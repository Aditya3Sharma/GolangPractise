package router

import (
	"practise/authentication"
	actualcontrollers "practise/controllers/actualControllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/employees", actualcontrollers.GetAllEmployeesData).Methods("GET")
	router.HandleFunc("/api/employee/{id}", actualcontrollers.UpdateEmployeeData).Methods("PUT")
	router.HandleFunc("/api/employees", actualcontrollers.InsertOneEmployeeData).Methods("POST")
	router.HandleFunc("/api/employee/{id}", actualcontrollers.DeleteOneEmployeeData).Methods("DELETE")
	router.HandleFunc("/api/employees", actualcontrollers.DeleteAllEmployeeData).Methods("DELETE")

	router.HandleFunc("/api/signup", authentication.SignUpHandler).Methods("POST")
	router.HandleFunc("/api/login", authentication.LoginHandler).Methods("POST")
	router.HandleFunc("/api/logout", authentication.LogoutHandler).Methods("POST")

	return router

}
