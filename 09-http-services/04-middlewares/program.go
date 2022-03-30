package main

import (
	"fmt"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func main() {

	/*
		http.HandleFunc("/foo", profile(logger(foo)))
		http.HandleFunc("/bar", profile(logger(bar)))
	*/

	http.HandleFunc("/foo", chain(foo, logger, profile))
	http.HandleFunc("/bar", chain(bar, logger, profile))

	http.ListenAndServe(":8080", nil)
}

func logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		handler(w, r)
	}
}

func profile(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			elapsed := end.Sub(start) / time.Millisecond
			fmt.Printf("%s took %dms\n", r.URL.Path, elapsed)
		}()
		handler(w, r)
	}
}

func chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func foo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Write([]byte("foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Write([]byte("bar"))
}
