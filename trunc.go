package main

import (
	"fmt"
)

func main() {
	var userInput float64
	fmt.Println("Please enter a random decimal number to truncate:")
	fmt.Scan(&userInput)
	fmt.Println("The truncated number is", int(userInput))
}
