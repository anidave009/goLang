package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Movie struct {
	Title    string
	Year     string
	Poster   string
	Response string
	Error    string
}

func main() {
	// 1. Check if user provided a movie name
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie name")
		return
	}

	//extract moviename from args
	movieName := os.Args[1]

	apiKey := os.Getenv("OMDB_API_KEY")

	apiURL := "https://www.omdbapi.com/?apikey=" + apiKey +
		"&t=" + url.QueryEscape(movieName)

	// fmt.Println("Request URL:", apiURL)

	//make a request to the opemmoviedb api
	response, err := http.Get(apiURL)

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	// Always close the response body when we're done
	defer response.Body.Close()

	// Check the HTTP status
	fmt.Println("Response status:", response.Status)

	//parse the JSON response into a Movie struct
	var movie Movie

	err = json.NewDecoder(response.Body).Decode(&movie)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if movie.Response == "False" {
		fmt.Println("Movie not found:", movie.Error)
		return
	}

	fmt.Println("Title:", movie.Title)
	fmt.Println("Year:", movie.Year)
	fmt.Println("Poster URL:", movie.Poster)
	fmt.Println("Response:", movie.Response)

	//make a http request to download the poster image
	posterResponse, err := http.Get(movie.Poster)
	if err != nil {
		fmt.Println("Error downloading poster:", err)
		return
	}
	defer posterResponse.Body.Close()

	//create a file to save the poster image
	file, err := os.Create("poster.jpg")

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	//copy the poster image data to the file
	_, err = io.Copy(file, posterResponse.Body)

	if err != nil {
		fmt.Println("Error saving poster:", err)
		return
	}
	fmt.Println("Poster downloaded successfully!")
}

// type Geo struct {
// 	Lat string
// 	Lng string
// }

// type Address struct {
// 	Street  string
// 	Suite   string
// 	City    string
// 	Zipcode string
// 	Geo     Geo
// }

// type Company struct {
// 	Name        string
// 	CatchPhrase string
// 	Bs          string
// }

// type User struct {
// 	ID       int
// 	Name     string
// 	Username string
// 	Email    string
// 	Address  Address
// 	Phone    string
// 	Website  string
// 	Company  Company
// }

// url := "https://jsonplaceholder.typicode.com/users/1"

// response, err := http.Get(url)

// if err != nil {
// 	fmt.Println("Error", err)
// 	return
// }

// defer response.Body.Close()

// var user User

// err = json.NewDecoder(response.Body).Decode(&user)

// if err != nil {
// 	fmt.Println("Error reading response:", err)
// 	return
// }

// fmt.Println(user)
// fmt.Println(user.Name)
// fmt.Println(user.Address.City)
// fmt.Println(user.Address.Geo.Lat)
// fmt.Println(user.Company.Name)

// fmt.Println(response.Status)
// data, err := io.ReadAll(response.Body)

// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return
// 	}

// 	var prettyJSON bytes.Buffer

// 	err = json.Indent(&prettyJSON, data, "", "  ")

// 	if err != nil {
// 		fmt.Println("Error formatting JSON:", err)
// 		return
// 	}

// 	fmt.Println(prettyJSON.String())

// $env:OMDB_API_KEY="your-new-api-key" // set it in gopractise module.
// echo $env:OMDB_API_KEY

//change directory and come to module2/OpenMovieDb and run the code.
// go run openMovieDb.go "The Matrix" // run the code with movie name as argument.
