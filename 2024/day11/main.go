package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("/Users/jonathanwalter/dev/Advent-of-code/2024/day11/input.txt")
	if err != nil {
		panic(err)
	}

	input := string(content)
	initialNumbers := strings.Fields(input)
	fmt.Println(part1(AtoiList(initialNumbers)))
}

func part1(inputs []uint64) int {
	for c := 0; c < 75; c++ {
		var amended []uint64
		for _, input := range inputs {
			switch {
			case input == 0:
				amended = append(amended, 1)
			case len(strconv.FormatUint(input, 10))%2 == 0:
				str := strconv.FormatUint(input, 10)
				a := str[:len(str)/2]
				b := str[len(str)/2:]
				ai, _ := strconv.ParseUint(a, 10, 64)
				bi, _ := strconv.ParseUint(b, 10, 64)
				amended = append(amended, ai)
				amended = append(amended, bi)

			default:
				amended = append(amended, input*2024)
			}

		}
		inputs = amended
	}

	return len(inputs)
}

func AtoiList(list []string) []uint64 {
	var ints []uint64

	for _, i := range list {
		i, _ := strconv.ParseUint(i, 10, 64)
		ints = append(ints, i)
	}
	return ints
}
