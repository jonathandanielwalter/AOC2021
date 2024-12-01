package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	run(scanner)
}

func run(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		string := strings.Split(line[0], "")
		nums := turnStringsToInts(strings.Split(line[1], ","))

		// log.Println(vals, groups)

		// var combos int

		for i, _ := range nums {
			
			


		}

	}
}

func turnStringsToInts(strs []string) []int {
	var ints []int

	for _, str := range strs {
		val, _ := strconv.Atoi(str)
		ints = append(ints, val)
	}

	return ints
}
