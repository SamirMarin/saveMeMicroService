package server

import (
	"github.com/SamirMarin/saveMeMicroService/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"html/template"
	"fmt"
)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("server/templates/*"))
}

// Handler for POST requests on /help
// Extracts Emergency info from POST request, and stores in database
func help(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("IM here!!!!")
	var emrInfo models.EmergencyInfo
	if r.Body == nil {
		fmt.Println("body is nil")
	}
	err := json.NewDecoder(r.Body).Decode(&emrInfo)
	if err != nil {
		fmt.Println("json error")
	}
	fmt.Println(emrInfo.Id)
	fmt.Println(emrInfo.Desc)
	fmt.Println(emrInfo.Priority)
	fmt.Println(emrInfo.Lat)
	fmt.Println(emrInfo.Lon)
	fmt.Println(emrInfo.UpdateTime)
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
func test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//title := "edit"
	type data struct{
		Title string
		Body string
	}
	//dataVal := data{"hello","test"}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, nil)
}
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "edit.html", nil)
	if err != nil {
		fmt.Println("error")
	}
}


func Run() {
	router := httprouter.New()
	router.GET("/index", index)
	router.POST("/help", help)
	router.GET("/map/:updateTime", getMap)
	fmt.Println("now before")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("here")
}
