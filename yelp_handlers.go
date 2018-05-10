package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Response is the entire body of the YELP API Response
type Response struct {
	Business []Business `json:"businesses"`
	Total    int        `json:"total"`
}

// Business struct for each business
type Business struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

// Location holds address of business
type Location struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	City     string `json:"city"`
	ZipCode  string `json:"zip_code"`
	Country  string `json:"country"`
	State    string `json:"state"`
}

func getBusinessesHandler(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search?term=starbucks&location=Singapore", nil)
	req.Header.Set("Authorization", "Bearer VNBosgLfAisHadcgnVbBUUawb--mf61EurggtiBUDcdG-StSyzuwQ39yk0n5OGMIbKxcjFlhaHh8c0YNIrbUmv-UXRaHtYNqcEAkrQ8c-5UMXUerZj85DoRZqU3yWnYx")
	response, err := client.Do(req)
	var responseObject Response

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Printf(string(data))

		json.Unmarshal(data, &responseObject)
	}

	//Convert the "business" variable to json

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of businesses to the response
	// w.Write(businessListBytes)

	tmpl := template.Must(template.ParseFiles("./assets/layout.html"))

	tmpl.Execute(w, responseObject)
}
