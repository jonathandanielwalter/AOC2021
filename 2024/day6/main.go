package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var visited map[string]bool{}

func main() {
	file, err := os.Open("/Users/jonathanwalter/dev/Advent-of-code/2024/day6/testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows [][]string
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		rows = append(rows, row)
	}

	fmt.Println(part1(rows))
}

func part1(rows [][]string) int {
	var currentX int
	var currentY int

	for x, row := range rows {
		for y, _ := range row {
			if rows[y][x] != "." && rows[y][x] != "#" {
				currentY = y
				currentX = x
				break
			}
		}
	}

	visited[formatCoords(currentX, currentY)] = true
	traverse(rows, currentX, currentY)

	return len(visited)
}

func traverse(rows [][]string, currentX, currentY int) {
	visited[formatCoords(currentX, currentY)] = true

	switch {
	case currentX == 0 && getDirection(rows[currentY][currentX]) == "left":
		return
	case currentX == len(rows[0])-1 && getDirection(rows[currentY][currentX]) == "right":
		return
	case currentY == 0 && getDirection(rows[currentY][currentX]) == "up":
		return
	case currentY == len(rows)-1 && getDirection(rows[currentY][currentX]) == "down":
		return
	}
}

func rotate(rows [][]string, x,y int) {
	switch rows[y][x] {
	case "^":
		rows[y][x] = ">"
	case ">":
		rows[y][x] = "V"
	case "<":
		rows[y][x] = "^"
	case "V":
		rows[y][x] = "<"
	}
}

func getDirection(chev string) string {
	switch chev {
	case "^":
		return "up"
	case ">":
		return "right"
	case "<":
		return "left"
	case "V":
		return "down"
	}
	return ""
}

func formatCoords(x,y int) string{
	return fmt.Sprintf("%v,%v", x,y)
}