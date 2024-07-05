package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Bro</h1>")
}

func TestHttp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	bString := string(body)

	fmt.Println(bString)
}
