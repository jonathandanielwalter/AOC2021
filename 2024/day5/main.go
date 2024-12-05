package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/jonathanwalter/dev/Advent-of-code/2024/day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	onlists := false
	var rules []string
	var lists []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			onlists = true
			continue
		}
		if !onlists {
			rules = append(rules, scanner.Text())
		}

		if onlists {
			lists = append(lists, scanner.Text())
		}
	}

	fmt.Println(part1(rules, lists))

}

func part1(rules []string, lists []string) any {

	ruleMap := map[string][]string{}

	var count int
	var wrongCount int
	for _, rule := range rules {
		split := strings.Split(rule, "|")
		comesBefore := split[0]
		comesAfter := split[1]

		if val, ok := ruleMap[comesBefore]; ok {
			ruleMap[comesBefore] = append(val, comesAfter)
		} else {
			ruleMap[comesBefore] = []string{comesAfter}
		}
	}

	for _, list := range lists {
		listAllowed := true
		//var numbersVisited []string

		splitList := strings.Split(list, ",")

		for i, number := range splitList {
			allowedNumbers := ruleMap[number]
			for _, n := range splitList[i+1:] {

				if !slices.Contains(allowedNumbers, n) {
					listAllowed = false
					break
				}
			}

		}

		if listAllowed {
			middle := len(splitList) / 2
			val, _ := strconv.Atoi(splitList[middle])
			count += val
		} else {
			//do incorrect list logic

			slices.SortFunc(splitList, func(a, b string) int {
				allowedNumbers := ruleMap[a]
				if slices.Contains(allowedNumbers, b) {
					return -1
				}
				return 1
			})

			//for i := 0; i < len(splitList)-1; i++ {
			//	allowedNumbers := ruleMap[splitList[i]]
			//
			//	//for j := 0; j < len(splitList)-i-1; j++ {
			//	//	if !slices.Contains(allowedNumbers, splitList[j]) {
			//	//		splitList[j], splitList[j+1] = splitList[j+1], splitList[j]
			//	//	}
			//	//}
			//}

			val, _ := strconv.Atoi(splitList[len(splitList)/2])
			wrongCount += val
		}

	}

	fmt.Println("Wrong count : ", wrongCount)
	return count
}
