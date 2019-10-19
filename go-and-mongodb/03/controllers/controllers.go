package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"../models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	Client      *mongo.Client
	Collections map[string]models.Collection
}

func NewController(cli *mongo.Client) *Controller {
	collections := make(map[string]models.Collection)
	collections["users"] = models.Users{}

	return &Controller{cli, collections}
}

func (c Controller) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if collection, ok := c.Collections["users"]; ok {
		var docs []models.User
		filter := bson.D{{}}
		cur, e := c.Client.Database(collection.DbName()).Collection(collection.Name()).Find(context.TODO(), filter)
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}

		for cur.Next(context.TODO()) {
			var doc models.User
			e := cur.Decode(&doc)
			if e != nil {
				http.Error(w, e.Error(), http.StatusBadRequest)
				return
			}
			docs = append(docs, doc)
		}
		// closes the cursor
		cur.Close(context.TODO())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(docs)
	}
}

func (c Controller) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if collection, ok := c.Collections["users"]; ok {
		var doc models.User
		objId, _ := primitive.ObjectIDFromHex(p.ByName("id"))
		c.Client.Database(collection.DbName()).Collection(collection.Name()).FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&doc)
		if doc.Username == "" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doc)
	}
}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if collection, ok := c.Collections["users"]; ok {
		var body models.User
		var doc models.User
		json.NewDecoder(r.Body).Decode(&body)
		conn := c.Client.Database(collection.DbName()).Collection(collection.Name())

		result, e := conn.InsertOne(context.TODO(), body)
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}

		objId := result.InsertedID.(primitive.ObjectID)
		conn.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&doc)

		w.Header().Set("Location", strings.Join([]string{r.URL.Path, objId.Hex()}, "/"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(doc)
	}
}

func (c Controller) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if collection, ok := c.Collections["users"]; ok {
		objId, _ := primitive.ObjectIDFromHex(p.ByName("id"))
		result, e := c.Client.Database(collection.DbName()).Collection(collection.Name()).DeleteOne(context.TODO(), bson.M{"_id": objId})
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}

		if result.DeletedCount == 0 {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(204)
	}
}

func (c Controller) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if collection, ok := c.Collections["users"]; ok {
		var doc models.User
		var body interface{}
		objId, _ := primitive.ObjectIDFromHex(p.ByName("id"))
		conn := c.Client.Database(collection.DbName()).Collection(collection.Name())

		json.NewDecoder(r.Body).Decode(&body)
		result, e := conn.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": body})
		if e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}
		if result.MatchedCount == 0 {
			http.Error(w, "No User found", http.StatusNotFound)
			return
		}
		conn.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&doc)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doc)
	}
}
