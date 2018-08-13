package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"github.com/joho/godotenv"
)

var Coll *mgo.Collection

func Connect() {
	godotenv.Load()
	if addr := os.Getenv("MONGO_ADDRESS"); addr == "" {
		os.Setenv("MONGO_ADDRESS", "localhost:27017")
	}
	session, err := mgo.Dial(os.Getenv("MONGO_ADDRESS"))

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	Coll = session.DB("todo-db").C("todos")
}
