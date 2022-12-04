package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\r\n")

	log.Println("Part One : ", partOne(inputs))
	log.Println("Part Two : ", partTwo(inputs))
}

func partOne(inputs []string) int {
	numberOfFullContains := 0

	for _, input := range inputs {
		assignments := strings.Split(input, ",")

		if fullyContains(assignments[0], assignments[1]) {
			numberOfFullContains++
		}
	}

	return numberOfFullContains
}

func partTwo(inputs []string) int {
	overlaps := 0

	for _, input := range inputs {
		assignments := strings.Split(input, ",")

		if overlap(assignments[0], assignments[1]) {
			overlaps++
		}
	}

	return overlaps
}

func fullyContains(leftAssignmentPair string, rightAssignmentPair string) bool {
	leftValuesStrings := strings.Split(leftAssignmentPair, "-")
	rightValuesStrings := strings.Split(rightAssignmentPair, "-")

	leftValues := convertToInts(leftValuesStrings)
	rightValues := convertToInts(rightValuesStrings)

	/// 4-6,3-7
	if leftValues[0] >= rightValues[0] && leftValues[1] <= rightValues[1] {
		return true
	}

	// 3-7,4-6
	if leftValues[0] <= rightValues[0] && leftValues[1] >= rightValues[1] {
		return true
	}

	return false
}

func overlap(leftAssignmentPair string, rightAssignmentPair string) bool {
	leftValuesStrings := strings.Split(leftAssignmentPair, "-")
	rightValuesStrings := strings.Split(rightAssignmentPair, "-")

	leftValues := convertToInts(leftValuesStrings)
	rightValues := convertToInts(rightValuesStrings)

	if leftValues[0] <= rightValues[1] && leftValues[1] >= rightValues[0] {
		return true
	}

	if leftValues[0] >= rightValues[1] && leftValues[1] <= rightValues[0] {
		return true
	}

	return false
}

func convertToInts(strs []string) []int {
	numbers := []int{}
	for _, s := range strs {
		number, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}

		numbers = append(numbers, number)
	}

	return numbers
}
