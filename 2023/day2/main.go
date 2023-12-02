package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

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
	var total int
	var part2Total int

	for scanner.Scan() {

		var highestBlue int
		var highestGreen int
		var highestRed int

		line := scanner.Text()

		gameSplit := strings.Split(line, ":")
		game := gameSplit[0]

		roundsString := gameSplit[1]
		rounds := strings.Split(strings.TrimSpace(roundsString), ";")

		roundValid := true

		//split rounds
		for _, round := range rounds {
			colours := strings.Split(round, ",")

			for _, colour := range colours {
				colourValue := strings.Split(strings.TrimSpace(colour), " ")
				valueStr := colourValue[0]
				colour := colourValue[1]

				value, err := strconv.Atoi(valueStr)
				if err != nil {
					panic(err)
				}

				switch colour {
				case "green":
					if value > maxGreen {
						roundValid = false
					}

					if value > highestGreen {
						highestGreen = value
					}

				case "blue":
					if value > maxBlue {
						roundValid = false
					}

					if value > highestBlue {
						highestBlue = value
					}

				case "red":
					if value > maxRed {
						roundValid = false
					}

					if value > highestRed {
						highestRed = value
					}
				}
			}

		}

		if roundValid {
			total += getGameVal(game)
		}
		gameTotal := highestBlue * highestGreen * highestRed
		part2Total += gameTotal
	}

	log.Println("part one val", total)
	log.Println("part two val", part2Total)
}

func getGameVal(gameVal string) int {
	game := strings.Split(gameVal, " ")

	gameNumber, err := strconv.Atoi(game[1])

	if err != nil {
		panic(err)
	}

	return gameNumber
}
