package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//partOne(scanner)
	partTwo(scanner)
}

func partOne(scanner *bufio.Scanner) {
	var total int
	for scanner.Scan() {
		//var number string
		line := scanner.Text()

		regex := regexp.MustCompile("[^0-9]")
		strNumber := regex.ReplaceAllString(line, "")

		firstDigit := strNumber[0]
		lastDigit := strNumber[len(strNumber)-1]

		number, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
		if err != nil {
			panic(err)
		}
		total += number
	}

	log.Printf("total is %v", total)
}

func partTwo(scanner *bufio.Scanner) {
	var total int
	numberRegex := regexp.MustCompile("1|2|3|4|5|6|7|8|9")
	//stringRegex := regexp.MustCompile("[one|two|three|four|five|six|seven|eight|nine]")
	for scanner.Scan() {
		line := scanner.Text()
		var first string
		var last string
		var err error

		var firstSub string
		var lastSub string

		for _, val := range line {
			if numberRegex.MatchString(string(val)) {
				first = string(val)
				if err != nil {
					panic(err)
				}
				break
			} else {
				firstSub = firstSub + string(val)
				contains, number := containsNumber(firstSub)

				if contains {
					first = getNumbs(number)
					break

				}

			}
		}

		for x := len(line) - 1; x >= 0; x-- {
			val := line[x]
			if numberRegex.MatchString(string(val)) {
				last = string(val)
				if err != nil {
					panic(err)
				}
				break
			} else {
				lastSub = string(val) + lastSub
				contains, number := containsNumber(lastSub)

				if contains {
					last = getNumbs(number)
					break
				}

			}
		}

		strNumber := first + last
		log.Println(strNumber)

		number, err := strconv.Atoi(strNumber)
		if err != nil {
			panic(err)
		}
		total += number
	}

	log.Println(total)
}

func getNumbs(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	return "0"
}

func containsNumber(str string) (bool, string) {
	numbers := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for _, n := range numbers {
		if strings.Contains(str, n) {
			return true, n
		}
	}

	return false, ""
}
