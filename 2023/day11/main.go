package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	run(scanner)
}

func run(scanner *bufio.Scanner) {
	pairs := map[string]int{}
	galaxy := [][]string{}

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		galaxy = append(galaxy, row)
	}

	expand(galaxy)
	log.Println(galaxy)
	for y, row := range galaxy {
		for x, cell := range row {
			if cell == "#" {
				start := fmt.Sprintf("%v,%v", x, y)

				//traverse right
				traverseRight(pairs, row, start, x, y, 0)

				vericalDistance := 0
				//traverseDown left and right
				if y != len(galaxy)-1 { //if we're not at the last row
					for g := y + 1; g < len(galaxy); g++ {
						if galaxy[g][x] == "." {
							vericalDistance++
						} else if galaxy[g][x] == "X" {
							vericalDistance = vericalDistance + 1000000
						} else { //reached a planet
							end := fmt.Sprintf("%v,%v", x, g)
							vericalDistance++
							addToMap(pairs, start, end, vericalDistance)
							
						}
						traverseRight(pairs, galaxy[g], start, x, g, vericalDistance)
						traverseLeft(pairs, galaxy[g], start, x, g, vericalDistance)
					}

				}
			}
		}
	}
	// log.Printf("%v\n", pairs)

	total := 0
	for _, v := range pairs {
		total += v
	}
	log.Println(total)
	// log.Println(len(pairs))
}

func traverseRight(pairs map[string]int, row []string, startingCoord string, x, y, verticalDistance int) {
	horizontalDistance := 1
	//traverse right
	for i := x + 1; i < len(row); i++ {
		if row[i] == "." {
			horizontalDistance++
		} else if row[i] == "X" {
			horizontalDistance = horizontalDistance + 1000000
		} else { //reached a planet
			end := fmt.Sprintf("%v,%v", i, y)
			addToMap(pairs, startingCoord, end, horizontalDistance+verticalDistance)
			horizontalDistance++
		}
	}
}

func traverseLeft(pairs map[string]int, row []string, startingCoord string, x, y, verticalDistance int) {
	horizontalDistance := 1
	//traverse right
	for i := x-1; i >= 0; i-- {
		if row[i] == "." {
			horizontalDistance++
		} else if row[i] == "X" {
			horizontalDistance = horizontalDistance + 1000000
		} else { //reached a planet
			end := fmt.Sprintf("%v,%v", i, y)
			addToMap(pairs, startingCoord, end, horizontalDistance+verticalDistance)
			horizontalDistance++
		}
	}
}

func addToMap(theMap map[string]int, start string, end string, distance int) {

	if start != end {
		_, startFirst := theMap[fmt.Sprintf("%s-%s", start, end)]
		_, endFirst := theMap[fmt.Sprintf("%s-%s", end, start)]

		if !(startFirst || endFirst) {
			theMap[fmt.Sprintf("%s-%s", start, end)] = distance
		}
	}

}

func expand(galaxy [][]string) [][]string {
	for y := 0; y < len(galaxy); y++ {
		allTheSame := true
		for _, column := range galaxy[y] {
			if column != "." {
				allTheSame = false
				break
			}
		}
		if allTheSame {
			for x, _ := range galaxy[y] {
				galaxy[y][x] = "X"
			}

		}
	}

	var sameIndexes []int

	for x := 0; x < len(galaxy[0]); x++ {
		allTheSame := true
		for y := 0; y < len(galaxy); y++ {
			if galaxy[y][x] == "#" {
				allTheSame = false
			}
		}
		if allTheSame {
			sameIndexes = append(sameIndexes, x)
		}
	}

	for _, index := range sameIndexes {
		for y := 0; y < len(galaxy); y++ {
			//galaxy[y] = append(galaxy[y][:index+i+1], galaxy[y][index+i:]...)
			galaxy[y][index] = "X"
		}
	}

	//for _, row := range galaxy {
	//	log.Println(row)
	//}
	return galaxy
}
