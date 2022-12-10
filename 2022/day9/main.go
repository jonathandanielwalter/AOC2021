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
	Rope     []*Pair
)

//4978 too low
//5048 too low

// part 2
// 4989 too high

type Node struct {
	Name       int
	ParentNode *Node
	X          int
	Y          int
	ChildNode  *Node
}

type Pair struct {
	X int
	Y int
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\n") // running on mac

	makeRope(10)

	visitMap["0,0"] = true

	for _, move := range inputs {
		processMove(move, Rope)
	}

	// log.Println(visitMap)
	log.Println(len(visitMap))
}

func makeRope(size int) {
	for i := 0; i < size; i++ {
		Rope = append(Rope, &Pair{X: 0, Y: 0})
	}
}

func processMove(move string, rope []*Pair) {
	moveComponents := strings.Fields(move)

	direction := moveComponents[0]
	amount, _ := strconv.Atoi(moveComponents[1])

	for i := 0; i < amount; i++ {
		if direction == "U" {
			head := rope[0]
			// headPreviousX := head.X
			// headPreviousY := head.Y
			head.Y++
			moveRope()
			// bubble(Rope[1], head, headPreviousX, headPreviousY, rope, 1)
		}
		if direction == "D" {
			head := rope[0]
			// headPreviousX := head.X
			// headPreviousY := head.Y
			head.Y--
			moveRope()
			// bubble(Rope[1], head, headPreviousX, headPreviousY, rope, 1)
		}
		if direction == "L" {
			head := rope[0]
			// headPreviousX := head.X
			// headPreviousY := head.Y
			head.X--
			moveRope()
			// bubble(Rope[1], head, headPreviousX, headPreviousY, rope, 1)
		}
		if direction == "R" {
			head := rope[0]
			// headPreviousX := head.X
			// headPreviousY := head.Y
			head.X++
			moveRope()
			// bubble(Rope[1], head, headPreviousX, headPreviousY, rope, 1)
		}

		// last := rope[len(rope)-1]
		// xString := strconv.Itoa(last.X)
		// yString := strconv.Itoa(last.Y)
		// coords := fmt.Sprintf("%s,%s", xString, yString)
		// // log.Println(coords)
		// visitMap[coords] = true

	}
}

func moveRope() {
	for i := 1; i < len(Rope); i++ {
		xDelta, yDelta := delta(Rope[i-1].X, Rope[i].X, Rope[i-1].Y, Rope[i].Y)

		if Abs(xDelta) <= 1 && Abs(yDelta) <= 1 {
			return
		}

		if yDelta > 0 {
			Rope[i].Y++
		} else if yDelta < 0 {
			Rope[i].Y--
		}
		if xDelta > 0 {
			Rope[i].X++
		} else if xDelta < 0 {
			Rope[i].X--
		}

	}

	last := Rope[len(Rope)-1]
	xString := strconv.Itoa(last.X)
	yString := strconv.Itoa(last.Y)
	coords := fmt.Sprintf("%s,%s", xString, yString)
	// log.Println(coords)
	visitMap[coords] = true
}

func delta(parentX, childX, parentY, childY int) (int, int) {
	xDelta := parentX - childX
	yDelta := parentY - childY
	return xDelta, yDelta
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
