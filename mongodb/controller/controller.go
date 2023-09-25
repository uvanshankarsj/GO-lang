package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main.go/model"
)

//const connection="mongodb+srv://uvan:J86VjPgRC98OFD3O@cluster0.5z3zed2.mongodb.net/?retryWrites=true&w=majority"
const connection="mongodb+srv://uvanshankar02:dzVRc4fqiyY7u5dw@cluster0.ejoioxu.mongodb.net/?retryWrites=true&w=majority"
//const connection = "mongodb+srv://devasunder64:Pass123word@cluster0.jmgkwf9.mongodb.net/?retryWrites=true&w=majority"
//const dbname = "Decision"
//const colname = "path"

const dbname = "netflix"
const colname = "Watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connection)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo dbconnectio success")
	collection = (*mongo.Collection)(client.Database(dbname).Collection(colname))
	fmt.Println("collection instance is ready")
}

func insertonemovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted 1 ovie in db successfully", inserted.InsertedID)
}

func updatemovie(MovieID string) {
	id, _ := primitive.ObjectIDFromHex(MovieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified successfully", result.ModifiedCount)
}

func deleteonemovie(MovieID string) {
	id, _ := primitive.ObjectIDFromHex(MovieID)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted successfully", result.DeletedCount)
}

func deleteallmovie() int {
	// id, _ := primitive.ObjectIDFromHex(MovieID)
	deleteresult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of moive deleted successfully", deleteresult.DeletedCount)
	return int(deleteresult.DeletedCount)
}

func getallmovies() []primitive.M {
	fmt.Println("getall movies")
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

func Getallmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	allmovies := getallmovies()
	json.NewEncoder(w).Encode(allmovies)
}

func Createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	fmt.Println("working")
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertonemovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func Markaswatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "POST")

	params := mux.Vars(r)
	updatemovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func Deleteonemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "DELETE")

	params := mux.Vars(r)
	deleteonemovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func Delteallmoveies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "DELETE")

	count := deleteallmovie()
	json.NewEncoder(w).Encode(count)
}
