package Database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var mapGenre = make(map[int32]string)

func initMapGenre() {
	mapGenre[28] = "Action"
	mapGenre[12] = "Adventure"
	mapGenre[16] = "Animation"
	mapGenre[35] = "Comedy"
	mapGenre[80] = "Crime"
	mapGenre[99] = "Documentary"
	mapGenre[18] = "Drama"
	mapGenre[10751] = "Family"
	mapGenre[14] = "Fantasy"
	mapGenre[36] = "History"
	mapGenre[27] = "Horror"
	mapGenre[9648] = "Mystery"
	mapGenre[10402] = "Music"
	mapGenre[10749] = "Romance"
	mapGenre[878] = "Sci-Fi"
	mapGenre[10770] = "TV Movie"
	mapGenre[53] = "Thriller"
	mapGenre[10752] = "War"
	mapGenre[37] = "Western"
}

type MovieDetails struct {
	Adult         bool    `json:"adult"`
	BackdropPath  string  `json:"backdrop_path"`
	ID            int     `json:"id"`
	OriginalTitle string  `json:"original_title"`
	GenreIDs      []int32 `json:"genre_ids"`
	Popularity    float32 `json:"popularity"`
	PosterPath    string  `json:"poster_path"`
	ReleaseDate   string  `json:"release_date"`
	Title         string  `json:"title"`
	Overview      string  `json:"overview"`
	Video         bool    `json:"video"`
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     uint32  `json:"vote_count"`
	Language      string  `json:"original_language"`
}

type MovieSearch struct {
	Page         int            `json:"page"`
	TotalPages   int            `json:"total_pages"`
	TotalResults int            `json:"total_results"`
	Results      []MovieDetails `json:"results"`
}

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

var imageLink = "https://image.tmdb.org/t/p/w600_and_h900_bestv2"

func getIMDBid(TMDbID int) (string, int) {
	APIKey := "317433652d5b6b2fb88b19cd436ee5d6"
	SearchRequest := "https://api.themoviedb.org/3/movie/" + strconv.Itoa(TMDbID) + "?api_key=" + APIKey + "&language=en-US"
	response, err := http.Get(SearchRequest)

	if err != nil {
		fmt.Print(err.Error())
		return "", 0
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", 0
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(responseData), &result)
	id, ok := result["imdb_id"].(string)
	if !ok {
		return "", 0
	}
	runtime, _ := result["runtime"].(float64)
	return id, int(runtime)
}
func getRated(IMDbID string) string {
	apiRequest := "http://www.omdbapi.com/?apikey=9bd55f8b&i=" + IMDbID
	response, err := http.Get(apiRequest)

	if err != nil {
		fmt.Print(err.Error())
		return ""
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(responseData), &result)
	rated, ok := result["Rated"].(string)
	if !ok {
		return ""
	}
	return rated

}
func toRequiredJSON(movieInfo []MovieDetails) []Movie {
	var tempList []Movie
	for i := 0; i < len(movieInfo); i++ {
		temp := Movie{}
		temp.Title = movieInfo[i].Title
		temp.Poster = ""
		if movieInfo[i].PosterPath != "" {
			temp.Poster = imageLink + movieInfo[i].PosterPath
		}
		temp.Backdrop = ""
		if movieInfo[i].BackdropPath != "" {
			temp.Backdrop = imageLink + movieInfo[i].BackdropPath
		}
		temp.TMDbID = movieInfo[i].ID
		for j := 0; j < len(movieInfo[i].GenreIDs); j++ {
			temp.Genre = append(temp.Genre, mapGenre[movieInfo[i].GenreIDs[j]])
		}
		temp.VoteCount = 0
		temp.VoteTotal = 0
		temp.Language = movieInfo[i].Language
		if temp.Language == "" {
			continue
		}
		temp.Overview = movieInfo[i].Overview
		if temp.Overview == "" {
			continue
		}
		temp.ReleaseDate = movieInfo[i].ReleaseDate
		if temp.ReleaseDate == "" {
			continue
		}
		temp.Popularity = 0
		temp.IMDbID, temp.Duration = getIMDBid(temp.TMDbID)
		if temp.IMDbID == "" || !(temp.Duration > 0) {
			continue
		}
		temp.Rated = getRated(temp.IMDbID)
		if temp.Rated == "" {
			temp.Rated = "Not Rated"
		}
		tempList = append(tempList, temp)
	}
	return tempList
}

func insertMovies(movieList []Movie, db *sql.DB) {

	sqlStatement := ""
	id := 0
	for i := 0; i < len(movieList); i++ {

		sqlStatement = `
		INSERT INTO moviesbasic (title, releasedate, language, poster, backdrop, tmdbid, popularity, overview, genre, imdbid, duration, rated, votetotal, votecount)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 0, 0)
		RETURNING movieid`
		tempGenre := ""
		for j := 0; j < len(movieList[i].Genre); j++ {
			tempGenre += `"` + movieList[i].Genre[j] + `"` + ","
		}
		sz := len(tempGenre)
		if sz > 0 && tempGenre[sz-1] == ',' {
			tempGenre = tempGenre[:sz-1]
		}
		tempGenre = "{" + tempGenre + "}"

		err := db.QueryRow(sqlStatement, movieList[i].Title, movieList[i].ReleaseDate, movieList[i].Language, movieList[i].Poster, movieList[i].Backdrop, movieList[i].TMDbID, movieList[i].Popularity, movieList[i].Overview, tempGenre, movieList[i].IMDbID, movieList[i].Duration, movieList[i].Rated).Scan(&id)
		if err != nil {
			// panic(err)
			fmt.Println(err)
		}
		fmt.Println("New record ID is:", id)
	}
	fmt.Printf("Inserted %d rows\n", len(movieList))

}

func Result(db *sql.DB, query string) []Movie {
	fmt.Println("starting query", query)
	rows, err := db.Query("SELECT * FROM moviesbasic WHERE title ILIKE '%' || $1 || '%'", query)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	fmt.Println("rows: ", rows)
	resultList := []Movie{}
	for rows.Next() {
		var movieinfo Movie
		var genre string
		err = rows.Scan(&movieinfo.MovieID, &movieinfo.Title, &movieinfo.Duration, &movieinfo.ReleaseDate, &genre, &movieinfo.Rated, &movieinfo.VoteTotal, &movieinfo.VoteCount, &movieinfo.Language, &movieinfo.Poster, &movieinfo.Backdrop, &movieinfo.IMDbID, &movieinfo.TMDbID, &movieinfo.Overview, &movieinfo.Popularity)
		if err != nil {
			fmt.Println(err)
		}
		genre = genre[1 : len(genre)-1]
		movieinfo.Genre = strings.Split(genre, ",")
		resultList = append(resultList, movieinfo)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	return resultList
}

func SearchTitle(query string, db *sql.DB) []Movie {
	initMapGenre()
	APIKey := "317433652d5b6b2fb88b19cd436ee5d6"
	language := "en-US"
	query_ := url.QueryEscape(query)

	SearchRequest := "https://api.themoviedb.org/3/search/movie?api_key=" + APIKey + "&language=" + language + "&query=" + query_ + "&include_adult=false"
	response, err := http.Get(SearchRequest)

	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var moveiList MovieSearch
	json.Unmarshal(responseData, &moveiList)
	fmt.Println("Received Movie List")
	fmt.Println(len(moveiList.Results))
	moviesList := toRequiredJSON(moveiList.Results)
	insertMovies(moviesList, db)
	moviesList = Result(db, query)
	return moviesList

}
