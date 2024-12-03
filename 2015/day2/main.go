package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(content)

	var floor int
	visitedBasement := false
	for i, c := range input {

		switch string(c) {
		case "(":
			floor++
		case ")":
			floor--
		}

		if !visitedBasement && floor == -1 {
			fmt.Println(i + 1)
			visitedBasement = true
		}
	}

	fmt.Println(floor)
}
