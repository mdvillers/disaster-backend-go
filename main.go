package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mdvillers/disaster-backend-go/db"
	"github.com/mdvillers/disaster-backend-go/routes"

	"github.com/mdvillers/disaster-backend-go/config"

	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	routes.InitalizeRoutes(r)
	fmt.Println(config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}

func main() {
	config.Load()
	db.Connect()
	handleRequests()
}
