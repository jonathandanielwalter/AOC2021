package main

import (
	"bufio"
	"log"
	"os"
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
	// partOne(scanner)
	partTwo(scanner)

}

func partOne(scanner *bufio.Scanner) {
	var total int

	for scanner.Scan() {
		rows := [][]int{}
		line := strings.Fields(scanner.Text())

		var allZeros bool
		rows = append(rows, convertStringsToInts(line))
		rowNumber := 0
		for !allZeros {
			newRow := []int{}
			newRowTotal := 0
			currentRow := rows[rowNumber]
			for i := 0; i < len(currentRow)-1; i++ {
				diff := currentRow[i+1] - currentRow[i]
				newRowTotal += diff
				newRow = append(newRow, diff)
			}
			rows = append(rows, newRow)
			if newRowTotal == 0 {
				allZeros = true
			}
			rowNumber++
		}

		//iterate backwards
		lastRow := len(rows) - 1
		for i := lastRow; i > 0; i-- {
			currentRow := rows[i]
			rowAbove := rows[i-1]
			newValue := currentRow[len(currentRow)-1] + rowAbove[len(rowAbove)-1]

			rowAbove = append(rowAbove, newValue)
			rows[i-1] = rowAbove
		}

		total += rows[0][len(rows[0])-1]

	}
	log.Println("part 1 ", total)
}

func partTwo(scanner *bufio.Scanner) {
	var total int

	for scanner.Scan() {
		rows := [][]int{}
		line := strings.Fields(scanner.Text())

		var allZeros bool
		rows = append(rows, convertStringsToInts(line))
		rowNumber := 0
		for !allZeros {
			newRow := []int{}
			newRowTotal := 0
			currentRow := rows[rowNumber]
			for i := 0; i < len(currentRow)-1; i++ {
				diff := currentRow[i+1] - currentRow[i]
				newRowTotal += diff
				newRow = append(newRow, diff)
			}
			rows = append(rows, newRow)
			if newRowTotal == 0 {
				allZeros = true
			}
			rowNumber++
		}

		//iterate backwards
		lastRow := len(rows) - 1
		for i := lastRow; i > 0; i-- {
			currentRow := rows[i]
			rowAbove := rows[i-1]
			newValue := rowAbove[0] - currentRow[0]

			rowAbove = append(rowAbove, newValue)
			rowAbove = append([]int{newValue}, rowAbove...)
			rows[i-1] = rowAbove
		}

		total += rows[0][0]

	}
	log.Println("part 2 ", total)
}

func convertStringsToInts(strings []string) []int {
	ints := []int{}

	for _, str := range strings {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}

	return ints
}
