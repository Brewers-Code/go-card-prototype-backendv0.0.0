package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Planet is a planet type
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

// Person is a person type
type Person struct {
	Name        string `json:"name"`
	HomewordURL string `json:"homeworld"`
	Homeworld   Planet
}

func (p *Person) getHomeWorld() {
	res, err := http.Get(p.HomewordURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading res body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

// type AllPeople is a collection of Person types
type AllPeople struct {
	People []Person `json:"results"`
}

const PORT = ":8080"
const BaseURL = "https://swapi.dev/api/"

// -------------------
// Endpoint handlers
// -------------------
func getPeople(w http.ResponseWriter, r *http.Request) {
	// handle get req to the API endpoint
	// this will return the res and err
	res, err := http.Get(BaseURL + "people")

	// 1st handle err
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to req star wars people")
	}

	// parse http res:
	// parse Body of the http res
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse req body")
	}

	var people AllPeople
	// an arr of a set of integers. Parse the set of bytes to the rep we want
	// fmt.Println(string(bytes))

	// parse encoded data and store the result in AllPeople struct people
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	// Make another API call to the URL stored
	for _, persons := range people.People {
		// Fetch and adddata for each person HomewordURL
		persons.getHomeWorld()
		fmt.Println("Person", persons)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	// where are we writting, req that as passed through
	// these are built-in go types
	fmt.Println("Home!")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/people", getPeople)

	fmt.Printf("Server is running on http://localhost%s\n", PORT)
	// the method will get executed no matter what. log.Fatal is a backup wrapper
	log.Fatal(http.ListenAndServe(PORT, nil))
}
