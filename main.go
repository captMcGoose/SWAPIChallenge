package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sort"
    "github.com/gorilla/mux"
)

// Person structure to hold individual character data
type Person struct {
    Name string `json:"name"`
}

// Response structure to hold the response from SWAPI
type Response struct {
    Results []Person `json:"results"`
}

// This is the function that gets executed when we get a GET request on the /people endpoint on our service:
// It will fetch the data from the SWAPI's /people endpoint, transform it, and send it through. 
func getSortedPeople(w http.ResponseWriter, r *http.Request) {
    // Fetch data from the SWAPI
    log.Println("Fetching data from SWAPI...")
    resp, err := http.Get("https://swapi.dev/api/people/")
    if err != nil {
        log.Println("Error fetching data from SWAPI:", err)
        http.Error(w, "Error fetching data from SWAPI", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Decode the response body
    var response Response
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        log.Println("Error decoding SWAPI response:", err)
        http.Error(w, "Error decoding SWAPI response", http.StatusInternalServerError)
        return
    }

    // Transform the data. Sort the results by the "name" field
    log.Println("Transforming data...")
    sort.SliceStable(response.Results, func(i, j int) bool {
        return response.Results[i].Name < response.Results[j].Name
    })

    // Send the sorted results as JSON via our own /people endpoint
    log.Println("Sending data to the endpoint page...")
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response.Results); err != nil {
        log.Println("Error encoding response:", err)
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }
}

func main() {
    // Initialize router
    r := mux.NewRouter()

    // Define our /people endpoint
    r.HandleFunc("/people", getSortedPeople).Methods("GET")

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Error starting server:", err)
    }
}
