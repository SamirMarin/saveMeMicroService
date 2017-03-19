package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func help(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")

}

func getMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "THIS IS Were I send a map to Ios\n")
}

func Run() {
	router := httprouter.New()
	router.POST("/help", help)
	router.GET("/map", getMap)
	log.Fatal(http.ListenAndServe(":8080", router))
}
