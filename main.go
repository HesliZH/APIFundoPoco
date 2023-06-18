package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//Tratamento do Start da API
	log.Fatal(http.ListenAndServe(":8080", router))
}
