package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Patient struct to receive data from DB
type Patient struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FName     string             `json:"fname"`
	LName     string             `json:"lname"`
	Email     string             `json:"email"`
	BloodType string             `json:"blood"`
	Cpf       int                `json:"cpf"`
	Birth     int                `json:"birth"`
	Phone     int                `json:"phone"`
	Mobile    int                `json:"mobile"`
}

func main() {
	//get a connection creating a new client for handling it.
	client, err := mongo.Connect(nil, options.Client().ApplyURI("mongodb://192.168.122.95:27017"))
	if err != nil {
		log.Fatal(err)
	}

	//get a collection pointer to query
	collection := client.Database("clinic").Collection("patient")

	//get a cursor as find object
	cur, err := collection.Find(nil, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(nil)

	//loop throught cursor to get the items from the collection
	for cur.Next(nil) {
		var result Patient
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
