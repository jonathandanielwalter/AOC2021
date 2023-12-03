package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
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

func partOne(scanner *bufio.Scanner) {
	var sum int
	var sum2 int

	matrix := createMatrix(scanner)

	for y, row := range matrix {

		var currentNumber string
		var currentNumberAdjacent bool

		for x, item := range row {
			/////// part one logic
			if unicode.IsDigit(item) {
				currentNumber += string(item)

				if !currentNumberAdjacent && isAdjacentSymbol(x, y, matrix) {
					currentNumberAdjacent = true
				}
			}
			//if the current number has ended or it isnt a number
			if (x == len(row)-1 || !unicode.IsDigit(row[x+1])) && currentNumberAdjacent {
				//log.Println("current number is ", currentNumber)
				val, err := strconv.Atoi(currentNumber)
				if err != nil {
					panic(err)
				}
				sum += val

				//reset
				currentNumber = ""
				currentNumberAdjacent = false
			}

			if (x == len(row)-1 || !unicode.IsDigit(row[x+1])) && !currentNumberAdjacent {
				//reset
				currentNumber = ""
				currentNumberAdjacent = false
			}
			/////// part one logic

			//////// part 2 logic
			if item == '*' {
				adjacentNumbers := getAdjacentNumbers(x, y, matrix)

				if len(adjacentNumbers) > 1 {
					first, err := strconv.Atoi(adjacentNumbers[0])
					if err != nil {
						log.Panic(err)
					}
					second, err := strconv.Atoi(adjacentNumbers[1])
					if err != nil {
						log.Panic(err)
					}

					sum2 = sum2 + (first * second)
				}

			}
		}

	}

	log.Println("part one", sum)
	log.Println("part two", sum2)
}

func isSymbol(point rune) bool {
	if point != '.' && !unicode.IsNumber(point) {
		return true
	}
	return false
}

func createMatrix(scanner *bufio.Scanner) [][]rune {
	var matrix [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}
	return matrix
}

func isAdjacentSymbol(x, y int, matrix [][]rune) bool {
	//check left
	if x != 0 {
		if isSymbol(matrix[y][x-1]) {
			return true
		}
	}
	//check right
	if x < len(matrix[y])-1 {
		if isSymbol(matrix[y][x+1]) {
			return true
		}
	}

	//check above
	if y != 0 {
		if isSymbol(matrix[y-1][x]) {
			return true
		}

		//above left
		if x != 0 {
			if isSymbol(matrix[y-1][x-1]) {
				return true
			}
		}

		//above right
		if x < len(matrix[y])-1 {
			if isSymbol(matrix[y-1][x+1]) {
				return true
			}
		}
	}
	//check below
	if y < len(matrix)-1 {
		if isSymbol(matrix[y+1][x]) {
			return true
		}

		//below left
		if x != 0 {
			if isSymbol(matrix[y+1][x-1]) {
				return true
			}
		}

		//below right
		if x < len(matrix[y])-1 {
			if isSymbol(matrix[y+1][x+1]) {
				return true
			}
		}
	}

	return false
}

func getAdjacentNumbers(x, y int, matrix [][]rune) []string {
	var adjacentNumbers []string

	//check left
	if x != 0 {
		if unicode.IsDigit(matrix[y][x-1]) {
			adjacentNumbers = append(adjacentNumbers, buildNumber(x-1, y, matrix))
		}
	}
	//check right
	if x < len(matrix[y])-1 {
		if unicode.IsDigit(matrix[y][x+1]) {
			adjacentNumbers = append(adjacentNumbers, buildNumber(x+1, y, matrix))
		}
	}

	//check above
	if y != 0 {
		//middle out
		if unicode.IsDigit(matrix[y-1][x]) {
			adjacentNumbers = append(adjacentNumbers, buildNumber(x, y-1, matrix))
		} else {
			if x != 0 && unicode.IsDigit(matrix[y-1][x-1]) {
				adjacentNumbers = append(adjacentNumbers, buildNumber(x-1, y-1, matrix))
			}
			if x < len(matrix[y])-1 && unicode.IsDigit(matrix[y-1][x+1]) {
				adjacentNumbers = append(adjacentNumbers, buildNumber(x+1, y-1, matrix))
			}
		}

		// //iterates from top left through the right so we dont duplicate the same number
		// if x != 0 && unicode.IsDigit(matrix[y-1][x-1]) {
		// 	adjacentNumbers = append(adjacentNumbers, buildNumber(x-1, y-1, matrix))
		// } else if unicode.IsDigit(matrix[y-1][x]) {
		// 	adjacentNumbers = append(adjacentNumbers, buildNumber(x, y-1, matrix))
		// } else if x < len(matrix[y])-1 && unicode.IsDigit(matrix[y-1][x+1]) {
		// 	adjacentNumbers = append(adjacentNumbers, buildNumber(x+1, y-1, matrix))
		// }
	}
	//check below
	if y < len(matrix)-1 {
		//middle out
		if unicode.IsDigit(matrix[y+1][x]) {
			adjacentNumbers = append(adjacentNumbers, buildNumber(x, y+1, matrix))
		} else {
			if x != 0 && unicode.IsDigit(matrix[y+1][x-1]) {
				adjacentNumbers = append(adjacentNumbers, buildNumber(x-1, y+1, matrix))
			}
			if x < len(matrix[y])-1 && unicode.IsDigit(matrix[y+1][x+1]) {
				adjacentNumbers = append(adjacentNumbers, buildNumber(x+1, y+1, matrix))
			}
		}

		// //iterates from top left through the right so we dont duplicate the same number
		// if x != 0 && unicode.IsDigit(matrix[y+1][x-1]) {
		// 	adjacentNumbers = append(adjacentNumbers, buildNumber(x-1, y+1, matrix))
		// } else if unicode.IsDigit(matrix[y+1][x]) {
		// 	adjacentNumbers = append(adjacentNumbers, buildNumber(x, y+1, matrix))
		// } else if x < len(matrix[y])-1 && unicode.IsDigit(matrix[y+1][x+1]) {
		// 	adjacentNumbers = append(adjacentNumbers, buildNumber(x+1, y+1, matrix))
		// }
	}

	return adjacentNumbers
}

func buildNumber(x, y int, matrix [][]rune) string {
	leftPointer := x - 1
	var number string = string(matrix[y][x])

	//look left
	for leftPointer >= 0 {
		if unicode.IsDigit(matrix[y][leftPointer]) {
			number = string(matrix[y][leftPointer]) + number
			leftPointer--
		} else {
			break
		}
	}

	rightPointer := x + 1
	//look right
	for rightPointer < len(matrix[y]) {
		if unicode.IsDigit(matrix[y][rightPointer]) {
			number = number + string(matrix[y][rightPointer])
			rightPointer++
		} else {
			break
		}
	}

	return number
}
