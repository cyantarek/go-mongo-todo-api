package handlers

import (
	"net/http"
	"github.com/cyantarek/go-mongo-rest-api-crud/src/models"
	"github.com/cyantarek/go-mongo-rest-api-crud/src/db"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/gorilla/mux"
	"log"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	db.Coll.Find(nil).All(&todos)
	json.NewEncoder(w).Encode(todos)
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	check := IDChecker(vars, w)

	if !check {
		return
	}

	var todo models.Todo
	err := db.Coll.FindId(bson.ObjectIdHex(vars["id"])).One(&todo)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func CreateATodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = bson.NewObjectId()
	todo.CreatedAt = time.Now()

	err := db.Coll.Insert(todo)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

func UpdateATodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	check := IDChecker(vars, w)

	if !check {
		return
	}

	var originalTodo models.Todo
	var updatedTodo models.Todo

	json.NewDecoder(r.Body).Decode(&updatedTodo)
	updatedTodo.UpdatedAt = time.Now()

	err := db.Coll.FindId(bson.ObjectIdHex(vars["id"])).One(&originalTodo)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	updatedTodo.ID = originalTodo.ID
	updatedTodo.CreatedAt = originalTodo.CreatedAt

	err = db.Coll.UpdateId(originalTodo.ID, &updatedTodo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "updated successfully"})

}

func DeleteATodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	check := IDChecker(vars, w)
	if !check {
		return
	}
	err := db.Coll.RemoveId(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "deleted successfully"})
}

func IDChecker(vars map[string]string, w http.ResponseWriter) bool {
	check := bson.IsObjectIdHex(vars["id"])
	if !check {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "invalid id given"})
		return false
	}
	return true
}
