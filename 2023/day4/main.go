package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file, err := os.Open("/Users/jonathanwalter/dev/Advent-of-code/2023/day4/input.txt")
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
	total := 0

	for scanner.Scan() {
		var cardTotal int
		line := scanner.Text()

		line = line[7:]

		pair := strings.Split(line, "|")

		winningNumbers := toMap(strings.Fields(pair[0]))
		gottenNumbers := toMap(strings.Fields(pair[1]))

		for number, _ := range gottenNumbers {
			_, ok := winningNumbers[number]

			if ok {
				if cardTotal == 0 {
					cardTotal++
				} else {
					cardTotal = cardTotal * 2
				}
			}
		}

		total += cardTotal
	}

	log.Println(total)
}

func partTwo(scanner *bufio.Scanner) {
	cardCount := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		cardNumber := getCardNumber(line)
		line = line[7:]
		pair := strings.Split(line, "|")

		//add original card count
		cardCount[cardNumber]++

		currentCardCount := cardCount[cardNumber]

		for i := 0; i < currentCardCount; i++ {
			nextCard := cardNumber + 1
			winningNumbers := toMap(strings.Fields(pair[0]))
			gottenNumbers := toMap(strings.Fields(pair[1]))

			_, winnerCount := scoreCard(winningNumbers, gottenNumbers)

			for x := 0; x < winnerCount; x++ {
				cardCount[nextCard]++
				nextCard++
			}

		}
	}
	log.Println(totalValue(cardCount))
}

func totalValue(cards map[int]int) int {
	var total int
	for _, val := range cards {
		total += val
	}
	return total
}

func scoreCard(winningNumbers map[string]bool, gottenNumbers map[string]bool) (int, int) {
	var cardTotal int
	var hits int
	for number, _ := range gottenNumbers {
		_, ok := winningNumbers[number]

		if ok {
			if cardTotal == 0 {
				cardTotal++
				hits++
			} else {
				cardTotal = cardTotal * 2
				hits++
			}
		}
	}
	return cardTotal, hits
}

func toMap(numbers []string) map[string]bool {
	numberMap := map[string]bool{}

	for _, number := range numbers {
		numberMap[number] = true
	}

	return numberMap
}

func getCardNumber(fullLine string) int {
	cardNumber := strings.Fields(strings.Split(fullLine, ":")[0])[1]
	number, err := strconv.Atoi(cardNumber)
	if err != nil {
		panic(err)
	}
	return number
}
