package main

import (
	"context"
	"log"
	"net/http"

	"./controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, e := mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		log.Fatal(e)
	}

	router := httprouter.New()
	controller := controllers.NewController(client)
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUser)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
