package models

import (
	"github.com/SamirMarin/saveMeMicroService/db"
	"gopkg.in/mgo.v2/bson"
)

type EmergencyInfo struct {
	Id          string	`json:"id"`
	Description string	`json:"description"`
	Priority    int		`json:"priority"`
	Location    LatLong	`json:"location"`
	UpdateTime  string	`json:"update_time"`
}

type LatLong struct {
	Lat float64
	Lon float64
}

func (e EmergencyInfo) StoreEmergencyInfo() {
	localSession := db.Session.Copy()
	defer localSession.Close()
	localDB := localSession.DB("db").C("info")
	err := localDB.Insert(e)
	if err != nil {
		panic(err)
	}
}

func (e EmergencyInfo) GetAllEmergencyInfo(limit string) []EmergencyInfo {
	localSession := db.Session.Copy()
	defer localSession.Close()
	var allEmergencyInfo []EmergencyInfo
	localDB := localSession.DB("db").C("info")
	err := localDB.Find(bson.M{"update_time" : {"$gte" : limit}}).All(&allEmergencyInfo)
	if err != nil {
		panic(err)
	}
	return allEmergencyInfo
}

func (e EmergencyInfo) RemoveEmergencyInfo() {
	localSession := db.Session.Copy()
	defer localSession.Close()

	addInfo := localSession.DB("db").C("info")
	err := addInfo.Remove(bson.M{"id": e.Id, "description": e.Description, "priority": e.Priority,
		"location": e.Location, "update_time": e.UpdateTime})
	if err != nil {
		//panic(err)
	}
}
