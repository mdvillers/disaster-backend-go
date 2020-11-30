package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdvillers/disaster-backend-go/controllers"
)

//InitalizeRoutes initalizes all routes
func InitalizeRoutes(r *mux.Router) {
	//admin routes
	r.HandleFunc("/admin/getuser", controllers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/admin/create", controllers.CreateAdmin).Methods(http.MethodPost)
	r.HandleFunc("/admin/signin", controllers.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/admin/{id}", controllers.DeleteAdmin).Methods(http.MethodDelete)

	//district routes
	r.HandleFunc("/district/view", controllers.GetDistricts).Methods(http.MethodGet)
	r.HandleFunc("/district/insert", controllers.InsertDistrict).Methods(http.MethodPost)
	r.HandleFunc("/district/delete/{id}", controllers.DeleteDistrict).Methods(http.MethodDelete)

	//vm routes
	r.HandleFunc("/vm/view", controllers.GetVMs).Methods(http.MethodGet)
	r.HandleFunc("/vm/insert", controllers.InsertVM).Methods(http.MethodPost)
	r.HandleFunc("/vm/delete/{id}", controllers.DeleteVM).Methods(http.MethodDelete)
}
