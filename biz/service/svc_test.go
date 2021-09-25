package service

import (
	"context"
	"testing"
)

var accountSvc AccountServiceImpl

func TestAccountServiceImpl_GenCapture(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(accountSvc.GenCapture(context.Background(), 6))
	}
}

func TestAccountServiceImpl_GenCaptureImage(t *testing.T) {
	accountSvc.GenCaptureImage(context.Background(), 4)
}
