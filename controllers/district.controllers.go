package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mdvillers/disaster-backend-go/db"
	"github.com/mdvillers/disaster-backend-go/models"

	"github.com/gorilla/mux"
)

//GetDistricts implies getting all districts
func GetDistricts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var districts []models.District
	db.DB.Preload("vms").Find(&districts)
	json.NewEncoder(w).Encode(districts)
}

//InsertDistrict means to Create an district
func InsertDistrict(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var district models.District
	_ = json.NewDecoder(r.Body).Decode(&district)
	db.DB.Create(&district)
	json.NewEncoder(w).Encode(district)
}

//DeleteDistrict means to Delete an district
func DeleteDistrict(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var district models.District
	db.DB.Where("id=?", id).Find(&district)
	db.DB.Delete(&district)
	json.NewEncoder(w).Encode(district)
}
