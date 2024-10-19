package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/manishSharma1-dev/goPractice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb://127.0.0.1:27017/"
const Dbname = "Gopractice"
const Collname = "dbpraccoll"

var Collection *mongo.Collection

func init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	ClientOptions := options.Client().ApplyURI(ConnectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), ClientOptions)

	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected Successfully")

	Collection = client.Database(Dbname).Collection(Collname)

	fmt.Println("Collection Instance is ready")

}

// mongodb Helper
func insertonemovie(movie model.Netflix) {
	insertedData, err := Collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Data with id : ", insertedData.InsertedID)
}

func updateonerecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"Watched": true}}

	updatedvalue, err := Collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		panic(err)
	}

	fmt.Println("Value Updated in go helper method", updatedvalue.ModifiedCount)
}

func deleteonerecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"id": id}

	deletecount, err := Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	fmt.Println("movie deleted", deletecount)

}

func deleteallfromdb() int64 {
	filter := bson.D{{}}
	deletedallcoount, err := Collection.DeleteMany(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	fmt.Println("All Documents are deleted from the backend", deletedallcoount.DeletedCount)

	return deletedallcoount.DeletedCount
}

func getallmovies() []primitive.M {
	cur, err := Collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)

		if err != nil {
			panic(err)
		}

		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())

	return movies
}

// MONGODB controllers

func AddoneMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding one Movie")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertonemovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func GetallMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all movies")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	movies := getallmovies()
	json.NewEncoder(w).Encode(movies)
}

func UpdateoneMovietomarkwatch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating Movies watchlist value")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateonerecord(params["id"])
	json.NewEncoder(w).Encode("Movie watch list Updated")
}

func DeleteOnemovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one movie")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)

	deleteonerecord(params["id"])
	json.NewEncoder(w).Encode("One Movie Deleted")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all movie")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	num := deleteallfromdb()
	json.NewEncoder(w).Encode(num)
}
