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

	//for _, input := range inputs {
	deduceLeterValuesForInputs(inputs[0])
	//}
}

func deduceLeterValuesForInputs(inputs []string) {

	//2 letters means number 1 (either position 3 or 6 for each letter)

	candidatesForPosition := make(map[int][]string)

	// allLettersDeduced bool = false

	//deal with 1
	for _, input := range inputs {
		if len(input) == 2 { //if its 1 specifically then we add the nunbers to potentials for 3 and 6
			for _, char := range input {

				addInputToMap(&candidatesForPosition, 3, char)
				addInputToMap(&candidatesForPosition, 6, char)
			}
		}
	}

	//deal with 7
	for _, input := range inputs {
		if len(input) == 3 { //if its 7 specifically then we know that there can only be one candidate for the position, whatever wasnt at 3/6
			for _, char := range input {
				if _, ok := candidatesForPosition[1]; !ok {
					addInputToMap(&candidatesForPosition, 1, char)
				}

			}
		}
	}

	//deal with 2, 3 9, 6, 5
	for _, input := range inputs {
		if len(input) == 5 { //if its 2 specifically then we can add the candidates for 4 5 and 7 but also confirm the number in 3 and 6
			for _, char := range input {

				addInputToMap(&candidatesForPosition, 4, char)
				addInputToMap(&candidatesForPosition, 5, char)
				addInputToMap(&candidatesForPosition, 7, char)

				if contains(candidatesForPosition[3], char) {
					//remove other characters at position 3 and remove that character from position 6

					arr := make([]string, 1)
					arr[0] = strconv.QuoteRune(char)

					candidatesForPosition[3] = arr

					if contains(candidatesForPosition[6], char) {
						candidatesForPosition[6] = allButThisChar(candidatesForPosition[6], char)
					}
				}
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

func allButThisChar(candidates []string, c rune) []string {
	newList := []string{}

	for _, candidate := range candidates {
		if !(candidate == strconv.QuoteRune(c)) {
			newList = append(newList, candidate)
		}
	}

	return newList
}
