package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var xRegister int = 1
var currentCycle = 0
var cycles = []int{220, 180, 140, 100, 60, 20}
var row = []string{}
var total int

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\r\n") // running on windows

	for _, input := range inputs {
		run(input)
	}

	log.Println("part one total:", total)
}

func run(input string) {
	if input == "noop" {

		cycles[xRegister]

		passNextImportantCycle(1)
		currentCycle++
	} else {
		value, _ := strconv.Atoi(strings.Fields(input)[1])

		passNextImportantCycle(2)
		currentCycle += 2
		xRegister += value

	}
}

func getSignalStrength(runtimeCycle int) int {
	log.Println("register is:", xRegister)
	return runtimeCycle * xRegister
}

func passNextImportantCycle(increment int) {

	if len(cycles) == 0 {
		return
	}
	neededCycleNumber := cycles[len(cycles)-1]
	if currentCycle+increment >= neededCycleNumber {

		log.Println("current cycle:",
			currentCycle)

		// log.Println(getSignalStrength(neededCycleNumber))

		total += getSignalStrength(neededCycleNumber)

		cycles = cycles[:len(cycles)-1]
		log.Println(cycles)
	}
}
