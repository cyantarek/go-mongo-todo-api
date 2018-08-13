package tests

import (
	"testing"
	"gopkg.in/mgo.v2"
	"net/http"
)

var client = &http.Client{}

func TestMongoDb(t *testing.T) {
	db, err := mgo.Dial("localhost")

	if err != nil || db.Ping() != nil {
		t.Error("Mongodb is not available")
	}
}

func TestGetAllTodo(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8078/todos", nil)

	if err != nil {
		t.Fail()
	}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fail()
	}
}