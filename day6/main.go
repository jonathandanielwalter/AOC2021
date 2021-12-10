package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type LanternfishCounter struct {
	InternalCounter     int
	NumberOfLanternfish int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//var lanternfishCounters []LanternfishCounter

	var numbers []string

	//build from input
	for scanner.Scan() {
		row := scanner.Text()

		numbers = strings.Split(row, ",")
	}

	fishMap := make(map[int]int)

	for _, numberString := range numbers {
		number, _ := strconv.Atoi(numberString)
		if val, ok := fishMap[number]; ok {
			fishMap[number] = val + 1
		} else {
			fishMap[number] = 1
		}

	}
	log.Printf("%v", fishMap)

	for i := 0; i < 256; i++ {
		originalZero := fishMap[0]

		fishMap[0] = fishMap[1]
		fishMap[1] = fishMap[2]
		fishMap[2] = fishMap[3]
		fishMap[3] = fishMap[4]
		fishMap[4] = fishMap[5]
		fishMap[5] = fishMap[6]
		fishMap[6] = fishMap[7]
		fishMap[7] = fishMap[8]

		fishMap[8] = originalZero
		fishMap[6] = fishMap[6] + originalZero
	}

	var count int

	for _, fish := range fishMap {
		count = count + fish
	}

	println(count)
}
