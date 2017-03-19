package server

import (
	"fmt"
	"github.com/SamirMarin/saveMeMicroService/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Handler for POST requests on /help
// Extracts Emergency info from POST request, and stores in database
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

// Handler for GET requests on /map
// Returns a JSON that contains all the EmergencyInfo entries in the db
func getMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "THIS IS Were I send a map to Ios\n")
}

func Run() {
	router := httprouter.New()
	router.POST("/help", help)
	router.GET("/map", getMap)
	log.Fatal(http.ListenAndServe(":8080", router))
}
