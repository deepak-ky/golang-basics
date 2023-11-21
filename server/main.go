package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootHandler)

	fmt.Println("Listening and Serving 8080")
	http.ListenAndServe(":8080", nil)

}

type RootApiResponse struct {
	Greeting string `json:"api-greeting"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "default"
	}

	fmt.Println("Received request for ", name)

	rootApiResponse := RootApiResponse{
		Greeting: fmt.Sprintf("Hi %s",name),
	}

	marshalledData, err := json.Marshal(rootApiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applicatoin/json")
	w.Write(marshalledData)
}
