package main

import (
	//	"fmt"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	testPort = ":8080"
	testAddr = "http://localhost" + testPort
)

var (
	testClnt           = &http.Client{}
	errInvalidResponse = errors.New("invalid response")
)

func TestMain(m *testing.M) {
	s := &srv{}
	go s.Listen(testPort)
	time.Sleep(time.Second)
	sc := m.Run()

	os.Exit(sc)
}

func TestBasic(t *testing.T) {
	if err := httpReq(testAddr); err != nil {
		t.Error(err)
	}
}

func BenchmarkBasic(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		if err = httpReq(testAddr); err != nil {
			b.Error(err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkParaBasic(b *testing.B) {
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		var err error
		for pb.Next() {
			if err = httpReq(testAddr); err != nil {
				b.Error(err)
			}
		}
	})

	b.ReportAllocs()
}

func httpReq(loc string) (err error) {
	fmt.Println("Req start")
	var (
		resp *http.Response
		buf  [len(jsonStr)]byte
	)

	if resp, err = testClnt.Get(loc); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		goto END
	}

	if _, err = resp.Body.Read(buf[:]); err != nil {
		if err != io.EOF {
			goto END
		}

		err = nil
	}

	if str := string(buf[:]); str != jsonStr {
		err = errInvalidResponse
		goto END
	}

END:
	resp.Body.Close()
	fmt.Println("Request done", err)
	return
}
