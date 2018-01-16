package repository

import (
	"gopkg.in/mgo.v2"
	"fmt"
)


func MongoConnect() *mgo.Session{

	session, err := mgo.Dial("10.50.115.16:27017")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}