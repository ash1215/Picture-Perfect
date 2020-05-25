package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	Database "./Database"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var myRouter *mux.Router = mux.NewRouter().StrictSlash(true)
var db *sql.DB

func startDB() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345678"
		dbname   = "PicturePerfect"
	)
	fmt.Println("Connecting to the database")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnMovies(w http.ResponseWriter, r *http.Request) {
	var RequestedMovies []Database.Movie
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	title := string(reqBody)
	fmt.Println(title)
	// RequestedMovies = append(RequestedMovies, GetMovieByTitle(title))
	// RequestedMovies = ReadTsv("IMDb Database/title.basics.tsv", title, "movie", RequestedMovies)

	RequestedMovies = Database.SearchTitle(title, db)
	json.NewEncoder(w).Encode(RequestedMovies)
	fmt.Println("Sent requested")
}

func handleRequests() {
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api", returnMovies).Methods("POST")
	log.Fatal(http.ListenAndServe(":8500", handlers.CORS(header, methods, origins)(myRouter)))
}

func main() {
	db = startDB()
	defer db.Close()
	handleRequests()
}
