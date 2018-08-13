package routes

import (
	"github.com/gorilla/mux"
	"github.com/cyantarek/go-mongo-rest-api-crud/src/handlers"
	"net/http"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Handle("/", http.RedirectHandler("/todos", http.StatusFound)).Methods("GET")

	r.HandleFunc("/todos", handlers.GetAllTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.GetATodo).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateATodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateATodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteATodo).Methods("DELETE")
	return r
}
