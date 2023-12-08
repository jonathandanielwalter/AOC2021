package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/jonathanwalter/dev/Advent-of-code/2023/day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//partOne(scanner)
	partTwo(scanner)
}

type pair struct {
	left  string
	right string
}

type node struct {
	name       string
	parentNode *node
	leftNode   *node
	rightNode  *node
	level      int
}

var theMap map[string]pair

func partOne(scanner *bufio.Scanner) {
	theMap = map[string]pair{}
	instructions := ""
	for scanner.Scan() {
		line := scanner.Text()
		if instructions == "" {
			instructions = line
			continue
		} else if line == "" {
			continue
		}
		data := strings.Split(line, " = ")
		leftRight := strings.Split(strings.Trim(data[1], "()"), ",")
		theMap[data[0]] = pair{left: strings.TrimSpace(leftRight[0]), right: strings.TrimSpace(leftRight[1])}
	}

	var jumps int
	var currentPlace = "AAA"
	for currentPlace != "ZZZ" {

		for _, instruction := range instructions {
			switch instruction {
			case 'R':
				jumps++
				currentPlace = theMap[currentPlace].right
				if currentPlace == "ZZZ" {
					break
				}

			case 'L':
				jumps++
				currentPlace = theMap[currentPlace].left
				if currentPlace == "ZZZ" {
					break
				}
			}
		}
	}
}

func partTwo(scanner *bufio.Scanner) {
	theMap = map[string]pair{}
	instructions := ""
	for scanner.Scan() {
		line := scanner.Text()
		if instructions == "" {
			instructions = line
			continue
		} else if line == "" {
			continue
		}
		data := strings.Split(line, " = ")
		leftRight := strings.Split(strings.Trim(data[1], "()"), ",")
		theMap[data[0]] = pair{left: strings.TrimSpace(leftRight[0]), right: strings.TrimSpace(leftRight[1])}
	}

	var jumps []int
	for node := range theMap { //nodes
		if string(node[2]) == "A" {
			currentPlace := node
			count := 0
			for string(currentPlace[2]) != "Z" {
				for _, instruction := range instructions {
					count++
					switch instruction {
					case 'R':
						currentPlace = theMap[currentPlace].right
					case 'L':
						currentPlace = theMap[currentPlace].left
					}
				}
			}
			jumps = append(jumps, count)
		}

	}

	log.Println(LCM(jumps...))
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM(x ...int) int {
	if len(x) == 1 {
		return x[0]
	} else if len(x) > 2 {
		return LCM(x[0], LCM(x[1:]...))
	}

	return x[0] * x[1] / GCD(x[0], x[1])
}
func buildChildren(currentNode *node) {
	//look Left
	if currentNode.leftNode != nil {
		childNode := createNode(currentNode.leftNode.name, currentNode)
		if currentNode.parentNode != nil && currentNode.parentNode.name == childNode.name {
			childNode.leftNode = childNode.parentNode
			currentNode.leftNode = childNode
		} else if currentNode.name != childNode.name {
			currentNode.leftNode = childNode
			buildChildren(childNode)
		}

	}
	//look right
	if currentNode.rightNode != nil {
		childNode := createNode(currentNode.rightNode.name, currentNode)
		if currentNode.parentNode != nil && currentNode.parentNode.name == childNode.name {
			childNode.rightNode = childNode.parentNode
			currentNode.rightNode = childNode
		} else if currentNode.name != childNode.name {
			currentNode.rightNode = childNode
			buildChildren(childNode)
		}
	}
}

func createNode(key string, parent *node) *node {
	if val, ok := theMap[key]; ok {
		return &node{
			parentNode: parent,
			name:       key,
			leftNode:   &node{name: val.left},
			rightNode:  &node{name: val.right},
			level:      0,
		}
	}

	return nil
}
