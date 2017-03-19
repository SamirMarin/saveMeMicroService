package main

import (
	"github.com/SamirMarin/saveMeMicroService/server"
	"github.com/SamirMarin/saveMeMicroService/db"
)

func main(){
	db.ConnectMongoDb()
	server.Run()
}
