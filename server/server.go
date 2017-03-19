package server

import (
	"fmt"
	"github.com/SamirMarin/saveMeMicroService/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

// Handler for POST requests on /help
// Extracts Emergency info from POST request, and stores in database
func help(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	var emrInfo models.EmergencyInfo
	var latLong models.LatLong
	lat := convertToFloat(ps.ByName("lat"))
	lon := convertToFloat(ps.ByName("lon"))
	latLong.Lat = lat
	latLong.Lon = lon
	priority := convertToInt(ps.ByName("priority"))
	description := ps.ByName("description")
	time := ps.ByName("time")
	emrInfo.Description = description
	emrInfo.Priority = priority
	emrInfo.UpdateTime = time
	emrInfo.Location = latLong
	emrInfo.StoreEmergencyInfo()
}

// Handler for GET requests on /map
// Returns a JSON that contains all the EmergencyInfo entries in the db
// that have a date greater than ps.time
func getMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "THIS IS Were I send a map to Ios\n")
}
func convertToFloat(flt string) float64 {
	f, err := strconv.ParseFloat(flt, 64)

	if err != nil {
		log.Fatal(err)
	}
	return f
}
func convertToInt(strInt string) int {
	newInt, err := strconv.Atoi(strInt)
	if err != nil {
		log.Fatal(err)
	}
	return newInt

}

func Run() {
	router := httprouter.New()
	router.POST("/help", help)
	router.GET("/map", getMap)
	log.Fatal(http.ListenAndServe(":8080", router))
}
