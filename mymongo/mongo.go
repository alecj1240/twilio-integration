package mymongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ZeitID           string
	TwilioKeySID     string
	TwilioKeySecret  string
	TwilioAccountSID string
}

func Connect(client *mongo.Client) *mongo.Collection {
	collection := client.Database("hackathon").Collection("users")

	return collection
}

func CreateUser(client *mongo.Client, ZeitID string, TwilioKeySID string, TwilioKeySecret string, TwilioAccountSID string) {
	newUser := User{ZeitID, TwilioKeySID, TwilioKeySecret, TwilioAccountSID}
	collection := Connect(client)
	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func FindUser(client *mongo.Client, ZeitID string) User {
	filter := bson.D{{"zeitid", ZeitID}}
	collection := Connect(client)
	// create a value into which the result can be decoded
	var result User

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return User{"", "", "", ""}
	}

	return result
}
