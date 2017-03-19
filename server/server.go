package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"github.com/SamirMarin/saveMeMicroService/models"
)

func help(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	var emrInfo models.EmergencyInfo
	var latLong models.LatLong
	lat := float64(ps.ByName("lat"))
	lon := float64(ps.ByName("lon"))
	latLong.Lat = lat
	latLong.Lon = lon
	priority := int(ps.ByName("priority"))
	description := ps.ByName("EmergencyType")
	time := ps.ByName("time")
	emrInfo.EmergencyType = description
	emrInfo.Priority = priority
	emrInfo.UpdateTime = time
	emrInfo.Location = latLong
	emrInfo.StoreEmergencyInfo()
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



