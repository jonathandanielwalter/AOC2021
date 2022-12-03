package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.Println("part one :", partOne())
	log.Println("part two :", partTwo())
}

func partOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalValue int = 0

	for scanner.Scan() {
		line := scanner.Text()

		firstHalfString := []rune(line[:len(line)/2])
		secondHalfString := []rune(line[len(line)/2:])

		var firstHalfMap = make(map[rune]int)
		var secondHalfMap = make(map[rune]int)

		for i := 0; i < len(line)/2; i++ {
			firstHalfMap[firstHalfString[i]]++
			secondHalfMap[secondHalfString[i]]++
		}

		for letter := range firstHalfMap {
			_, ok := secondHalfMap[letter]

			if ok {
				totalValue += getGiftValue(letter)
			}
		}
	}
	return totalValue
}

func partTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	currentGroup := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		currentGroup = append(currentGroup, line)

		if len(currentGroup) == 3 {
			total += findSharedGiftValue(currentGroup)
			currentGroup = currentGroup[:0]
		}

	}

	return total
}

func findSharedGiftValue(currentGroup []string) int {
	gifts := map[rune]int{}

	for _, gift := range currentGroup[0] {
		gifts[gift] = 1
	}

	for _, gift := range currentGroup[1] {
		_, ok := gifts[gift]
		if ok {
			gifts[gift] = 2
		}
	}

	for _, gift := range currentGroup[2] {
		count, ok := gifts[gift]
		if ok && count == 2 {
			gifts[gift] = 3
		}
	}

	for gift, count := range gifts {
		if count == 3 {
			return getGiftValue(gift)
		}
	}

	return 0
	// 	alreadySeenMap := map[rune]bool{}
	// 	for _, gift := range elfBag {

	// 		_, giftExists := gifts[gift]
	// 		_, alreadySeen := alreadySeenMap[gift]

	// 		if alreadySeen {
	// 			continue
	// 		}

	// 		if giftExists {
	// 			if gifts[gift] == 3 {
	// 				return getGiftValue(gift)
	// 			} else {
	// 				gifts[gift]++
	// 			}

	// 		} else {
	// 			gifts[gift] = 1
	// 			alreadySeenMap[gift] = true
	// 		}

	// 	}
	// }

	// return 0
}

func getGiftValue(gift rune) int {
	if gift > 'a' && gift <= 'z' {
		return int(gift - 'a' + 1)
	} else if gift > 'A' && gift <= 'Z' {
		return int(gift - 'A' + 27)
	}
	return 0
}
