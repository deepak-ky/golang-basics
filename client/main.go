package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "http://localhost:8080?name=Deepak"

	// Make the GET request
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in making GET Request: " ,err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error in reading body due to : ", err)
		return
	}


	fmt.Println(string(body))
}
