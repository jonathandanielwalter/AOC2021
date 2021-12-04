package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	allBinaryNumbers := [][]string{}

	var bitLength int

	for scanner.Scan() {
		number := scanner.Text()
		bitLength = len(number)
		var binaryNubers []string = strings.Split(number, "")
		allBinaryNumbers = append(allBinaryNumbers, binaryNubers)
	}

	oxyRating, _ := strconv.ParseInt(strings.Join(getOxygenRating(allBinaryNumbers, bitLength), ""), 2, 64)
	co2Rating, _ := strconv.ParseInt(strings.Join(getC02Rating(allBinaryNumbers, bitLength), ""), 2, 64)

	println(oxyRating * co2Rating)
}

func partOne() {
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[int][]string)

	for scanner.Scan() {
		number := scanner.Text()
		var binaryNubers []string = strings.Split(number, "")

		for i, s := range binaryNubers {
			m[i] = append(m[i], s)
		}

	}
	gammaValue := calcGamma(m)
	epsilonValue := flipBinary(gammaValue)

	gammaInt, err := strconv.ParseInt(gammaValue, 2, 64)
	epsilonInt, err := strconv.ParseInt(epsilonValue, 2, 64)

	println("gamma Int ", gammaInt)
	println("epsilonInt ", epsilonInt)

	print(gammaInt * epsilonInt)
}

func calcGamma(m map[int][]string) string {
	var gammaNumber string

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	for key := range keys {
		var count1 int
		var count0 int

		for _, number := range m[key] {
			if number == "1" {
				count1++
			} else {
				count0++
			}
		}

		if count1 > count0 {
			gammaNumber = gammaNumber + "1"
		} else {
			gammaNumber = gammaNumber + "0"
		}
	}
	return gammaNumber
}

func flipBinary(binary string) string {
	var newNumber string

	var binaryNubers []string = strings.Split(binary, "")
	for _, s := range binaryNubers {
		if s == "1" {
			newNumber = newNumber + "0"
		} else {
			newNumber = newNumber + "1"
		}
	}
	return newNumber
}

func getOxygenRating(allBinaryNumbers [][]string, bitLength int) []string {
	for i := 0; i < bitLength-1; i++ {
		if len(allBinaryNumbers) > 1 {
			mostCommon := getMostCommonAtIndex(allBinaryNumbers, i)
			allBinaryNumbers = removeValuesWithCommonValueAtIndex(allBinaryNumbers, mostCommon, i)
		}
	}
	return allBinaryNumbers[0]
}

func getC02Rating(allBinaryNumbers [][]string, bitLength int) []string {
	for i := 0; i < bitLength-1; i++ {
		if len(allBinaryNumbers) > 1 {
			leastCommon := getLeastCommonAtIndex(allBinaryNumbers, i)
			allBinaryNumbers = removeValuesWithCommonValueAtIndex(allBinaryNumbers, leastCommon, i)

		}

	}
	return allBinaryNumbers[0]
}

func getMostCommonAtIndex(binaryNumbers [][]string, index int) string {
	var count1 int
	var count0 int

	for _, number := range binaryNumbers {
		//log.Printf("%v", number)
		//	println("searching at index ", index)
		if number[index] == "1" {
			//println(number[index])
			count1++
		} else {
			//println(number[index])
			count0++
		}
	}

	if count1 > count0 {
		//println("Most common bit was 1")
		return "1"
	} else if count0 == count1 {
		//println("Equal bit count, returning 1 ")
		return "1"
	}
	// println("Most common bit was 0")
	return "0"
}

func getLeastCommonAtIndex(binaryNumbers [][]string, index int) string {
	var count1 int
	var count0 int

	for _, number := range binaryNumbers {
		if number[index] == "0" {
			count0++
		} else {
			count1++
		}
	}

	if count1 < count0 {
		return "1"
	} else if count0 == count1 {
		return "0"
	}
	return "0"
}

func removeValuesWithCommonValueAtIndex(allBinaryNumbers [][]string, mostCommon string, index int) [][]string {
	//mutatedBinaryNumbers := allBinaryNumbers
	log.Printf("%v", allBinaryNumbers)
	for i, number := range allBinaryNumbers {
		println(i)
		if number[index] != mostCommon {
			allBinaryNumbers = removeIndex(allBinaryNumbers, i)
		}
	}
	return allBinaryNumbers
}

func removeIndex(allBinaryNumbers [][]string, index int) [][]string {
	allBinaryNumbers[index] = allBinaryNumbers[len(allBinaryNumbers)-1]
	return allBinaryNumbers[:len(allBinaryNumbers)-1]
}
