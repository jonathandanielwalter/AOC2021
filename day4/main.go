package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Board struct {
	Rows    []Row
	Columns []Column
}

type Row struct {
	Numbers []string
}

type Column struct {
	Numbers []string
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bingoCalledNumbers []string
	var allBoards []Board

	var previousLineBlank bool = false
	var currentBoard Board
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, ",") {
			bingoCalledNumbers = strings.Split(row, ",")
			previousLineBlank = false
		} else if row == "" {
			previousLineBlank = true
		} else {
			if previousLineBlank {
				var board = createBoard(row) //create new board
				allBoards = append(allBoards, board)
				currentBoard := &board
				previousLineBlank = false

			} else {
				currentBoard = addRowAndAddToColumns(currentBoard, row)
				previousLineBlank = false
			}
		}

	}

	winner := playBingo(allBoards, bingoCalledNumbers)
	fmt.Printf("%v", winner)
}

func playBingo(allBoards []Board, bingoCalledNumbers []string) Board {

	for _, number := range bingoCalledNumbers {
		for _, board := range allBoards {
			board = mark(board, number)

			for _, column := range board.Columns {
				if len(column.Numbers) == 0 {
					return board
				}
			}
			for _, row := range board.Rows {
				if len(row.Numbers) == 0 {
					return board
				}
			}
		}
	}
	return Board{}
}

func mark(board Board, number string) Board {
	for _, row := range board.Rows {
		for i := len(row.Numbers) - 1; i >= 0; i-- {
			//fmt.Printf("comparing %v and %v\n", row.Numbers[i], number)
			if row.Numbers[i] == number {
				row.Numbers = append(row.Numbers[:i], row.Numbers[i+1:]...)
				//	fmt.Printf("%v\n", row.Numbers)
			}
		}

		for _, column := range board.Columns {
			for i := len(column.Numbers) - 1; i >= 0; i-- {
				if column.Numbers[i] == number {
					column.Numbers = append(column.Numbers[:i], column.Numbers[i+1:]...)
				}
			}
		}
	}

	return board
}

func createBoard(initialRow string) Board {
	println("creating board")
	var board Board = Board{}
	numbers := strings.Fields(initialRow)
	var row Row = Row{}
	row.Numbers = numbers
	board.Rows = append(board.Rows, row)

	for _, number := range numbers {
		var column = Column{}
		column.Numbers = append(column.Numbers, number)
		board.Columns = append(board.Columns, column)
	}

	fmt.Printf("%v", board)
	return board
}

func addRowAndAddToColumns(board *Board, newRow string) {
	numbers := strings.Fields(newRow)

	var row Row = Row{}
	row.Numbers = numbers
	board.Rows = append(board.Rows, row)

	for i, number := range numbers {
		board.Columns[i].Numbers = append(board.Columns[i].Numbers, number)
	}
	fmt.Println(len(board.Columns[0].Numbers))
}
