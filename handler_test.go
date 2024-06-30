package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello Brooo</h1>")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "<h1>Hello Bro</h1>")
	})
	mux.HandleFunc("/ar", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "<h1>Nama Saya Aris</h1>")
	})
	mux.HandleFunc("/is", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "<h1>Di Malam yang dingin</h1>")
	})
	mux.HandleFunc("/mah", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "<h1>dan gelap sepi</h1>")
	})
	mux.HandleFunc("/mu/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "<h1>Maafkanlah Aku</h1>")
	})
	mux.HandleFunc("/mu/di/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "<h1>Hancurkan dirimu, Sumpah serapah.</h1>")
	})
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
