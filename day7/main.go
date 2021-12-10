package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var numbers []string
	scanner := bufio.NewScanner(file)
	var largestNumber int

	//build from input
	for scanner.Scan() {
		row := scanner.Text()
		numbers = strings.Split(row, ",")

		for _, number := range numbers {
			current, _ := strconv.Atoi(number)
			if current > largestNumber {
				largestNumber = current
			}
		}
	}

	crabs := []int{}
	for _, numberSrt := range numbers {
		number, _ := strconv.Atoi(numberSrt)
		crabs = append(crabs, number)
	}

	var leastFuelUsed int
	//var indexForLeastFuel int

	for i := 0; i < largestNumber; i++ {
		currentTotalFuelCost := 0

		//fmt.Printf("%v", crabs)
		for _, crabPos := range crabs {
			numberOfMoves := Abs(i - crabPos)
			currentTotalFuelCost += calculateTotalFuelUsage(numberOfMoves)
			//fmt.Printf("total fuel cost for place %v is %v\n", i, currentTotalFuelCost)
		}
		if leastFuelUsed == 0 {
			leastFuelUsed = currentTotalFuelCost
		} else {
			if currentTotalFuelCost < leastFuelUsed {
				//	println("least used fuel is ", indexForLeastFuel)
				//indexForLeastFuel = i
				leastFuelUsed = currentTotalFuelCost
			}
		}
	}

	fmt.Println(leastFuelUsed)
}

func calculateTotalFuelUsage(numberOfMoves int) int {
	var totalUsage int

	for i := 1; i <= numberOfMoves; i++ {
		totalUsage = totalUsage + i
	}
	return totalUsage
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
