package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	City string `json:"city,omitempty",`
}

func mapSubExpNames(regex string, s string) map[string]string {
	re := regexp.MustCompile(regex)
	m := re.FindStringSubmatch(s)
	n := re.SubexpNames()

	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i := 0; i < len(m); i++ {
		r[n[i]] = m[i]
	}
	return r
}

type DbStore struct {
	client *mongo.Client
}

var dbStore DbStore

func NewDbStore(uri string) DbStore {
	clientOptions := options.Client().ApplyURI(uri)
	client, e := mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		log.Fatal(e)
	}

	// test connection
	e = client.Ping(context.TODO(), nil)
	if e != nil {
		log.Fatal(e)
	}

	return DbStore{client}
}

func getTrainers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var trainers []Trainer
	cur, e := dbStore.client.Database("test").Collection("trainers").Find(context.TODO(), bson.D{{}}, options.Find())
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	for cur.Next(context.TODO()) {
		var trainer Trainer
		err := cur.Decode(&trainer)
		if err != nil {
			log.Fatal(err)
		}

		trainers = append(trainers, trainer)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trainers)
}

func getTrainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var trainer Trainer
	id, err := primitive.ObjectIDFromHex(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e := dbStore.client.Database("test").Collection("trainers").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&trainer)

	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trainer)
}

func createTrainer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t Trainer
	json.NewDecoder(r.Body).Decode(&t)

	result, e := dbStore.client.Database("test").Collection("trainers").InsertOne(context.TODO(), &t)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	// extracts the trainer id
	objId := fmt.Sprintf("%v", result.InsertedID)
	m := mapSubExpNames(`\"(?P<id>\w+)\"`, objId)

	// HTTP Location Header - /trainers/:id
	if id, ok := m["id"]; ok {
		w.Header().Set("Location", strings.Join([]string{r.URL.Path, id}, "/"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func updateTrainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var t Trainer
	var nt Trainer
	json.NewDecoder(r.Body).Decode(&t)

	id, e := primitive.ObjectIDFromHex(p.ByName("id"))
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": t.Name}}
	result, e := dbStore.client.Database("test").Collection("trainers").UpdateOne(context.TODO(), filter, update)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	if result.MatchedCount == 0 {
		http.Error(w, "Trainer does not exist", http.StatusNotFound)
		return
	}

	e = dbStore.client.Database("test").Collection("trainers").FindOne(context.TODO(), filter).Decode(&nt)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nt)
}

func deleteTrainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, e := primitive.ObjectIDFromHex(p.ByName("id"))
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	result, e := dbStore.client.Database("test").Collection("trainers").DeleteOne(context.TODO(), bson.M{"_id": id})
	fmt.Println(result.DeletedCount)
	if result.DeletedCount == 0 {
		http.Error(w, "Trainer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
func main() {
	dbStore = NewDbStore("mongodb://localhost:27017")
	router := httprouter.New()
	router.GET("/trainers", getTrainers)
	router.GET("/trainers/:id", getTrainer)
	router.POST("/trainers", createTrainer)
	router.DELETE("/trainers/:id", deleteTrainer)
	router.PUT("/trainers/:id", updateTrainer)

	log.Fatal(http.ListenAndServe(":8000", router))
}
