package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	port      = "6969"
	mongoHost = "localhost"
	mongoPort = "27017"
	counter   = atomic.Int64{}
)

func main() {
	if newPort, ok := os.LookupEnv("PORT"); ok {
		port = newPort
	}
	if newmp, ok := os.LookupEnv("MONGO_PORT"); ok {
		mongoPort = newmp
	}
	if newmh, ok := os.LookupEnv("MONGO_HOST"); ok {
		mongoHost = newmh
	}
	log.Printf("Args: port: %s, mongoHost: %s, mongoPort: %s", port, mongoHost, mongoPort)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	usersCollection := client.Database("testing").Collection("users")

	rtr := mux.NewRouter()
	rtr.HandleFunc("/user/{id}", getUserRouter(usersCollection)).Methods("GET")
	rtr.HandleFunc("/counter", counterRouter).Methods("GET")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

var m map[string]bool = map[string]bool{}

func getUserRouter(c *mongo.Collection) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		counter.Add(1)
		params := mux.Vars(r)
		id := params["id"]
		filter := bson.D{
			{Key: "$and",
				Value: bson.A{
					bson.D{
						{Key: "jmbg", Value: bson.D{
							{Key: "$eq", Value: id}}},
					},
				},
			},
		}
		result := c.FindOne(context.TODO(), filter)
		if result.Err() == nil {
			w.WriteHeader(400)
			w.Write([]byte("Already registered: " + id))
			return
		}
		w.Write([]byte("Hello " + id))
		d := bson.M{"jmbg": id}
		_, err := c.InsertOne(context.TODO(), d)
		if err != nil {
			panic(err)
		}
		log.Println("Registered: ", id)
	}
}

func counterRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%d", counter.Load())))
}
