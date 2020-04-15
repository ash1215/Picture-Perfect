package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetMovieByTitle(title string) Movie {
	title = strings.Replace(title, " ", "+", -1)
	apiRequest := "http://www.omdbapi.com/?apikey=9bd55f8b&t=" + title
	response, err := http.Get(apiRequest)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var movie1 Movie
	json.Unmarshal(responseData, &movie1)
	return movie1
}

func GetMovieByID(titleID string) Movie {
	apiRequest := "http://www.omdbapi.com/?apikey=9bd55f8b&i=" + titleID
	response, err := http.Get(apiRequest)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var movie1 Movie
	json.Unmarshal(responseData, &movie1)
	return movie1
}
func ReadTsv(filename string, query string, titleType string, RequestedMovies []Movie) []Movie {

	// Convert query string to lowercase
	query = strings.ToLower(query)

	// Open TSV file
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	// Read File into a Variable
	tsvReader := csv.NewReader(f)
	for {
		line, err := tsvReader.Read()
		if err == io.EOF {
			break
		}
		if len(line) == 0 {
			continue
		}
		if strings.Contains(strings.ToLower(line[0]), query) && strings.Contains(strings.ToLower(line[0]), titleType) {
			titleID := line[0][0:9]
			if titleID != RequestedMovies[0].ID {
				mymovie := GetMovieByID(titleID)
				fmt.Println(mymovie)
				if mymovie.ID == titleID {
					fmt.Println(mymovie)
					fmt.Println("appended")
					RequestedMovies = append(RequestedMovies, mymovie)
				}
			}
		}

	}
	return RequestedMovies

}
