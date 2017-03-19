package db

import "gopkg.in/mgo.v2"

var (
	Session mgo.Session
)

func ConnectMongoDb () {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	Session = *session
}
