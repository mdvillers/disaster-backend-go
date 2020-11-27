package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mdvillers/disaster-backend-go/db"
	"github.com/mdvillers/disaster-backend-go/models"

	"github.com/gorilla/mux"
)

//GetUser implies getting the single User info
func GetUser(w http.ResponseWriter, r *http.Request) {
	var admins []models.Admin
	db.DB.Find(&admins)
	json.NewEncoder(w).Encode(admins)
}

//CreateAdmin means to Create an admin
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	var admin models.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)
	db.DB.Create(&admin)
	json.NewEncoder(w).Encode(admin)
}

//DeleteAdmin means to Delete an admin
func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var admin models.Admin
	db.DB.Where("id=?", id).Find(&admin)
	db.DB.Delete(&admin)
	json.NewEncoder(w).Encode(admin)
}

//SignIn to get token
func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Signin here</h1>")
}
