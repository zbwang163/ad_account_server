package test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/zbwang163/ad_account_server/common/biz_error"
	"lukechampine.com/frand"
	"os"
	"runtime"
	"testing"
)

func main() float64 {
	var result float64
	var mergeSlice []int
	var n = len(mergeSlice)
	if n%2 == 0 {
		result = float64(mergeSlice[n/2]+mergeSlice[n/2+1]) / 2.0
	} else {
	}
	return result
}

func TestName4(t *testing.T) {
	type Model struct {
		Uid int64 `json:"uid"`
	}
	a := Model{Uid: 123456789}
	bs, _ := json.Marshal(a)
	var b map[string]string
	json.Unmarshal(bs, &b)
	t.Log(b["uid"])
}

func TestName3(t *testing.T) {
	var err error = nil
	if e, ok := err.(*biz_error.BizError); ok {
		t.Log("1", e)
	}
	t.Log("2")
}

func TestName2(t *testing.T) {
	file, err := os.Open("./buffer/test.txt") //test.txt的内容是“world”
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println(file.Sync())
	buf := bytes.NewBufferString("hello ")
	buf.ReadFrom(file)        //将text.txt内容追加到缓冲器的尾部
	fmt.Println(buf.String()) //打印“hello world”
}

func TestLogin(t *testing.T) {
	//resp, err := client.AccountClient.Login(context.Background(), &accountRpc.LoginRequest{})
	//if err != nil {
	//	t.Fatalf("err :%v", err)
	//}
	//t.Log(resp)
}

func TestCSPRNG(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := frand.New().Bytes(8)

		t.Logf("%x", a)
	}
}
func TestName(t *testing.T) {
	src := []byte("Hello Gopher!")
	s := hex.EncodeToString(src)
	t.Log(s)
	t.Logf("%x", src)
}

func TestName1(t *testing.T) {
	test(t)
}
func test(t *testing.T) {
	test2(t)
}

func test2(t *testing.T) {
	pc, file, line, ok := runtime.Caller(2)
	t.Log(pc)
	t.Log(file)
	t.Log(line)
	t.Log(ok)
	f := runtime.FuncForPC(pc)
	t.Log(f.Name())

	pc, file, line, ok = runtime.Caller(0)
	t.Log(pc)
	t.Log(file)
	t.Log(line)
	t.Log(ok)
	f = runtime.FuncForPC(pc)
	t.Log(f.Name())

	pc, file, line, ok = runtime.Caller(1)
	t.Log(pc)
	t.Log(file)
	t.Log(line)
	t.Log(ok)
	f = runtime.FuncForPC(pc)
	t.Log(f.Name())
}
