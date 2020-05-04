package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "PicturePerfect"
)

type Movie struct {
	MovieID     string   `json:"MovieID"`
	Title       string   `json:"Title"`
	Duration    int      `json:"Duration"`
	ReleaseDate string   `json:"ReleaseDate"`
	Genre       []string `json:"Genre"`
	Rated       string   `json:"Rated"`
	VoteTotal   int64    `json:"VoteTotal"`
	VoteCount   int64    `json:"VoteCount"`
	Trailer     string   `json:"Trailer"`
	Language    string   `json:"Language"`
	Poster      string   `json:"Poster"`
	Backdrop    string   `json:"Backdrop"`
	IMDbID      string   `json:"IMDbID"`
	TMDbID      int      `json:"TMDbID"`
	Overview    string   `json:"Overview"`
	Popularity  float32  `json:"Popularity"`
}

func main() {
	fmt.Println("Here I am back")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT * FROM moviesbasic WHERE title ILIKE 'ra.one'")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var movieinfo Movie
		var genre string
		err = rows.Scan(&movieinfo.MovieID, &movieinfo.Title, &movieinfo.Duration, &movieinfo.ReleaseDate, &genre, &movieinfo.Rated, &movieinfo.VoteTotal, &movieinfo.VoteCount, &movieinfo.Language, &movieinfo.Poster, &movieinfo.Backdrop, &movieinfo.IMDbID, &movieinfo.TMDbID, &movieinfo.Overview, &movieinfo.Popularity)
		if err != nil {
			fmt.Println(err)
		}
		genre = genre[1 : len(genre)-1]
		movieinfo.Genre = strings.Split(genre, ",")
		fmt.Println(movieinfo)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
