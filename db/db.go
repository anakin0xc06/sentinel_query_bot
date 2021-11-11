package db

import (
	"context"
	"fmt"
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/anakin0xc06/sentinel_query_bot/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDB ...
func NewDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

// UpdateUser ...
func UpdateUser(botDB *mongo.Collection, chatID int64, update bson.M) {
	var user types.User
	err := botDB.FindOne(context.TODO(), bson.M{"chatid": bson.M{"$eq": chatID}}).Decode(&user)
	if err != nil {
		insertResult, err := botDB.InsertOne(context.TODO(), bson.M{"chatid": chatID})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}
	botDB.UpdateOne(context.TODO(), bson.M{"chatid": bson.M{"$eq": chatID}}, update)
}

// DeleteUser ...
func DeleteUser(botDB *mongo.Collection, chatID int64) {
	var user types.User
	err := botDB.FindOne(context.TODO(), bson.M{"chatid": bson.M{"$eq": chatID}}).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}
	botDB.DeleteOne(context.TODO(), bson.M{"chatid": bson.M{"$eq": chatID}})
}

func UpdateStatus(botDB *mongo.Collection, username string, chatID int64, status string) {
	update := bson.M{"$set": bson.M{"status": status}}
	botDB.UpdateOne(context.TODO(), bson.M{"chatid": bson.M{"$eq": chatID}}, update)
}

func GetStatus(botDB *mongo.Collection, username string, chatID int64) string {
	var user types.User
	err := botDB.FindOne(context.TODO(), bson.M{"chatid": bson.M{"$eq": chatID}}).Decode(&user)
	if err != nil {
		log.Println(err)
		return ""
	}
	return user.Status
}
