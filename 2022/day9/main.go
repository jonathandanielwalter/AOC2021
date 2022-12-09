package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	visitMap = make(map[string]bool)
)

//4978 too low
//5048 too low

type Node struct {
	Name       int
	ParentNode *Node
	X          int
	Y          int
	ChildNode  *Node
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\n") // running on mac

	visitMap["0,0"] = true

	head := createRope(11)

	log.Println(head)

	for _, move := range inputs {
		processMove(move, &head)
	}

	log.Println(len(visitMap))
}

func createRope(totalNumberOfNodes int) Node {
	var rootNode *Node
	var parentNode *Node
	var currentNode *Node

	for i := 0; i < totalNumberOfNodes; i++ {

		currentNode = &Node{
			Name:       i + 1,
			ParentNode: parentNode,
			X:          0,
			Y:          0,
		}

		if i == 0 {
			rootNode = currentNode
			parentNode = rootNode
		} else {
			parentNode.ChildNode = currentNode
			parentNode = currentNode
		}

	}
	return *rootNode
}

func processMove(move string, head *Node) {
	moveComponents := strings.Fields(move)

	direction := moveComponents[0]
	amount, _ := strconv.Atoi(moveComponents[1])

	//adjacent := isTailAdjacent()

	//log.Println(direction)
	currentNode := head
	for i := 0; i < amount; i++ {
		for currentNode != nil {
			//log.Println("currentNodeName", currentNode.Name)
			//log.Printf("moving %s by 1", direction)
			if currentNode.ParentNode == nil {
				currentNode = currentNode.ChildNode
				continue
			}
			if direction == "U" {
				if isNodeAdjacentAfterMove(currentNode.ParentNode.X, currentNode.ParentNode.Y+1, currentNode.X, currentNode.Y) {
					currentNode.ParentNode.Y++
				} else {
					//log.Println("tail not adjacent after move")
					currentNode.X = currentNode.ParentNode.X
					currentNode.Y = currentNode.ParentNode.Y
					currentNode.ParentNode.Y++
				}
			}
			if direction == "D" {
				if isNodeAdjacentAfterMove(currentNode.ParentNode.X, currentNode.ParentNode.Y-1, currentNode.X, currentNode.Y) {
					currentNode.ParentNode.Y--
				} else {
					currentNode.X = currentNode.ParentNode.X
					currentNode.Y = currentNode.ParentNode.Y
					currentNode.ParentNode.Y--
				}

			}
			if direction == "L" {
				if isNodeAdjacentAfterMove(currentNode.ParentNode.X-1, currentNode.ParentNode.Y, currentNode.X, currentNode.Y) {
					currentNode.ParentNode.X--
				} else {
					currentNode.X = currentNode.ParentNode.X
					currentNode.Y = currentNode.ParentNode.Y
					currentNode.ParentNode.X--
				}
			}
			if direction == "R" {
				if isNodeAdjacentAfterMove(currentNode.ParentNode.X+1, currentNode.ParentNode.Y, currentNode.X, currentNode.Y) {
					currentNode.ParentNode.X++
				} else {
					//log.Println("tail not adjacent after move")
					currentNode.X = currentNode.ParentNode.X
					currentNode.Y = currentNode.ParentNode.Y
					currentNode.ParentNode.X++
				}
			}

			if currentNode.ChildNode == nil {
				xString := strconv.Itoa(currentNode.X)
				yString := strconv.Itoa(currentNode.Y)
				coords := fmt.Sprintf("%s,%s", xString, yString)
				log.Printf("number %v is at coords %s, %s", currentNode.Name, xString, yString)
				visitMap[coords] = true
			}
			currentNode = currentNode.ChildNode
		}
		currentNode = head
	}
}

func isNodeAdjacentAfterMove(headX, headY, tailX, tailY int) bool {
	if headX == tailX && headY == tailY {
		return true
	}

	xSub := headX - tailX

	//adjacent on the x axis and on the same Y so adjacent
	if (xSub == 1 || xSub == -1) && headY == tailY {
		return true
	}

	ysub := headY - tailY
	//log.Println("y difference is ", ysub)

	//adjacent on the Y axis and same X axis
	if (ysub == 1 || ysub == -1) && headX == tailX {
		return true
	}

	if (xSub > 1 || xSub < -1) || (ysub > 1 || ysub < -1) {
		//log.Println("Not adjacent after move")
		return false
	}

	return true
}

//func processMove(move string, head Node) {
//	moveComponents := strings.Fields(move)
//
//	direction := moveComponents[0]
//	amount, _ := strconv.Atoi(moveComponents[1])
//
//	//adjacent := isTailAdjacent()
//
//	//log.Println(direction)
//
//	for i := 0; i < amount; i++ {
//
//		nextNode := &head
//		for nextNode != nil{
//
//		}
//
//		//log.Printf("moving %s by 1", direction)
//		if direction == "U" {
//			if isNodeAdjacentAfterMove(headX, headY+1, tailX, tailY) {
//				headY++
//			} else {
//				//log.Println("tail not adjacent after move")
//				tailX = headX
//				tailY = headY
//				headY++
//			}
//		}
//		if direction == "D" {
//			if isNodeAdjacentAfterMove(headX, headY-1, tailX, tailY) {
//				headY--
//			} else {
//				tailX = headX
//				tailY = headY
//				headY--
//			}
//
//		}
//		if direction == "L" {
//			if isNodeAdjacentAfterMove(headX-1, headY, tailX, tailY) {
//				headX--
//			} else {
//				tailX = headX
//				tailY = headY
//				headX--
//			}
//		}
//		if direction == "R" {
//			if isNodeAdjacentAfterMove(headX+1, headY, tailX, tailY) {
//				headX++
//			} else {
//				//log.Println("tail not adjacent after move")
//				tailX = headX
//				tailY = headY
//				headX++
//			}
//		}
//
//		xString := strconv.Itoa(tailX)
//		yString := strconv.Itoa(tailY)
//		coords := fmt.Sprintf("%s,%s", xString, yString)
//		//log.Println(coords)
//		visitMap[coords] = true
//	}
//}
