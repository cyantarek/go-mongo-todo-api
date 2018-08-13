package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

var Coll *mgo.Collection

func Connect() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		log.Println(err)
	}
	Coll = session.DB("todo-db").C("todos")
}