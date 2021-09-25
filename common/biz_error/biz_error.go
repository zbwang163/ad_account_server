package biz_error

import "fmt"

type errorType struct {
	Code int
	Type string
}

type BizError struct {
	errorType
	ErrorMsg string
}

func (b *BizError) Error() string {
	if b == nil || b.Code == 0 {
		return ""
	}
	return fmt.Sprintf("code:%v, type:%v, error_message:%v", b.GetErrorCode(), b.GetErrorType(), b.GetBizErrorMessage())
}

func (b *BizError) GetErrorCode() int {
	return b.Code
}

func (b *BizError) GetErrorType() string {
	return b.Type
}

func (b *BizError) GetBizErrorMessage() string {
	return b.ErrorMsg
}

var (
	mysqlError    = errorType{1000, "读写mysql错误"}
	redisError    = errorType{1001, "读写redis错误"}
	resourceError = errorType{1002, "请求资源不存在或无权限"}
	paramError    = errorType{1003, "参数错误"}
	internalError = errorType{1004, "内部逻辑错误"}
	minIOError    = errorType{1005, "minio内部错误"}
)

func NewMysqlError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{mysqlError, errMsg}
}

func NewRedisError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{redisError, errMsg}
}

func NewResourceError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{resourceError, errMsg}
}

func NewParamError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{paramError, errMsg}
}

func NewMinIOError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{minIOError, errMsg}
}

func NewInternalError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{internalError, errMsg}
}
