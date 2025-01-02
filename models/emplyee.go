package employee

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	Designation string             `json:"designation"`
	Salary      int                `json:"salary"`
}

type Login struct {
	Username       string `json:"username,omitempty"`
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}
