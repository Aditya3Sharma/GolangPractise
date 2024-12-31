package actualcontrollers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practise/controllers/helpers"
	employee "practise/models"

	"github.com/gorilla/mux"
)

func GetAllEmployeesData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	employees := helpers.GetAllEmployees()
	json.NewEncoder(w).Encode(employees)
}

func InsertOneEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Mthods", "POST")

	var employee employee.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)

	helpers.InsertOneEmployee(employee)
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlncode")
	w.Header().Set("Allow-Control-Allow-Mthods", "PUT")

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

	params := mux.Vars(r)
	helpers.DeleteOneEmployee(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlncode")
	w.Header().Set("Allow-Control-Allow-Mthods", "DELETE")

	helpers.DeleteAllEmployee()
	json.NewEncoder(w).Encode("Deleted all")

}
