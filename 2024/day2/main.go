package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/jonat/playground/Advent-of-code/2024/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)

		report := stringsToInts(split)

		reports = append(reports, report)
	}

	part1 := partOne(reports)
	fmt.Println(part1)

	part2 := partTwo(reports)
	fmt.Println(part2)
}

func partOne(reports [][]int) int {
	var safeReports int

	for _, report := range reports {
		var ascending bool

		safe := true

		if report[1] > report[0] {
			ascending = true
		}

		for i := 0; i < len(report)-1; i++ {

			//if its the first one, lets work out if we're going up or down
			if i == 0 {
				if report[i] == report[i+1] {
					safe = false
					break
				}

			}

			diff := diff(report[i], report[i+1])

			if ascending {
				if diff < 1 || diff > 3 || report[i+1] < report[i] {
					safe = false
					break
				}
			} else {
				if diff < 1 || diff > 3 || report[i] < report[i+1] {
					safe = false
					break
				}
			}

		}

		if safe {
			safeReports++
		}

	}

	return safeReports
}

func partTwo(reports [][]int) int {
	var safeReports int

	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		} else {
			for i, _ := range report {
				if isSafeWithIndexRemoved(report, i) {
					safeReports++
					break
				}
			}
		}

	}

	return safeReports
}

func isSafe(report []int) bool {
	var ascending bool

	if report[1] > report[0] {
		ascending = true
	}

	for i := 0; i < len(report)-1; i++ {

		//if its the first one, lets work out if we're going up or down
		if i == 0 {
			if report[i] == report[i+1] {
				return false
			}

		}

		diff := diff(report[i], report[i+1])

		if ascending {
			if diff < 1 || diff > 3 || report[i+1] < report[i] {
				return false
			}
		} else {
			if diff < 1 || diff > 3 || report[i] < report[i+1] {
				return false
			}
		}

	}

	return true
}

func isSafeWithIndexRemoved(report []int, i int) bool {
	var ascending bool

	var reducedReport []int
	reducedReport = append(reducedReport, report[:i]...)
	reducedReport = append(reducedReport, report[i+1:]...)

	if reducedReport[1] > reducedReport[0] {
		ascending = true
	}

	for i := 0; i < len(reducedReport)-1; i++ {

		//if its the first one, lets work out if we're going up or down
		if i == 0 {
			if reducedReport[i] == reducedReport[i+1] {
				return false
			}

		}

		diff := diff(reducedReport[i], reducedReport[i+1])

		if ascending {
			if diff < 1 || diff > 3 || reducedReport[i+1] < reducedReport[i] {
				return false
			}
		} else {
			if diff < 1 || diff > 3 || reducedReport[i] < reducedReport[i+1] {
				return false
			}
		}

	}

	return true
}

func diff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}

func stringsToInts(strs []string) []int {
	var ints []int

	for _, val := range strs {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}

	return ints
}
