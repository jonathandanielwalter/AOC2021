package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var seedToSoil []input
var soilToFertilizer []input
var fertilizerToWater []input
var waterToLight []input
var lightToTemperature []input
var temperatureToHumidity []input
var humidityToLocation []input

func main() {
	file, err := os.ReadFile("/Users/jonathanwalter/dev/Advent-of-code/2023/day5/input.txt")
	//file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	blocks := strings.Split(string(file), "\n")
	blocks = removeBlanks(blocks)

	//log.Println(blocks)
	partOne(blocks)
}

func partOne(lines []string) {
	seedStrings := strings.Fields(strings.Split(lines[0], ":")[1])
	seedInts := convertSeedInts(seedStrings)
	createInputSets(lines[1:])

	//var lowestLocation int
	for _, seed := range seedInts {
		soil := getDestinationValue(seed)
		log.Println(soil)
	}
}

func getDestinationValue(source int) int {
	var soil int
	for _, row := range seedToSoil {
		soil = row.source //if its not in a range use the soruce
		if source >= row.source && source < (row.source+row.valueRange) {
			diff := source - row.source

			soil = row.destination + diff
			return soil
		}
	}
	return source
}

func convertSeedInts(seeds []string) []int {
	var seedInts []int
	for _, seed := range seeds {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seedInts = append(seedInts, seedInt)
	}
	return seedInts
}

type input struct {
	destination int
	source      int
	valueRange  int
}

func createInputSets(lines []string) {
	var currentMap string
	for _, line := range lines {

		if strings.Contains(line, "map") {
			currentMap = line
			continue
		}

		if currentMap == "seed-to-soil map:" {
			seedToSoil = addValuesToSlice(seedToSoil, line)
		} else if currentMap == "soil-to-fertilizer map:" {
			soilToFertilizer = addValuesToSlice(soilToFertilizer, line)
		} else if currentMap == "fertilizer-to-water map:" {
			fertilizerToWater = addValuesToSlice(fertilizerToWater, line)
		} else if currentMap == "water-to-light map:" {
			waterToLight = addValuesToSlice(waterToLight, line)
		} else if currentMap == "light-to-temperature map:" {
			lightToTemperature = addValuesToSlice(lightToTemperature, line)
		} else if currentMap == "temperature-to-humidity map:" {
			temperatureToHumidity = addValuesToSlice(temperatureToHumidity, line)
		} else if currentMap == "humidity-to-location map:" {
			humidityToLocation = addValuesToSlice(humidityToLocation, line)
		}
	}
}

func addValuesToSlice(currentInputSet []input, line string) []input {
	values := strings.Fields(line)

	destination, _ := strconv.Atoi(values[0])
	source, _ := strconv.Atoi(values[1])
	valueRange, _ := strconv.Atoi(values[2])

	currentInputSet = append(currentInputSet, input{
		source:      source,
		destination: destination,
		valueRange:  valueRange,
	})

	return currentInputSet
}

func createMap(block string) map[int]int {
	mapper := map[int]int{}
	inputsStr := strings.Fields(strings.Split(block, ":")[1])
	inputs := make([]int, len(inputsStr))

	for i, s := range inputsStr {
		inputs[i], _ = strconv.Atoi(s)
	}

	log.Println(inputs)
	return mapper
}

func removeBlanks(list []string) []string {
	var output []string
	for _, str := range list {
		if str != "" {
			output = append(output, str)
		}
	}
	return output
}
