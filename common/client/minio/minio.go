package minio

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/zbwang163/ad_account_server/common/biz_error"
	"github.com/zbwang163/ad_account_server/common/logs"
	"io"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Client *minio.Client

type MinIo interface {
	PutObject(ctx context.Context, bucket string, name string, obj io.Reader) error
	GetObject()
}

func InitMinIO() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false
	var err error
	// Initialize minio client object.
	Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func PutObject(ctx context.Context, bucket string, name string, obj io.Reader) error {
	// 判断bucket是否存在, 不存在创建一个bucket
	exists, errBucketExists := Client.BucketExists(ctx, bucket)
	if errBucketExists == nil && !exists {
		err := Client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			logs.CtxError(ctx, "minio bucket创建失败, err:%v", err)
			return biz_error.NewMinIOError(err)
		}
	}
	buffer := bytes.NewBuffer([]byte{})
	n, err := buffer.ReadFrom(obj)
	if err != nil {
		logs.CtxError(ctx, "param reader err:%v", err)
		return biz_error.NewInternalError(errors.New("读取输入流失败"))
	}

	var contentType string
	if strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".jpeg") {
		contentType = "image/jpeg"
	} else if strings.HasSuffix(name, "png") {
		contentType = "image/png"
	}

	info, err := Client.PutObject(ctx, bucket, name, buffer, n, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logs.CtxError(ctx, "minio上传对象失败,err:%v", err)
		return biz_error.NewMinIOError(err)
	}
	logs.CtxInfo(ctx, "minio info:%v", info)
	return nil
}

func PreSignedObjectUrl(ctx context.Context, bucket string, name string, expire time.Duration) (string, error) {
	reqParams := make(url.Values)
	u, err := Client.PresignedGetObject(ctx, bucket, name, expire, reqParams)
	return u.String(), err
}

func PutPreSignedObject(ctx context.Context, bucket string, name string) (string, error) {
	//reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
	policy := minio.NewPostPolicy()
	policy.SetBucket(bucket)
	policy.SetKey(name)
	// Expires in 10 days.
	policy.SetExpires(time.Now().UTC().Add(time.Second * 120))

	u1, f, e := Client.PresignedPostPolicy(ctx, policy)
	fmt.Println(u1, f, e)
	u, err := Client.PresignedPutObject(ctx, bucket, name, time.Second*120)
	logs.CtxInfo(ctx, "%s", u)
	return u.String(), err
}
