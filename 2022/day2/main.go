package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// A = ROCK
// B = PAPER
// C = SCISSORS

// X = ROCK
// Y = PAPER
// Z = SCISSORS

func main() {
	// partOne()
	part2()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalScore int

	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, " ")

		totalScore += winLoseOrDrawPart1(values)
	}

	log.Println(totalScore)
}

func winLoseOrDrawPart1(values []string) int {
	var winnerMap = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	var drawMap = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	var shapeValue = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	leftValue := values[0]
	rightValue := values[1]

	if winnerMap[leftValue] == rightValue {
		return shapeValue[rightValue] + 6
	}

	if drawMap[leftValue] == rightValue {
		return shapeValue[rightValue] + 3
	}

	return shapeValue[rightValue]
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalScore int

	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, " ")

		totalScore += winLoseOrDrawPart2(values)
	}

	log.Println(totalScore)
}

func winLoseOrDrawPart2(values []string) int {

	leftValue := values[0]
	rightValue := values[1]

	var winnerMap = map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	var lossMap = map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	var shapeValue = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	//Need to lose
	if rightValue == "X" {
		return shapeValue[lossMap[leftValue]]
	}

	//Need to draw
	if rightValue == "Y" {
		return shapeValue[leftValue] + 3
	}

	return shapeValue[winnerMap[leftValue]] + 6

}
