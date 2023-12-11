package main

import (
	"bufio"
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

func run(scanner *bufio.Scanner) {
	galaxy := [][]string{}

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		galaxy = append(galaxy, row)
	}

	//log.Println(galaxy)

	expand(galaxy)
}

func expand(galaxy [][]string) {
	for y := 0; y < len(galaxy); y++ {
		allTheSame := true
		for _, column := range galaxy[y] {
			if column != "." {
				allTheSame = false
				break
			}
		}
		if allTheSame {
			galaxy = append(galaxy[:y+1], galaxy[y:]...)
			y++
		}
	}

	var sameIndexes []int

	for x := 0; x < len(galaxy[0]); x++ {
		allTheSame := true
		for y := 0; y < len(galaxy); y++ {
			if galaxy[y][x] != "." {
				allTheSame = false
			}
		}
		if allTheSame {
			sameIndexes = append(sameIndexes, x)
		}
	}

	for i, index := range sameIndexes {
		for y := 0; y < len(galaxy); y++ {
			galaxy[y] = append(galaxy[y][:index+i+1], galaxy[y][index+i:]...)
		}
	}

	for _, row := range galaxy {
		log.Println(row)
	}

}
