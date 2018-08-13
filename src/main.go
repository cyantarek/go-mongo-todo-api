package main

import (
	"log"
	"github.com/cyantarek/go-mongo-rest-api-crud/src/db"
	"github.com/cyantarek/go-mongo-rest-api-crud/src/routes"
	"github.com/cyantarek/go-mongo-rest-api-crud/src/handlers"
	"net/http"
)

const address = ":8078"

func main() {
	db.Connect()

	r := routes.GetRoutes()

	log.Println("Server started at http://localhost" + address)
	log.Fatalln(http.ListenAndServe(address, handlers.JsonResponse(r)))
}
