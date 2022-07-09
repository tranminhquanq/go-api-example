package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"golang-example/golang-movies/models"
)

var movies[]models.Movie

func main() {
	r := mux.NewRouter().StrictSlash(true)
	port := ":8080"
	if envP := os.Getenv("PORT"); envP != "" {
		port = envP
	}

	movies = append(movies, models.Movie{
		ID: "1",
		Isbn: "438227",
		Title: "Movie One",
		Director: &models.Director{
			LastName: "Minh Quang",
			FirstName: "Tran",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server running on PORT 8080")
	log.Fatal(http.ListenAndServe(port, r))

}

func getMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Response{
		Data: movies,
		Message: "successfully",
		Code: 0,
	})
}

func createMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var newMovie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(models.Response{
		Data: movies,
		Message: "successfully",
		Code: 0,
	})
}

func getMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(models.Response{
				Data: []models.Movie{movie},
				Message: "successfully",
				Code: 0,
			})
			break
		}
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			var newMovie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = params["id"]
			movies = append(movies, newMovie)
			json.NewEncoder(w).Encode(models.Response{
				Data: []models.Movie{newMovie},
				Message: "successfully",
				Code: 0,
			})
			break
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(models.Response{
				Data: []models.Movie{item},
				Message: "successfully",
				Code: 0,
			})
			break
		}
	}

}