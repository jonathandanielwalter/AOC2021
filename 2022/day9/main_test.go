package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestThis(t *testing.T) {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	restOfRope := createRope(11)

	//head := Node{
	//	Name:       1,
	//	ParentNode: nil,
	//	X:          0,
	//	Y:          0,
	//	ChildNode:  &restOfRope,
	//}

	inputs := strings.Split(string(b), "\n") // running on mac

	for _, move := range inputs {
		processMove(move, &restOfRope)
	}
}
