package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type struct1 struct {
	Status string `json:"statusValue"`
}

func main() {  
	//status := "UP"
	s1 := struct1{Status: "UP"}
	jsonBytes, _ := json.Marshal(s1)
	fmt.Println(jsonBytes)
	fmt.Println(string(jsonBytes))

	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	jsonEncoder := json.NewEncoder(file)
	err = jsonEncoder.Encode(s1)
	if err != nil {

	}

}
