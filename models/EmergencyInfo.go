package models

import (
	"time"
)

type EmergencyInfo struct {
	ID		int
	emergencyType	string
	priority	int
	location	time.Location
	emergencyTime	time.Time
}

func (e EmergencyInfo) storeInfo() {

}

func (e EmergencyInfo) getAllEmergencyInfo() {

}
