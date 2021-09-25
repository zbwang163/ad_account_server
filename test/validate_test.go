package test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestValidate(t *testing.T) {
	mac := hmac.New(sha256.New, []byte("ustc"))
	mac.Write([]byte("i love you"))
	expectedMAC := mac.Sum(nil)
	t.Logf("%x", expectedMAC)
	t.Log(hex.EncodeToString(expectedMAC))
}
