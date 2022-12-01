package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var largestCountSoFar int
	var currentCalorieCount int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentCalorieCount > largestCountSoFar {
				largestCountSoFar = currentCalorieCount
			}
			currentCalorieCount = 0
			continue
		}

		calories, err := strconv.Atoi(line)

		if err != nil {
			log.Println("error reading file: ", err)
		}

		currentCalorieCount += calories

	}

	log.Println(largestCountSoFar)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfCalories := []int{}

	var currentElfCalories int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elfCalories = append(elfCalories, currentElfCalories)
			currentElfCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)

		if err != nil {
			log.Println("error reading file: ", err)
		}

		currentElfCalories += calories

	}

	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] > elfCalories[j]
	})

	topThree := elfCalories[0] + elfCalories[1] + elfCalories[2]

	log.Println(topThree)

}
