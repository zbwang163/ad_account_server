package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"os"
	"testing"
)

func TestPutObject(t *testing.T) {
	InitMinIO()
	ctx := context.Background()
	bucketName := "mymusic"
	objectName := "IK0e.jpg"
	filePath := "./test_data/lK0e.jpg"
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = PutObject(ctx, bucketName, objectName, f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutPreSignedObject(t *testing.T) {
	InitMinIO()
	ctx := context.Background()
	bucketName := "mymusic"
	objectName := "capture1.jpg"
	//filePath := "./test_data/lK0e.jpg"
	//f, err := os.Open(filePath)
	//defer f.Close()
	//if err != nil {
	//	t.Fatal(err)
	//}
	u, err := PutPreSignedObject(ctx, bucketName, objectName)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestMinIO(t *testing.T) {
	InitMinIO()
	ctx := context.Background()
	bucketName := "mymusic"
	//location := "us-east-1"
	//err := MinIOClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	//if err != nil {
	//	// Check to see if we already own this bucket (which happens if you run this twice)
	//	exists, errBucketExists := MinIOClient.BucketExists(ctx, bucketName)
	//	if errBucketExists == nil && exists {
	//		log.Printf("We already own %s\n", bucketName)
	//	} else {
	//		log.Fatalln(err)
	//	}
	//} else {
	//	log.Printf("Successfully created %s\n", bucketName)
	//}

	// Upload the zip file
	objectName := "capture.jpg"
	filePath := "./test_data/zwRE.jpg"

	// Upload the zip file with FPutObject
	info, err := Client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	filePath1 := "./test_data/7243111.jpg"
	_ = Client.FGetObject(ctx, bucketName, objectName, filePath1, minio.GetObjectOptions{})
}
