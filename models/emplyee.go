package employee

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Designation string             `json:"designation"`
	Salary      int                `json:"salary"`
}
