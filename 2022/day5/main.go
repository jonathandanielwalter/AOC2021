package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	moves := strings.Split(string(b), "\r\n")

	stacks := getStacks()

	partOne(moves, stacks)
	partTwo(moves, stacks)
}

func getStacks() []Stack {
	b, err := os.ReadFile("stacks.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\r\n")

	stacks := []Stack{}

	numberRow := inputs[len(inputs)-1]

	numberOfStacks := len(strings.Fields(numberRow))

	if err != nil {
		log.Println(err)
	}

	for i := 0; i < numberOfStacks; i++ {
		stacks = append(stacks, Stack{})
	}

	for numberIndex, number := range numberRow {

		if string(number) != " " {
			for i := len(inputs) - 2; i >= 0; i-- {
				stackNumber, _ := strconv.Atoi(string(number))

				value := string(inputs[i][numberIndex])

				if value != " " {
					stacks[stackNumber-1].Push(value)
				}

			}
		}

	}

	return stacks
}

func partOne(inputs []string, stacks []Stack) {

	for _, input := range inputs {
		splitPut := strings.Fields(input)

		amount, err := strconv.Atoi(splitPut[1])
		from, err := strconv.Atoi(splitPut[3])
		to, err := strconv.Atoi(splitPut[5])

		if err != nil {
			log.Println("error handling moves")
		}
		for i := 0; i < amount; i++ {
			box, exists := stacks[from-1].Pop()

			if !exists {
				log.Println("tried to pop an empty stack")
			}

			stacks[to-1].Push(box)

		}

	}

	getTopCrates(stacks)
}

func partTwo(inputs []string, stacks []Stack) {
	for _, input := range inputs {
		splitPut := strings.Fields(input)

		amount, err := strconv.Atoi(splitPut[1])
		from, err := strconv.Atoi(splitPut[3])
		to, err := strconv.Atoi(splitPut[5])

		if err != nil {
			log.Println("error handling moves")
		}
		if amount == 1 {
			box, exists := stacks[from-1].Pop()
			if !exists {
				log.Println("tried to pop an empty stack")
			}
			stacks[to-1].Push(box)
			log.Println(stacks)
		} else {
			holdingStack := Stack{}

			for i := 0; i < amount; i++ {
				box, exists := stacks[from-1].Pop()
				if !exists {
					log.Println("tried to pop an empty stack")
				}

				holdingStack.Push(box)
			}

			for !holdingStack.IsEmpty() {
				box, exists := holdingStack.Pop()
				if !exists {
					log.Println("error popping from holding stack")
				}

				stacks[to-1].Push(box)
			}
		}

	}

	getTopCrates(stacks)
}

func getTopCrates(stacks []Stack) {
	sb := strings.Builder{}

	for _, stack := range stacks {
		top, exists := stack.Pop()
		if !exists {
			log.Println("tried to pop an empty stack")
		}
		sb.WriteString(top)
	}

	log.Println(sb.String())
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
