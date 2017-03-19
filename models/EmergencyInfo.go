package models

import (
	"time"
	"github.com/SamirMarin/saveMeMicroService/db"
	"gopkg.in/mgo.v2/bson"
)

type EmergencyInfo struct {
	id		int
	emergencyType	string
	priority	int
	location	time.Location
	emergencyTime	time.Time
}

func (e EmergencyInfo) storeInfo() {
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
	var allEmergencyInfos []EmergencyInfo
	addInfo := localSession.DB("db").C("info")
	err := addInfo.Find(bson.M{}).All(&allEmergencyInfos)
	if err != nil {
		panic(err)
	}
	return allEmergencyInfos
}
