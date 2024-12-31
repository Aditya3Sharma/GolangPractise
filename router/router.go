package router

import (
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

	return router

}
