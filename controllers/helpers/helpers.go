package helpers

import (
	"context"
	"fmt"
	"log"
	"practise/controllers/connection"
	employee "practise/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = connection.GetCollection()

func InsertOneEmployee(employee employee.Employee) {

	inserted, err := collection.InsertOne(context.Background(), employee)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted an emplyee in db with id:", inserted.InsertedID)
}

func UpdateEmployee(emp employee.Employee, employeeID string) {
	id, _ := primitive.ObjectIDFromHex(employeeID)
	fmt.Println(id)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"name":        emp.Name,
		"designation": emp.Designation,
		"salary":      emp.Salary,
	},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified employees count:", result.ModifiedCount)

}

func GetAllEmployees() []primitive.M {
	fmt.Println("Request to access all employees")
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var employees []primitive.M

	for cursor.Next(context.Background()) {
		var employee primitive.M
		err := cursor.Decode(&employee)

		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}
	defer cursor.Close(context.Background())
	return employees
}

func DeleteOneEmployee(employeeID string) {
	id, _ := primitive.ObjectIDFromHex(employeeID)
	filter := bson.M{"_id": id}

	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", deleted.DeletedCount)
}

func DeleteAllEmployee() {
	filter := bson.D{{}}
	deleted, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", deleted.DeletedCount)
}
