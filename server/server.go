package server

import (
	"encoding/json"
	"fmt"
	"github.com/SamirMarin/saveMeMicroService/models"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("server/templates/*"))
}

// Handler for POST requests on /help
// Extracts Emergency info from POST request, and stores in database
func help(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var emrInfo models.EmergencyInfo
	err := json.NewDecoder(r.Body).Decode(&emrInfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(emrInfo)
	emrInfo.StoreEmergencyInfo()
	fmt.Println("here we are about to get all info")
	fmt.Println(emrInfo.GetAllEmergencyInfo("2017-03-19T11:25:18.723Z"))
	fmt.Println("nothing in the info")
}

type arrOfEmergencyInfo struct {
	emergencyInfo []models.EmergencyInfo `json: "emergencyInfo"`
}
// Handler for GET requests on /map
// Returns a JSON that contains all the EmergencyInfo entries in the db
// that have a date greater than ps.time
func getMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Im here")
	var emrInfo models.EmergencyInfo
	emrInfo.UpdateTime = ps.ByName("updateTime")
	lstEmrInfo := emrInfo.GetAllEmergencyInfo("2017-03-19T11:25:18.723Z")
	resBody, err := json.MarshalIndent(lstEmrInfo, "", " ")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
	if err != nil {
		fmt.Println("this is the error", err)
	}
}
//test function written for testing without ios
func test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//title := "edit"
	type data struct {
		Title string
		Body  string
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
	log.Fatal(http.ListenAndServe("128.189.89.251:9999", router))
	fmt.Println("here")
}
