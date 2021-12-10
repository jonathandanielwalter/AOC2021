package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	coordinatePairs := readFile("input.txt")

	//fmt.Printf("%v\n", coordinatePairs)
	buildGraph(coordinatePairs)

}

func readFile(filename string) [][][]string {
	var allCoordinatePairs [][][]string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		var coordinatePairs [][]string
		lineCoordinatePair := strings.Split(row, " -> ")

		for _, pair := range lineCoordinatePair {
			for _, coordPair := range strings.Split(pair, " ") {
				coordinatePairs = append(coordinatePairs, strings.Split(coordPair, ","))
			}
		}

		allCoordinatePairs = append(allCoordinatePairs, coordinatePairs)

	}
	return allCoordinatePairs
}

func buildGraph(lineCoordinates [][][]string) {
	// var highestX int
	// var highestY int

	//var board [][]string

	hitCoords := make(map[string]int)

	//fmt.Printf("%v\n", lineCoordinates)

	for _, lineCoordPairs := range lineCoordinates {

		//fmt.Printf("%v\n", lineCoord)

		pair1 := lineCoordPairs[0]
		pair2 := lineCoordPairs[1]

		// fmt.Printf("%v\n", pair1)
		// fmt.Printf("%v\n", pair2)

		x1, _ := strconv.Atoi(pair1[0])
		y1, _ := strconv.Atoi(pair1[1])

		x2, _ := strconv.Atoi(pair2[0])
		y2, _ := strconv.Atoi(pair2[1])

		//if it is straight line
		if x1 == x2 || y1 == y2 {

			var difference int

			if y1 == y2 {
				//println(x1, x2, y1, y2)
				if x1 > x2 {
					difference = x1 - x2
					for i := x2; i <= x2+difference; i++ {
						//	fmt.Printf("Between %v and %v there is %v  and we are on %v\n", x1, x2, difference, i)
						key := fmt.Sprintf("%d,%d", i, y1)
						//println(key)
						addCoordinateToMap(hitCoords, key)
					}
				} else if x2 > x1 {
					//println("x2 greater")
					difference = x2 - x1
					for i := x1; i <= x1+difference; i++ {
						//	fmt.Printf("Between %v and %v there is %v  and we are on %v\n", x2, x1, difference, i)
						key := fmt.Sprintf("%d,%d", i, y1)
						addCoordinateToMap(hitCoords, key)
					}
				}

			} else {
				//println(x1, x2, y1, y2)
				if y1 > y2 {
					difference = y1 - y2
					for i := y2; i <= y2+difference; i++ {
						//	fmt.Printf("Between %v and %v there is %v  and we are on %v\n", y1, y2, difference, i)
						key := fmt.Sprintf("%d,%d", x1, i)
						addCoordinateToMap(hitCoords, key)
					}
				} else if y2 > y1 {
					difference = y2 - y1
					for i := y1; i <= y1+difference; i++ {
						//	fmt.Printf("Between %v and %v there is %v  and we are on %v\n", y2, y1, difference, i)
						key := fmt.Sprintf("%d,%d", x1, i)
						addCoordinateToMap(hitCoords, key)
					}
				}
			}
		} else {
			//diagonal line

			var leftY int
			var rightY int
			var leftX int
			var rightX int

			if x1 < x2 {
				leftX = x1
				leftY = y1
				rightX = x2
				rightY = y2
			} else {
				leftX = x2
				leftY = y2
				rightX = x1
				rightY = y1
			}

			// var startingY int
			// var startingX int
			// var endingY int
			// var endingX int

			//t
			// 8,0 -> 0,8
			// 6,4 -> 2,0
			// 0,0 -> 8,8
			// 5,5 -> 8,2
			// if y1 > y2 && x1 > x2 {
			// 	yDiff = y1 - y2
			// 	xDiff = x1 - x2
			// 	//	endingY = y2
			// } else if y1 > y2 && x1 < x2 {
			// 	yDiff = y1 - y2
			// 	xDiff = x2 - x1
			// 	//	endingY = y1
			// } else if y1 < y2 && x1 > x2 {
			// 	yDiff = y2 - y1
			// 	xDiff = x1 - x2
			// } else if y1 < y2 && x1 < x2 {
			// 	yDiff = y2 - y1
			// 	xDiff = x2 - x1
			// }

			// yDiff = y1 - y2
			// xDiff = x1 - x2

			// if x1 > x2 {
			// 	width = x1 - x2
			// 	startingX = x1
			// 	//	endingX = x2
			// } else {
			// 	width = x2 - x1
			// 	startingX = x2
			// 	//	endingX = x1
			// }

			var xcoords []int
			var ycoords []int
			// println("X Diff ", xDiff)
			// println("Y Diff ", yDiff)

			//then Y will decrease
			//fmt.Printf("lefty y %v right y %v \n", leftY, rightY)
			if leftY > rightY {
				for i := leftY; i >= rightY; i-- {
					ycoords = append(ycoords, i)
				}
			} else { //y will need to increase
				for i := leftY; i <= rightY; i++ {
					ycoords = append(ycoords, i)
				}
			}

			if leftX > rightX {
				for i := leftX; i >= rightX; i-- {
					xcoords = append(xcoords, i)
				}
			} else {
				for i := leftX; i <= rightX; i++ {
					xcoords = append(xcoords, i)
				}
			}

			// fmt.Printf("x coords %v\n", xcoords)
			// fmt.Printf("y coords %v\n", ycoords)

			for i := 0; i < len(xcoords); i++ {
				fmt.Printf("%v,%v\n", xcoords[i], ycoords[i])

				addCoordinateToMap(hitCoords, fmt.Sprintf("%v,%v", xcoords[i], ycoords[i]))
			}

		}

	}
	var count int
	for _, v := range hitCoords {
		if v > 1 {
			count++
			//println(v)
		}
	}
	println(count)
	//fmt.Printf("%v\n", hitCoords)
}

func addCoordinateToMap(coordMap map[string]int, coordinate string) map[string]int {
	if val, ok := coordMap[coordinate]; ok {
		coordMap[coordinate] = val + 1
	} else {
		coordMap[coordinate] = 1
	}
	return coordMap
}
