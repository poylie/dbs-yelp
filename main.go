package main

import (
	// Import the gorilla/mux library we just installed
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/yelp", getBusinessesHandler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search?term=starbucks&location=Singapore", nil)
	req.Header.Set("Authorization", "Bearer VNBosgLfAisHadcgnVbBUUawb--mf61EurggtiBUDcdG-StSyzuwQ39yk0n5OGMIbKxcjFlhaHh8c0YNIrbUmv-UXRaHtYNqcEAkrQ8c-5UMXUerZj85DoRZqU3yWnYx")
	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintf(w, string(data))
	}

	fmt.Fprintf(w, "Hello World!")
}
