package models

import (
	"time"
	"github.com/SamirMarin/saveMeMicroService/db"
	"gopkg.in/mgo.v2/bson"
)

type EmergencyInfo struct {
	Id		int		`json:"id"`
	EmergencyType	string		`json:"emergency_type"`
	Priority	int		`json:"priority"`
	Location	LatLong		`json:"location"`
	UpdateTime	string  	`json:"updateTime"`
}
type LatLong struct {
	Lat float64
	Lon float64
}

func (e EmergencyInfo) StoreEmergencyInfo() {
	localSession := db.Session.Copy()
	defer localSession.Close()
	addInfo := localSession.DB("db").C("info")
	err := addInfo.Insert(e)
	if err != nil {
		panic(err)
	}
}

func (e EmergencyInfo) GetAllEmergencyInfo() []EmergencyInfo {
	localSession := db.Session.Copy()
	defer localSession.Close()
	var allEmergencyInfo []EmergencyInfo
	addInfo := localSession.DB("db").C("info")
	err := addInfo.Find(bson.M{}).All(&allEmergencyInfo)
	if err != nil {
		panic(err)
	}
	return allEmergencyInfo
}

func (e EmergencyInfo) RemoveEmergencyInfo() {
	localSession := db.Session.Copy()
	defer localSession.Close()

	addInfo := localSession.DB("db").C("info")
	err := addInfo.Remove(bson.M{"id" : e.id})
	if err != nil {
		panic(err)
	}
}
