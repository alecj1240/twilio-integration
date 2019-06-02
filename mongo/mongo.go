package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ZeitID           string
	TwilioKeySID     string
	TwilioKeySecret  string
	TwilioAccountSID string
}

func Connect() *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://alecjones:C3ITZEfRfY4R%404F808H@clusterjones-0g5df.mongodb.net/test",
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("loonie").Collection("users")
	return collection
}

func CreateUser(ZeitID string, TwilioKeySID string, TwilioKeySecret string, TwilioAccountSID string) {
	newUser := User{ZeitID, TwilioKeySID, TwilioKeySecret, TwilioAccountSID}
	collection := Connect()
	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func FindUser(ZeitID string) User {
	filter := bson.D{{"zeitid", ZeitID}}
	collection := Connect()
	// create a value into which the result can be decoded
	var result User

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return User{"", "", "", ""}
	}

	return result
}
