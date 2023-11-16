package main

import (
	"fmt"
	"strings"
)

func main() {
	var ian string
	fmt.Print("Enter a string: ")
	fmt.Scan(&ian)
	ian = strings.ToLower(ian)
	I := strings.HasPrefix(ian, "i")
	N := strings.HasSuffix(ian, "n")
	A := strings.Contains(ian, "a")
	if I && N && A {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
