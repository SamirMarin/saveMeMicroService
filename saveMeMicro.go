package main

import (
	"github.com/SamirMarin/saveMeMicroSevice/server"
	"github.com/SamirMarin/saveMeMicroSevice/db"
	"gopkg.in/mgo.v2"
	"fmt"
)


var (
	Session mgo.Session
)



func main(){
	Session = *db.ConnectMongoDb()
	server.Run(Session)
}
func TestMe() {
	fmt.Println("heelo")
}
