package main

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
