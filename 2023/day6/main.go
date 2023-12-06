package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/jonathanwalter/dev/Advent-of-code/2023/day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//partOne(scanner)
	partTwo(scanner)

}

type race struct {
	time     int
	distance int
}

func partOne(scanner *bufio.Scanner) {
	races := createRaces(scanner)

	var total int

	for _, r := range races {
		ways := getWinningWays(r)

		if total == 0 {
			total += ways
		} else {
			total = total * ways
		}
	}

	log.Println(total)
}

func partTwo(scanner *bufio.Scanner) {
	race := createLongRace(scanner)
	log.Printf("%v", race)

	var startTime int
	var endTime int

	for currentTime := 1; currentTime <= race.time; currentTime++ {
		timeRemaining := race.time - currentTime
		distanceTravelled := timeRemaining * currentTime

		if distanceTravelled > race.distance {
			startTime = currentTime
			break
		}
	}

	for currentTime := race.time; currentTime > 0; currentTime-- {
		timeRemaining := race.time - currentTime
		distanceTravelled := timeRemaining * currentTime

		if distanceTravelled > race.distance {
			endTime = currentTime
			break
		}
	}

	log.Println(endTime - startTime + 1)

}

func getWinningWays(race race) int {
	var winningTimeCombos int
	for currentTime := 1; currentTime <= race.time; currentTime++ {
		timeRemaining := race.time - currentTime
		distanceTravelled := timeRemaining * currentTime

		if distanceTravelled > race.distance {
			winningTimeCombos++
		}
	}

	return winningTimeCombos
}

func createRaces(scanner *bufio.Scanner) []race {
	var times []int
	var distances []int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Time") {
			timeStrings := strings.Fields(strings.Split(line, ":")[1])
			for _, time := range timeStrings {
				timeInt, _ := strconv.Atoi(time)
				times = append(times, timeInt)
			}

		}
		if strings.Contains(line, "Distance") {
			distanceStrings := strings.Fields(strings.Split(line, ":")[1])
			for _, distnace := range distanceStrings {
				distanceInt, _ := strconv.Atoi(distnace)
				distances = append(distances, distanceInt)
			}

		}
	}

	var races []race
	for i, time := range times {
		races = append(races, race{
			time:     time,
			distance: distances[i],
		})
	}

	return races
}

func createLongRace(scanner *bufio.Scanner) race {
	var bigRace = race{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Time") {
			timeStrings := strings.Fields(strings.Split(line, ":")[1])
			var timeString string
			for _, time := range timeStrings {
				timeString = timeString + time
			}

			time, _ := strconv.Atoi(timeString)
			bigRace.time = time
		}
		if strings.Contains(line, "Distance") {
			distanceStrings := strings.Fields(strings.Split(line, ":")[1])
			var distanceString string
			for _, distnace := range distanceStrings {
				distanceString = distanceString + distnace
			}

			distance, _ := strconv.Atoi(distanceString)
			bigRace.distance = distance

		}
	}
	return bigRace
}
