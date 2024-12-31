package connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func GetCollection() *mongo.Collection {
	godotenv.Load()
	DB_PASS := os.Getenv("DBPassword")

	var connectionString = "mongodb+srv://aditya3sharma:" + DB_PASS + "@cluster0.rmyb5.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	var dbName = "angeleone"
	var collName = "employees"
	clienOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clienOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success connecting MongoDB")

	collection = client.Database(dbName).Collection(collName)

	return collection
}
