package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const numberOneConnections = 2
const numberFourConnections = 4
const numberSevenConnections = 3
const numberEightConnections = 7

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var outputs []string
	//build from input
	for scanner.Scan() {
		row := scanner.Text()

		//inputs := strings.Split(row, "|")[0]
		outputsString := strings.Split(row, "|")[1]
		outputs = append(outputs, strings.Split(outputsString, " ")...)
	}

	//	fmt.Printf("%v", outputs)

	var count int
	for _, output := range outputs {
		if len(output) == numberOneConnections || len(output) == numberFourConnections || len(output) == numberSevenConnections || len(output) == numberEightConnections {
			count++
		}
	}

	println(count)
}
