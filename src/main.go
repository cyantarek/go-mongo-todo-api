package main

import (
	"log"
	"net/http"
)

const address = ":8078"

func main() {
	log.Println("Server started at http://localhost" + address)
	log.Println(http.ListenAndServe(address, nil))
}
