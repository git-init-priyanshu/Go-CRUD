package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var movie Movie

	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			movie = item
			break
		}
	}

	json.NewEncoder(w).Encode(movie)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var movie Movie

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// Deleting movie
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	var id string = params["id"]

	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = id
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func main() {
	// Appending first movie
	movies = append(
		movies,
		Movie{
			ID:       "1",
			Isbn:     "438227",
			Title:    "Movie One",
			Director: &Director{Firstname: "John", Lastname: "Doe"},
		})

	r := mux.NewRouter()

	// Defining routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
