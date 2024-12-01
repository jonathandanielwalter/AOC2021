package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/jonat/playground/Advent-of-code/2024/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1, part2 := partOne(*scanner)
	fmt.Println(part1)
	fmt.Println(part2)
}

func partOne(scanner bufio.Scanner) (int, int) {
	var leftList []int
	var rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		leftValue, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		rightValue, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	sort.Sort(sort.IntSlice(leftList))
	sort.Sort(sort.IntSlice(rightList))

	var total int
	for i := 0; i < len(leftList); i++ {
		left := leftList[i]
		right := rightList[i]

		if left > right {
			total += left - right
		} else {
			total += right - left
		}
	}

	rightMap := map[int]int{}
	for _, val := range rightList {
		if _, ok := rightMap[val]; ok {
			rightMap[val] = rightMap[val] + 1
		} else {
			rightMap[val] = 1
		}
	}

	var multiplier int

	for _, val := range leftList {
		if _, ok := rightMap[val]; ok {
			multiplier += val * rightMap[val]
		}
	}

	return total, multiplier
}
