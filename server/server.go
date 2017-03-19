package server

import (
	"github.com/SamirMarin/saveMeMicroService/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
)

// Handler for POST requests on /help
// Extracts Emergency info from POST request, and stores in database
func help(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var emrInfo models.EmergencyInfo
	id := ps.ByName("id")
	description := ps.ByName("desc")
	priority := convertToInt(ps.ByName("priority"))
	lat := convertToFloat(ps.ByName("lat"))
	lon := convertToFloat(ps.ByName("lon"))
	updateTime := ps.ByName("updateTime")
	emrInfo.Id = id
	emrInfo.Desc = description
	emrInfo.Priority = priority
	emrInfo.Lat = lat
	emrInfo.Lon = lon
	emrInfo.UpdateTime = updateTime
	emrInfo.StoreEmergencyInfo()
}

// Handler for GET requests on /map
// Returns a JSON that contains all the EmergencyInfo entries in the db
// that have a date greater than ps.time
func getMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var emrInfo models.EmergencyInfo
	emrInfo.UpdateTime = ps.ByName("updateTime")
	lstEmrInfo := emrInfo.GetAllEmergencyInfo(emrInfo.UpdateTime)
	json.NewEncoder(w).Encode(lstEmrInfo)
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
	router.GET("/map/:updateTime", getMap)
	log.Fatal(http.ListenAndServe(":8080", router))
}
