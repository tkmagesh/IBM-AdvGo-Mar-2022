package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Hello World\n"))
}

func greet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("Hello %s\n", name)))
}

func greetPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("[POST] Hello %s\n", name)))
}
func main() {
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("/Users/tkmagesh77/Documents/Training/IBM-AdvGo-Nov-2021/03-http-services/02-httprouter/static"))
	router.GET("/", index)
	router.GET("/greet/:name", greet)
	router.POST("/greet/:name", greetPost)
	http.ListenAndServe(":8080", router)
}

/*

 */
