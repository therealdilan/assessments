package main

import (
	"fmt"
	"slices"
	"strconv"
)

var close = true

func main() {

	var sli = make([]int, 3)
	var input string

	for close == true {
		fmt.Println("Enter an Integer, type X to exit: ")
		fmt.Scanln(&input)

		if input == "X" || input == "x" {
			close = false
			break
		}

		val, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid Input!")
		} else {
			sli = append(sli, val)
			slices.Sort(sli)
			fmt.Println(sli)
		}

	}

}
