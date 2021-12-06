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

	//var bingoCalledNumbers string
	var allBoards []Board

	var previousLineBlank bool = false
	var currentBoard Board
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, ",") {
			//bingoCalledNumbers = row
			previousLineBlank = false
		} else if row == "" {
			previousLineBlank = true
		} else {
			if previousLineBlank {
				var board = createBoard(row) //create new board
				allBoards = append(allBoards, board)
				currentBoard = board
				previousLineBlank = false

			} else {
				addRowAndAddToColumns(currentBoard, row)
				previousLineBlank = false
			}
		}

	}
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

	//fmt.Printf("%v", board)
	return board
}

func addRowAndAddToColumns(board Board, newRow string) {
	numbers := strings.Fields(newRow)

	var row Row = Row{}
	row.Numbers = numbers
	board.Rows = append(board.Rows, row)

	for i, number := range numbers {
		board.Columns[i].Numbers = append(board.Columns[i].Numbers, number)
	}

	fmt.Printf("%v", board.Columns)

}
