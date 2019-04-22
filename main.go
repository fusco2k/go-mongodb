package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Patient struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.122.95:27017"))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("clinic").Collection("patient")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	
	for cur.Next(ctx) {
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
