package actualcontrollers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practise/authentication"
	"practise/controllers/helpers"
	employee "practise/models"

	"github.com/gorilla/mux"
)

func GetAllEmployeesData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	if err := authentication.Authorized(r); err != nil {
		return
	}
	employees := helpers.GetAllEmployees()
	json.NewEncoder(w).Encode(employees)
}

func InsertOneEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Mthods", "POST")

	if err := authentication.Authorized(r); err != nil {
		return
	}
	var employee employee.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)

	helpers.InsertOneEmployee(employee)
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlncode")
	w.Header().Set("Allow-Control-Allow-Mthods", "PUT")

	if err := authentication.Authorized(r); err != nil {
		return
	}

	var employee employee.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)

	params := mux.Vars(r)

	fmt.Println(employee)
	helpers.UpdateEmployee(employee, params["id"])
	json.NewEncoder(w).Encode(employee)
}

func DeleteOneEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlncode")
	w.Header().Set("Allow-Control-Allow-Mthods", "DELETE")

	if err := authentication.Authorized(r); err != nil {
		return
	}

	params := mux.Vars(r)
	helpers.DeleteOneEmployee(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlncode")
	w.Header().Set("Allow-Control-Allow-Mthods", "DELETE")

	if err := authentication.Authorized(r); err != nil {
		return
	}

	helpers.DeleteAllEmployee()
	json.NewEncoder(w).Encode("Deleted all")

}
