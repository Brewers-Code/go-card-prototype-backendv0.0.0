package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URL string = "mongodb://localhost:27017"
const PORT string = ":8080"

type CardData struct {
	Anime    string   `bson:"anime"`
	Fname    string   `bson:"fname"`
	Lname    string   `bson:"lname"`
	Position string   `bson:"position"`
	Quote    string   `bson:"quote"`
	Images   []string `bson:"images"`
}

type AllCardData struct {
	CardsData []CardData `bson:"results"`
}

func getCardsData(w http.ResponseWriter, r *http.Request) {
	// Create a new client and start monitoring the MongoDB server on N local port
	client, err := mongo.NewClient(options.Client().ApplyURI(URL))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mongo db connected succesfully")
	}

	collection := client.Database("haikyuuDB").Collection("haikyu")
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var data []bson.D
	if err = cur.All(context.Background(), &data); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(data)

}

func main() {
	// Setup CORS default settings
	mux := http.NewServeMux()
	mux.HandleFunc("/", getCardsData)
	handler := cors.Default().Handler(mux)

	fmt.Printf("Server is running on http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, handler))

}
