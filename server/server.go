package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func help(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func getMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Run() {
	router := httprouter.New()
	router.POST("/help", help)
	router.GET("/hello", getMap)
	log.Fatal(http.ListenAndServe(":8080", router))
}
