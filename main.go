package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// struct of type movie
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Every Movie has a Director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// when we send this data using Postman, we'll be using json names

var movies []Movie

// to get all the Movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// delete movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "Varanam Ayiram", Director: &Director{Firstname: "Gautham", Lastname: "Vasudev Menon"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "CIA", Director: &Director{Firstname: "Amal", Lastname: "Neerad"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting serve at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
