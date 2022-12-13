package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	Items        []int
	Operation    string
	DivisibleBy  int
	TrueOutcome  int
	FalseOutcome int
}

func main() {
	monkeys := readInputs()
}

func readInputs() {
	monkeys := []Monkey{}

	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\n") // running on mac
	//Monkey 0:
	//Starting items: 54, 98, 50, 94, 69, 62, 53, 85
	//	Operation: new = old * 13
	//	Test: divisible by 3
	//If true: throw to monkey 2
	//If false: throw to monkey 1
	//
	var currentMonkey Monkey
	for i, input := range inputs {
		if strings.Contains(input, "monkey") {
			currentMonkey = Monkey{
				Items: []int{},
			}
			continue
		}
		if strings.Contains(input, "starting") {
			line := strings.Split(input, "Starting items: ")
			numberStrings := strings.Split(line[1], ", ")

			for _, item := range numberStrings {
				itemNumber, err := strconv.Atoi(item)
				if err != nil {
					log.Println(err)
				}
				currentMonkey.Items = append(currentMonkey.Items, itemNumber)
			}
		}
		components := strings.Fields(input)
		if strings.Contains(input, "Operation") {

		}

		if strings.Contains(input, "Test") {
			divider, err := strconv.Atoi((components[len(components)-1]))
			if err != nil {
				log.Println(err)
			}
			currentMonkey.DivisibleBy = divider
		}

		if strings.Contains(input, "true") {

		}

		if strings.Contains(input, "false") {

		}

		if input == " " {
			monkeys = append(monkeys, currentMonkey)
		}

	}

}
