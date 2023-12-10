package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	run(scanner)

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
var loopCoordinates map[string]string

func run(scanner *bufio.Scanner) {
	loopCoordinates = map[string]string{}

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
					addCoordinateToMap(currentTile)
					loopLength = distanceTravelled
					currentTile = c
					break
				}
			} else {
				if !(c.x == previousTile.x && c.y == previousTile.y) && isConnectingPipe(c.fromDirection, c.toDirection, adjacentTile, currentTileValue) {
					addCoordinateToMap(currentTile)
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

	numberInside := 0

	//

	// for y, row := range rows {
	// 	firstReached := false
	// 	threshholdsCrossed := 0
	// 	for x, _ := range row {
	// 		coordString := fmt.Sprintf("%s,%s", fmt.Sprint(x), fmt.Sprint(y))
	// 		previousCoord := fmt.Sprintf("%s,%s", fmt.Sprint(x-1), fmt.Sprint(y))
	// 		if _, ok := loopCoordinates[coordString]; ok {
	// 			if !firstReached {
	// 				threshholdsCrossed++
	// 			} else {
	// 				//if this pipe is connected to the last pipe, we dont count it as a threshhold crossed
	// 				if x != 0 {
	// 					if _, ok := loopCoordinates[previousCoord]; ok {

	// 						previousTile := rows[y][x-1]
	// 						currentTile := rows[y][x]
	// 						nextTile := rows[y][x+1]
	// 						if isPipeConnectedBothWays("W", "E", currentTile, previousTile, nextTile) {
	// 							continue //ignore this tile
	// 						}else{
	// 							threshholdsCrossed++
	// 						}

	// 					}else{
	// 						threshholdsCrossed++
	// 					}
	// 				}
	// 			}
	// 		} else {
	// 			if threshholdsCrossed != 0 && threshholdsCrossed%2 == 1 {
	// 				numberInside++
	// 			}
	// 			// if _, ok := loopCoordinates[previousCoord]; ok {
	// 			// 	threshholdsCrossed++
	// 			// }

	// 		}
	// 	}
	// }

	var cleanRows [][]string

	for _, row := range rows{
		var tiles []string
		for i:=0; i < len(row); i++{
			tiles = append(tiles, ".")
		}
		cleanRows = append(cleanRows, tiles)
	}

	for k, v := range loopCoordinates{
		
	}

	for y, row := range rows {
		threshholdsCrossed := 0
		for x, _ := range row {
			coordString := fmt.Sprintf("%s,%s", fmt.Sprint(x), fmt.Sprint(y))
			if _, ok := loopCoordinates[coordString]; ok {
				continue
			}
			for i := x; i < len(row); i++ {
				coordString := fmt.Sprintf("%s,%s", fmt.Sprint(i), fmt.Sprint(y))
				if _, ok := loopCoordinates[coordString]; ok {
					if row[i] == "-" {
						continue
					} else {
						threshholdsCrossed++
					}
				}
			}
			if threshholdsCrossed != 0 && threshholdsCrossed%2 == 1 {
				numberInside++
			}
			threshholdsCrossed = 0
		}
	}

	log.Println("part 1", loopLength/2+1)
	log.Println("part 2", numberInside)
}

func addCoordinateToMap(c coord) {
	coordString := fmt.Sprintf("%s,%s", fmt.Sprint(c.x), fmt.Sprint(c.y))
	loopCoordinates[coordString] = true
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

func isPipeConnectedBothWays(fromDirection string, toDirection string, currentTile string, previousTile string, nextTile string) bool {
	var left bool
	var right bool
	for _, direction := range pipeDirectory[currentTile].connectingDirections {
		if direction == fromDirection {
			left = true
		}

	}
	for _, direction := range pipeDirectory[nextTile].connectingDirections {
		if direction == fromDirection {
			right = true
		}

	}

	return left && right
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
