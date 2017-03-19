package models

import (
	"time"
	"github.com/SamirMarin/saveMeMicroService/db"
	"gopkg.in/mgo.v2/bson"
)

type EmergencyInfo struct {
	id		int		`json:"id"`
	emergencyType	string		`json:"emergency_type"`
	priority	int		`json:"priority"`
	location	time.Location	`json:"location"`
	updateTime	time.Time	`json:"updateTime"`
}

func (e EmergencyInfo) storeEmergencyInfo() {
	localSession := db.Session.Copy()
	defer localSession.Close()

	addInfo := localSession.DB("db").C("info")
	err := addInfo.Insert(e)
	if err != nil {
		panic(err)
	}
}

func (e EmergencyInfo) getAllEmergencyInfo() []EmergencyInfo {
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

func (e EmergencyInfo) removeEmergencyInfo() {
	localSession := db.Session.Copy()
	defer localSession.Close()

	addInfo := localSession.DB("db").C("info")
	err := addInfo.Remove(bson.M{"id" : e.id})
	if err != nil {
		panic(err)
	}
}
