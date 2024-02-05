package main

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
