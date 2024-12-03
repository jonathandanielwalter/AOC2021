package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(content)
	part1(input)
	part2(input)
}

func part1(input string) {
	r := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)
	matches := r.FindAllString(input, -1)

	//fmt.Println(matches)

	var total int

	for _, match := range matches {
		reduced := match[4 : len(match)-1]

		multipliers := strings.Split(reduced, ",")

		left, _ := strconv.Atoi(multipliers[0])
		right, _ := strconv.Atoi(multipliers[1])

		total += left * right

	}

	fmt.Println(total)
}

func part2(input string) {
	r := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	matches := r.FindAllString(input, -1)

	//fmt.Println(matches)

	var total int
	enabled := true
	for _, match := range matches {
		switch match {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				reduced := match[4 : len(match)-1]

				multipliers := strings.Split(reduced, ",")

				left, _ := strconv.Atoi(multipliers[0])
				right, _ := strconv.Atoi(multipliers[1])

				total += left * right
			}

		}

	}

	fmt.Println(total)
}
