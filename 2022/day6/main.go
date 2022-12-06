package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\r\n")

	solve(inputs, 4)
	solve(inputs, 14)
}

func solve(inputs []string, blockSize int) {

	for _, input := range inputs {

		windowStart := 0
		windowEnd := blockSize

		var uniqueChars bool

		for windowEnd < len(input) {
			uniqueChars = containsUniqueChars(input[windowStart:windowEnd])

			if uniqueChars {
				log.Println(windowEnd)
				break
			} else {
				windowStart++
				windowEnd++

			}
		}
	}

}

func containsUniqueChars(str string) bool {

	theMap := map[rune]int{}
	for _, c := range str {
		if _, ok := theMap[c]; ok {
			return false
		} else {
			theMap[c] = 1
		}
	}

	return true
}
