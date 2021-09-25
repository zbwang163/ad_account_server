package smtp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-mailer/send"
	"net/mail"
	"sync"
)

var (
	Host    = "smtp.163.com:25"
	From    = mail.Address{Name: "广告信息平台", Address: "zbwang163@163.com"}
	FromPwd = "ZEOSZOEQEFQASZNV"
)

func Send(ctx context.Context, to string, content string) error {
	sender, err := send.NewSmtpSender(Host, From, FromPwd)
	if err != nil {
		return err
	}
	msg := &send.Message{
		Subject: "广告信息平台邮箱验证码",
		Content: bytes.NewBufferString(fmt.Sprintf("<h1>%s</h1>", content)),
		To:      fmt.Sprintf("<%s>", to),
	}
	err = sender.Send(msg, false)
	if err != nil {
		return err
	}
	return nil
}

func AsyncSend(ctx context.Context, to string, content string) {
	sender, err := send.NewSmtpSender(Host, From, FromPwd)
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	msg := &send.Message{
		Subject: "广告信息平台邮箱验证码",
		Content: bytes.NewBufferString(fmt.Sprintf("<h1>%s</h1>", content)),
		To:      fmt.Sprintf("<%s>", to),
	}
	err = sender.AsyncSend(msg, false, func(err error) {
		defer wg.Done()
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	wg.Wait()
}
