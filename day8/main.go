package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const numberZeroConnections = 6
const numberOneConnections = 2
const numberTwoConnection = 5
const numberFourConnections = 4
const numberSevenConnections = 3
const numberEightConnections = 7

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var inputs [][]string
	var outputs []string
	//build from input
	for scanner.Scan() {
		row := scanner.Text()

		inputString := strings.Split(row, "|")[0]
		inputs = append(inputs, strings.Split(inputString, " "))
		//inputs := strings.Split(row, "|")[0]
		outputsString := strings.Split(row, "|")[1]
		outputs = append(outputs, strings.Split(outputsString, " ")...)
	}

	//fmt.Printf("%v", inputs)

	// var count int
	// for _, output := range outputs {
	// 	if len(output) == numberOneConnections || len(output) == numberFourConnections || len(output) == numberSevenConnections || len(output) == numberEightConnections {
	// 		count++
	// 	}
	// }

	// println(count)

	for _, input := range inputs {
		deduceLeterValuesForInputs(input)
	}
}

func deduceLeterValuesForInputs(inputs []string) {

	//2 letters means number 1 (either position 3 or 6 for each letter)

	candidatesForPosition := make(map[int][]string)

	// allLettersDeduced bool = false

	for _, input := range inputs {
		if len(input) == 2 { //if its 1 specifically then we add the nunbers to potentials for 3 and 6
			for _, char := range input {

				addInputToMap(&candidatesForPosition, 1, char)
			}
		}
	}

	fmt.Printf("%v", candidatesForPosition)
}

func addInputToMap(candidates *map[int][]string, number int, inputChar rune) {
	myMap := *candidates

	if contains(myMap[number], inputChar) {
		//do nothing
	} else {
		myMap[number] = append(myMap[number], strconv.QuoteRune(inputChar))
	}
}

func contains(list []string, input rune) bool {
	for _, str := range list {
		if str == strconv.QuoteRune(input) {
			return true
		}
	}

	return false
}
