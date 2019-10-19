package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	var e error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, e = mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		log.Fatal(e)
	}

}

func main() {
	fmt.Println(client)
}
