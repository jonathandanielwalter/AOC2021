package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

func (c coord) isValid(grid [][]string) bool {
	if c.x < 0 || c.x > len(grid[0])-1 {
		return false
	}

	if c.y < 0 || c.y > len(grid)-1 {
		return false
	}

	return true
}

func (c coord) string() string {
	return fmt.Sprintf("%v,%v", c.x, c.y)
}

func main() {
	file, err := os.Open("C:\\Users\\jonat\\playground\\Advent-of-code\\2024\\day8\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows [][]string

	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, strings.Split(line, ""))
	}

	fmt.Println(part1(rows))
	fmt.Println(part2(rows))
}

func part1(rows [][]string) int {
	frequencyMap := map[string][]coord{}

	for y, row := range rows {
		for x, freq := range row {
			if freq == "." {
				continue
			}
			if _, ok := frequencyMap[freq]; ok {
				frequencyMap[freq] = append(frequencyMap[freq], coord{x: x, y: y})
			} else {
				frequencyMap[freq] = []coord{{x: x, y: y}}
			}
		}
	}

	antinodes := map[string]bool{}
	for _, v := range frequencyMap {
		// for every frequency coord of a type (e.g. a)
		for i := 0; i < len(v); i++ {

			for j := 0; j < len(v); j++ {
				if i == j {
					continue
				}

				currentNode := v[i]
				checkingNode := v[j]

				xDiff := currentNode.x - checkingNode.x
				yDiff := currentNode.y - checkingNode.y
				antinode := coord{
					x: currentNode.x + xDiff,
					y: currentNode.y + yDiff,
				}

				if antinode.isValid(rows) {
					antinodes[antinode.string()] = true
				}
			}

		}
	}

	return len(antinodes)
}

func part2(rows [][]string) int {
	frequencyMap := map[string][]coord{}

	for y, row := range rows {
		for x, freq := range row {
			if freq == "." {
				continue
			}
			if _, ok := frequencyMap[freq]; ok {
				frequencyMap[freq] = append(frequencyMap[freq], coord{x: x, y: y})
			} else {
				frequencyMap[freq] = []coord{{x: x, y: y}}
			}
		}
	}

	antinodes := map[string]bool{}
	for _, v := range frequencyMap {
		// for every frequency coord of a type (e.g. a)
		for i := 0; i < len(v); i++ {

			for j := 0; j < len(v); j++ {
				if i == j {
					continue
				}

				currentNode := v[i]
				checkingNode := v[j]

				antinodes[currentNode.string()] = true

				xDiff := currentNode.x - checkingNode.x
				yDiff := currentNode.y - checkingNode.y

				antinode := coord{
					x: currentNode.x + xDiff,
					y: currentNode.y + yDiff,
				}

				for antinode.isValid(rows) {

					antinodes[antinode.string()] = true

					antinode = coord{
						x: antinode.x + xDiff,
						y: antinode.y + yDiff,
					}
				}

			}

		}
	}

	return len(antinodes)
}
