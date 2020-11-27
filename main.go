package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mdvillers/disaster-backend-go/db"

	"github.com/mdvillers/disaster-backend-go/controllers"

	"github.com/mdvillers/disaster-backend-go/config"

	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/admin/getuser", controllers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/admin/create", controllers.CreateAdmin).Methods(http.MethodPost)
	r.HandleFunc("/admin/signin", controllers.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/admin/{id}", controllers.DeleteAdmin).Methods(http.MethodDelete)
	fmt.Println(config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}

func main() {
	config.Load()
	db.Connect()
	handleRequests()
}
