package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "PicturePerfect"
)

func main() {
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
	sqlStatement2 := `DROP TABLE moviesbasic`
	_, err = db.Exec(sqlStatement2)
	if err != nil {
		panic(err)
	}
	sqlStatement := `CREATE TABLE moviesbasic (
		MovieID SERIAL,
		Title TEXT,
		Duration INT,
		ReleaseDate DATE,
		Genre TEXT[],
		Rated TEXT,
		VoteTotal INT,
		VoteCount INT,
		Language TEXT,
		Poster TEXT,
		Backdrop TEXT,
		IMDbID TEXT PRIMARY KEY,
		TMDbID INT,
		Overview TEXT,
		Popularity FLOAT
	)`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("Added Table successfully")
}
