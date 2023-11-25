package main

import (
	"encoding/json"
	"fmt"
)

var name string
var address string

func main() {
	fmt.Println("Enter your name:")
	fmt.Scanln(&name) // Get the name

	fmt.Println("Enter your address:")
	fmt.Scanln(&address) // Get the address

	data := map[string]interface{}{
		"name":    name,
		"address": address,
	} // Create a data map

	fmt.Println(data)

	jsonFile, err := json.Marshal(data) // Turn into json file
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Your json file was made:", string(jsonFile)) // Print the json file
}
