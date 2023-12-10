package main

import (
	"bufio"
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
	partOne(scanner)

}

type pipe struct {
	connectingDirections []string
}

var pipeDirectory = map[string]pipe{
	"|": {
		connectingDirections: []string{"N", "S"},
	},
	"-": {
		connectingDirections: []string{"E", "W"},
	},
	"L": {
		connectingDirections: []string{"N", "E"},
	},
	"J": {
		connectingDirections: []string{"N", "W"},
	},
	"7": {
		connectingDirections: []string{"S", "W"},
	},
	"F": {
		connectingDirections: []string{"S", "E"},
	},
}

type coord struct {
	x             int
	y             int
	fromDirection string
}

var rows [][]string
var loopLength int

func partOne(scanner *bufio.Scanner) {
	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), ""))
	}
	log.Println(rows)

	var startTile coord

	for y, row := range rows {
		for x, tile := range row {
			if tile == "S" {
				startTile = coord{x: x, y: y}
			}
		}
	}

	checkTile(startTile, "", 0, startTile)
	log.Println(loopLength/2 + 1)


	
}

func checkTile(currentTile coord, travellingDirection string, distanceTravelled int, previousTile coord) {
	for _, c := range neighborCoordinates(currentTile.x, currentTile.y) {
		if rows[currentTile.y][currentTile.x] == "S" && distanceTravelled > 0 {
			break
		}
		adjacentTile := rows[c.y][c.x]

		switch adjacentTile {
		case ".":
			continue
		case "S": //found a loop
			if (c.x == previousTile.x && c.y == previousTile.y){
				continue
			}
			loopLength = distanceTravelled
			break
		default:
			if !(c.x == previousTile.x && c.y == previousTile.y) && isConnectingPipe(c.fromDirection, adjacentTile) {
				checkTile(c, c.fromDirection, distanceTravelled+1, currentTile)
			}
		}

	}
}

func isConnectingPipe(travellingDirection string, tile string) bool {
	for _, direction := range pipeDirectory[tile].connectingDirections {
		if direction == travellingDirection {
			return true
		}
	}

	return false
}

// only cares about directly, non diagonal neighbours
func neighborCoordinates(x, y int) []coord {
	var coords []coord

	if y != 0 {
		coords = append(coords, coord{x: x, y: y - 1, fromDirection: "S"})
	}
	if y != len(rows)-1 {
		coords = append(coords, coord{x: x, y: y + 1, fromDirection: "N"})
	}

	if x != 0 {
		coords = append(coords, coord{x: x - 1, y: y, fromDirection: "E"})
	}

	if x != len(rows[0])-1 {
		coords = append(coords, coord{x: x + 1, y: y, fromDirection: "W"})
	}

	return coords
}
