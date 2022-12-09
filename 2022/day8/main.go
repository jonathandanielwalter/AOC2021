package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var highestScore int = 0

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\n") // running on mac

	treeGrid := [][]int{}

	for _, input := range inputs {
		treeGrid = append(treeGrid, turnToInts(input))
	}

	log.Println(numberOfVisibleTrees(treeGrid))
	log.Println(highestScore)
}

func numberOfVisibleTrees(treeGrid [][]int) int {
	visible := 0

	for i, row := range treeGrid {
		//its an edge row
		if i == 0 || i == len(treeGrid)-1 {
			visible += len(row)
		} else {
			for x, _ := range row {
				if x == 0 || x == len(row)-1 {
					visible++
				} else {

					visibleUp, upDistance := checkUp(treeGrid, i, x)
					visableDown, downDistance := checkDown(treeGrid, i, x)
					visableLeft, leftDistance := checkLeft(row, x)
					visableRight, rightDistance := checkRight(row, x)

					if visibleUp || visableDown || visableLeft || visableRight {
						visibilityScore := upDistance * downDistance * leftDistance * rightDistance
						if visibilityScore > highestScore {
							highestScore = visibilityScore
						}
						visible++
					}
				}

			}
		}

	}

	return visible

}

func checkUp(grid [][]int, rowIndex int, columnIndex int) (bool, int) {
	treeHeight := grid[rowIndex][columnIndex]

	visible := true

	viewingDistance := 0

	for i := rowIndex - 1; i >= 0; i-- {
		if grid[i][columnIndex] >= treeHeight {
			//log.Printf("%v was larger than current tree %v", grid[rowIndex][i], treeHeight)
			visible = false
			viewingDistance++
			break
		} else {
			viewingDistance++
		}
	}

	return visible, viewingDistance
}

func checkDown(grid [][]int, rowIndex int, columnIndex int) (bool, int) {
	treeHeight := grid[rowIndex][columnIndex]
	visible := true

	viewingDistance := 0

	for i := rowIndex + 1; i <= len(grid)-1; i++ {
		if grid[i][columnIndex] >= treeHeight {
			//log.Printf("%v was larger than current tree %v", grid[rowIndex][i], treeHeight)
			visible = false
			viewingDistance++
			break
		} else {
			viewingDistance++
		}
	}

	return visible, viewingDistance
}

func checkLeft(row []int, index int) (bool, int) {
	visible := true
	treeHeight := row[index]

	viewingDistance := 0

	for i := index - 1; i >= 0; i-- {
		if row[i] >= treeHeight {
			//log.Printf("%v was larger than current tree %v", row[i], treeHeight)
			visible = false
			viewingDistance++
			break
		} else {
			viewingDistance++
		}
	}

	return visible, viewingDistance
}

func checkRight(row []int, index int) (bool, int) {
	visible := true
	treeHeight := row[index]

	viewingDistance := 0

	for i := index + 1; i <= len(row)-1; i++ {
		if row[i] >= treeHeight {
			//log.Printf("%v was larger than current tree %v", row[i], treeHeight)
			visible = false
			viewingDistance++
			break
		} else {
			viewingDistance++
		}
	}

	return visible, viewingDistance
}

func turnToInts(str string) []int {
	numbers := []int{}
	for _, c := range str {
		number, _ := strconv.Atoi(string(c))
		numbers = append(numbers, number)
	}

	//log.Println(numbers)
	return numbers
}
