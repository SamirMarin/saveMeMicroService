package models

import (
	"github.com/SamirMarin/saveMeMicroService/db"
	"gopkg.in/mgo.v2/bson"
)

type EmergencyInfo struct {
	Id          string  `json:"id"`
	Desc        string  `json:"desc"`
	Priority    int     `json:"priority"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	UpdateTime  string  `json:"update_time"`
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
	err := addInfo.Remove(bson.M{"id": e.Id, "desc": e.Desc, "priority": e.Priority,
		"lat": e.Lat, "lon": e.Lon, "update_time": e.UpdateTime})
	if err != nil {
		//panic(err)
	}
}
