package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mdvillers/disaster-backend-go/db"
	"github.com/mdvillers/disaster-backend-go/models"

	"github.com/gorilla/mux"
)

//GetVMs implies getting all vms
func GetVMs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vms []models.VM
	db.DB.Find(&vms)
	json.NewEncoder(w).Encode(vms)
}

//InsertVM means to Create an vm
func InsertVM(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vm models.VM
	_ = json.NewDecoder(r.Body).Decode(&vm)
	db.DB.Create(&vm)
	json.NewEncoder(w).Encode(vm)
}

//DeleteVM means to Delete an vm
func DeleteVM(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var vm models.VM
	db.DB.Where("id=?", id).Find(&vm)
	db.DB.Delete(&vm)
	json.NewEncoder(w).Encode(vm)
}
