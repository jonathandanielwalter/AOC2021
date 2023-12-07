package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var cardValues = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardValuesWithJoker = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type kind int

var FIVE_OF_A_KIND kind = 1
var FOUR_OF_A_KIND kind = 2
var FULL_HOUSE kind = 3
var THREE_OF_A_KIND kind = 4
var TWO_PAIR kind = 5
var ONE_PAIR kind = 6
var HIGH_CARD kind = 7

func main() {
	file, err := os.Open("/Users/jonathanwalter/dev/Advent-of-code/2023/day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partTwo(scanner)

}

type hand struct {
	hand     string
	bid      int
	handType kind
}

func partOne(scanner *bufio.Scanner) {
	var total int
	var hands []hand
	for scanner.Scan() {
		line := scanner.Text()

		cardsAndBid := strings.Split(line, " ")
		bid, _ := strconv.Atoi(cardsAndBid[1])

		hands = append(hands, hand{
			hand:     cardsAndBid[0],
			bid:      bid,
			handType: determineKind(cardsAndBid[0]),
			//totalValue: getTotalValue(cardsAndBid[0]),
		})

	}

	for i := 0; i < len(hands)-1; i++ {
		for j := 0; j < len(hands)-i-1; j++ {
			if hands[j].handType < hands[j+1].handType {
				hands[j], hands[j+1] = hands[j+1], hands[j]
			} else if hands[j].handType == hands[j+1].handType {
				nextHigher := isNextHandHigher(hands[j], hands[j+1])
				if !nextHigher {
					hands[j], hands[j+1] = hands[j+1], hands[j]
				}
			}
		}
	}

	for i, hand := range hands {
		rank := i + 1
		total += hand.bid * rank
		log.Println(total)
	}
	log.Println(total)
}

func partTwo(scanner *bufio.Scanner) {
	var total int
	var hands []hand
	for scanner.Scan() {
		line := scanner.Text()

		cardsAndBid := strings.Split(line, " ")
		bid, _ := strconv.Atoi(cardsAndBid[1])

		hands = append(hands, hand{
			hand:     cardsAndBid[0],
			bid:      bid,
			handType: determineKindWithJoker(cardsAndBid[0]),
		})

	}

	for i := 0; i < len(hands)-1; i++ {
		for j := 0; j < len(hands)-i-1; j++ {
			if hands[j].handType < hands[j+1].handType {
				hands[j], hands[j+1] = hands[j+1], hands[j]
			} else if hands[j].handType == hands[j+1].handType {
				nextHigher := isNextHandHigherWithJoker(hands[j], hands[j+1])
				if !nextHigher {
					hands[j], hands[j+1] = hands[j+1], hands[j]
				}
			}
		}
	}

	for i, hand := range hands {
		rank := i + 1
		total += hand.bid * rank
		log.Println(total)
	}
	log.Printf("%v", hands)
	log.Println(total)
}

func isNextHandHigher(currentHand hand, nextHand hand) bool {
	for i := 0; i < len(currentHand.hand); i++ {
		currentCardValue := cardValues[rune(currentHand.hand[i])]
		nextCardValue := cardValues[rune(nextHand.hand[i])]
		if currentCardValue != nextCardValue {
			if currentCardValue < nextCardValue {
				return true
			} else {
				return false
			}
		}
	}

	return false
}
func isNextHandHigherWithJoker(currentHand hand, nextHand hand) bool {
	for i := 0; i < len(currentHand.hand); i++ {
		currentCardValue := cardValuesWithJoker[rune(currentHand.hand[i])]
		nextCardValue := cardValuesWithJoker[rune(nextHand.hand[i])]
		if currentCardValue != nextCardValue {
			if currentCardValue < nextCardValue {
				return true
			} else {
				return false
			}
		}
	}

	return false
}

func determineKind(hand string) kind {
	handCount := map[rune]int{}
	for _, card := range hand {
		handCount[card]++
	}

	if len(handCount) == 5 {
		return HIGH_CARD
	}
	if len(handCount) == 4 {
		return ONE_PAIR
	}
	if len(handCount) == 3 {
		for _, v := range handCount {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	}
	if len(handCount) == 2 {
		for _, v := range handCount {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	}
	if len(handCount) == 1 {
		return FIVE_OF_A_KIND
	}

	return 0
}

func determineKindWithJoker(hand string) kind {
	handCount := map[rune]int{}
	jokerCount := 0
	for _, card := range hand {
		if card == 'J' {
			jokerCount++
		}
	}

	if jokerCount == 5 {
		return FIVE_OF_A_KIND
	}

	for _, card := range hand {
		if card != 'J' {
			handCount[card]++
		}
	}
	if len(handCount) == 1 {
		return FIVE_OF_A_KIND
	}
	if len(handCount) == 2 {
		for _, v := range handCount {
			if v == 4-jokerCount {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	}
	if len(handCount) == 3 {
		for _, v := range handCount {
			if v == 3-jokerCount {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	}

	if len(handCount) == 4 {
		if jokerCount > 0 {
			for _, v := range handCount {
				if v == 3-jokerCount {
					return THREE_OF_A_KIND
				}
			}
		}
		return ONE_PAIR
	}
	if len(handCount) == 5 {
		return HIGH_CARD
	}

	return 0
}
