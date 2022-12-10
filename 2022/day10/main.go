package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var xRegister int = 1
var currentCycle = 0
var cycles = []int{220, 180, 140, 100, 60, 20}
var rows = [][]string{
	make([]string, 40),
	make([]string, 40),
	make([]string, 40),
	make([]string, 40),
	make([]string, 40),
	make([]string, 40),
}
var total int

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\n") // running on windows

	for _, input := range inputs {
		run(input)
	}

	log.Println("part one total:", total)

	var out strings.Builder
	for i, line := range rows {
		out.WriteString(strings.Join(line, ""))
		if i < 6-1 {
			out.WriteByte('\n')
		}
	}

	log.Println(out.String())
}

func run(input string) {
	//row, col := getPosition(currentCycle, 40)

	row, col := getPosition(currentCycle, 40)
	if math.Abs(float64(col-xRegister)) <= 1 {
		rows[row][col] = "#"
	} else {
		rows[row][col] = "."
	}

	if input == "noop" {
		passNextImportantCycle(1)

		currentCycle++
	} else {
		value, _ := strconv.Atoi(strings.Fields(input)[1])

		//log.Println("cursor at ", xRegister)

		passNextImportantCycle(2)
		currentCycle += 2

		xRegister += value

	}
}

func getSignalStrength(runtimeCycle int) int {
	//log.Println("register is:", xRegister)
	return runtimeCycle * xRegister
}

func passNextImportantCycle(increment int) {

	if len(cycles) == 0 {
		return
	}
	neededCycleNumber := cycles[len(cycles)-1]
	if currentCycle+increment >= neededCycleNumber {
		total += getSignalStrength(neededCycleNumber)

		cycles = cycles[:len(cycles)-1]
		//log.Println(cycles)
	}
}

func getPosition(currentCycle, cols int) (int, int) {
	col := (currentCycle - 1) % cols
	row := (currentCycle - 1) / cols
	return row, col
}
