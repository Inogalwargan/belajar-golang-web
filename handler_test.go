package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// func TestHandler(t *testing.T) {
// 	var handler http.HandlerFunc = func(writer http.ResponseWriter, r *http.Request) {
// 		// logic web
// 		fmt.Fprint(writer, "Hello World")
// 	}

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: handler,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Hello world")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Ini halaman hi")
	})

	mux.HandleFunc("/about", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Ini halaman about")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Images")
	})
	mux.HandleFunc("/images/thumbnails", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Thumbnail")
	})

	servers := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := servers.ListenAndServe()
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
