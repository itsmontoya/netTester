package main

import (
	"time"

	//"github.com/itsmontoya/go/src/pkg/net/http"
	"net/http"
)

const (
	statusCode      = 200
	jsonContentType = "application/json"
	jsonStr         = `{ "greeting" : "Hello world!" }`
)

var jsonB = []byte(jsonStr)

func main() {
	var s srv
	s.Listen(":8080")
}

type srv struct{}

func (s *srv) Listen(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *srv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", jsonContentType)
	switch r.URL.Path {
	case "/a":
	case "/b":
		time.Sleep(time.Millisecond * 1)
	case "/c":
		time.Sleep(time.Millisecond * 5)
	case "/d":
		time.Sleep(time.Millisecond * 10)
	}

	w.Write(jsonB)
}
