package models

import (
	"github.com/SamirMarin/saveMeMicroService/db"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type EmergencyInfo struct {
	Id          string  `json:"id"`
	Desc        string  `json:"desc"`
	Priority    int     `json:"priority"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	UpdateTime  string  `json:"updatetime"`
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
	fmt.Println("before the stuck", limit)
	err := localDB.Find(bson.M{"updatetime" : bson.M{"$gte" : limit}}).All(&allEmergencyInfo)
	fmt.Println("yes I go through")
	if err != nil {
		//panic(err)
	}
	return allEmergencyInfo
}

func (e EmergencyInfo) RemoveEmergencyInfo() {
	localSession := db.Session.Copy()
	defer localSession.Close()

	addInfo := localSession.DB("db").C("info")
	err := addInfo.Remove(bson.M{"id" : e.Id, "$and" : []interface{} {
		bson.M{"desc" : e.Desc, "$and" : []interface{} {
			bson.M{"priority" : e.Priority, "$and" : []interface{} {
				bson.M{"lat" : e.Lat, "$and" : []interface{} {
					bson.M{"lon" : e.Lon, "$and" : []interface{} {
						bson.M{"updatetime" : e.UpdateTime}}}}}}}}}}})
	if err != nil {
		//panic(err)
	}
}

func (e EmergencyInfo) GetEmergencyInfo(id string) EmergencyInfo {
	localSession := db.Session.Copy()
	defer localSession.Close()
	var emergencyInfo EmergencyInfo
	localDB := localSession.DB("db").C("info")
	err := localDB.Find(bson.M{"id" : id}).One(&emergencyInfo)
	if err != nil {
		//panic(err)
	}
	return emergencyInfo
}
