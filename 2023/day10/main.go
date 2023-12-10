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
	toDirection   string
}

var rows [][]string
var loopLength int

func partOne(scanner *bufio.Scanner) {
	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), ""))
	}
	var currentTile coord
	for y, row := range rows {
		for x, tile := range row {
			if tile == "S" {
				currentTile = coord{x: x, y: y}
			}
		}
	}
	var previousTile coord
	var distanceTravelled int

	for true {

		coords := neighborCoordinates(currentTile.x, currentTile.y)
		for _, c := range coords {
			currentTileValue := rows[currentTile.y][currentTile.x]
			if currentTileValue == "S" && distanceTravelled > 0 {
				break
			}
			adjacentTile := rows[c.y][c.x]

			if adjacentTile == "." {
				continue
			} else if adjacentTile == "S" {
				if c.x == previousTile.x && c.y == previousTile.y {
					continue
				}
				if isFinalConnection(c.fromDirection, c.toDirection, currentTileValue) {
					loopLength = distanceTravelled
					currentTile = c
					break
				}
			} else {
				if !(c.x == previousTile.x && c.y == previousTile.y) && isConnectingPipe(c.fromDirection, c.toDirection, adjacentTile, currentTileValue) {
					previousTile = currentTile
					currentTile = c
					distanceTravelled++
					break
				}
			}

		}
		if rows[currentTile.y][currentTile.x] == "S" && distanceTravelled > 0 {
			break
		}
	}

	log.Println(loopLength/2 + 1)

}

func isConnectingPipe(fromDirection string, toDirection string, destinationTile string, originTile string) bool {
	allowableOrigins := pipeDirectory[originTile].connectingDirections

	allowed := false

	if originTile != "S" {
		for _, allowableOrigin := range allowableOrigins {
			if toDirection == allowableOrigin {
				allowed = true
			}
		}
		if !allowed {
			return false
		}
	}

	for _, direction := range pipeDirectory[destinationTile].connectingDirections {
		if direction == fromDirection {
			return true
		}

	}

	return false
}

func isFinalConnection(fromDirection string, toDirection string, originTile string) bool {
	allowableOrigins := pipeDirectory[originTile].connectingDirections

	for _, allowableOrigin := range allowableOrigins {
		if toDirection == allowableOrigin {
			return true
		}
	}

	return false
}

// only cares about directly, non diagonal neighbours
func neighborCoordinates(x, y int) []coord {
	var coords []coord

	if y != 0 {
		coords = append(coords, coord{x: x, y: y - 1, fromDirection: "S", toDirection: "N"})
	}
	if y != len(rows)-1 {
		coords = append(coords, coord{x: x, y: y + 1, fromDirection: "N", toDirection: "S"})
	}

	if x != 0 {
		coords = append(coords, coord{x: x - 1, y: y, fromDirection: "E", toDirection: "W"})
	}

	if x != len(rows[0])-1 {
		coords = append(coords, coord{x: x + 1, y: y, fromDirection: "W", toDirection: "E"})
	}

	return coords
}
