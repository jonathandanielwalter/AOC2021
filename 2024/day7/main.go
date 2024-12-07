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
	file, err := os.Open("C:\\Users\\jonat\\playground\\Advent-of-code\\2024\\day7\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int
	var partTwoTotal uint64

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")

		result, _ := strconv.Atoi(split[0])
		numbersStr := split[1]

		var numbers []int

		for _, str := range strings.Fields(numbersStr) {
			n, _ := strconv.Atoi(str)
			numbers = append(numbers, n)
		}

		if calibrate(result, numbers[0], numbers, 1) {
			total += result
		}

		var numbersUint []uint64
		for _, num := range numbers {
			numbersUint = append(numbersUint, uint64(num))
		}

		if calibratepart2(uint64(result), numbersUint[0], numbersUint, 1) {
			partTwoTotal += uint64(result)
		}

	}

	fmt.Println(total)
	fmt.Println(partTwoTotal)
}

func calibrate(result int, currentTotal int, numbers []int, currentIndex int) bool {

	if currentIndex == len(numbers) {
		return result == currentTotal
	}

	return calibrate(result, currentTotal+numbers[currentIndex], numbers, currentIndex+1) || calibrate(result, currentTotal*numbers[currentIndex], numbers, currentIndex+1)
}

func calibratepart2(result uint64, currentTotal uint64, numbers []uint64, currentIndex int) bool {

	if currentIndex == len(numbers) {
		return result == currentTotal
	}

	return calibratepart2(result, currentTotal+numbers[currentIndex], numbers, currentIndex+1) ||
		calibratepart2(result, currentTotal*numbers[currentIndex], numbers, currentIndex+1) ||
		calibratepart2(result, concat(currentTotal, numbers[currentIndex]), numbers, currentIndex+1)
}

func concat(a, b uint64) uint64 {
	aStr := strconv.Itoa(int(a))
	bStr := strconv.Itoa(int(b))

	i, err := strconv.ParseUint(aStr+bStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
