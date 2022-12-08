package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var total int = 0
var smallestSize int = 999999999999
var spaceRequired int = 0

func main() {

	//dirTotals = map[string]int{}
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	inputs := strings.Split(string(b), "\n")

	root := buildFileStructure(inputs)

	rootTotal := traverse(&root)

	//if rootTotal <= 100000 {
	//	total += rootTotal
	//}

	log.Println("part one", total)

	log.Println("root", rootTotal)

	totalSpace := 70000000
	updateSize := 30000000
	spaceAvailable := totalSpace - rootTotal
	spaceRequired = updateSize - spaceAvailable

	log.Println("space required", spaceRequired)

	traverseTwoPointOh(&root)
	log.Println(smallestSize)
}

func traverse(current *Directory) int {
	var dirTotal int = 0
	for _, file := range current.Files {
		dirTotal += file.Size
	}

	for _, childDir := range current.ChildDirectories {
		dirTotal += traverse(childDir)
	}

	if dirTotal <= 100000 {
		//log.Printf("dir total is %v adding to total of %v", dirTotal, total)
		total += dirTotal
	} else {
		//log.Printf("dir %s total %v not less or equal to 100K", current.Name, dirTotal)
	}

	return dirTotal
}

func traverseTwoPointOh(current *Directory) int {
	var dirTotal int = 0
	for _, file := range current.Files {
		dirTotal += file.Size
	}

	if len(current.ChildDirectories) == 0 {
		log.Println(dirTotal)
	}

	for _, childDir := range current.ChildDirectories {
		dirTotal += traverseTwoPointOh(childDir)
	}

	if dirTotal > smallestSize {
		log.Println(dirTotal)
	}

	if dirTotal > spaceRequired && dirTotal < smallestSize {
		smallestSize = dirTotal
	}

	return dirTotal
}

func buildFileStructure(commands []string) Directory {
	var rootDir Directory
	var currentDir *Directory
	var parentDir *Directory

	for _, command := range commands {
		commandComponents := strings.Fields(command)
		if strings.Contains(command, " cd ") {

			if strings.Contains(command, "..") {
				//current dir becomes the parent i.e. moves us up a directory
				if parentDir != nil {
					currentDir = parentDir
					parentDir = currentDir.Parent
				}
			} else {
				dirName := commandComponents[len(commandComponents)-1]

				if dirName == "/" {
					rootDir = Directory{
						Name:             "/",
						Parent:           nil,
						ChildDirectories: make(map[string]*Directory),
						Files:            make(map[string]File),
					}

					currentDir = &rootDir
					parentDir = nil
				} else {
					_, dirExists := currentDir.ChildDirectories[dirName]

					if dirExists {
						currentDir = currentDir.ChildDirectories[dirName]
					} else {
						newDir := &Directory{
							Name:             dirName,
							Parent:           currentDir,
							ChildDirectories: make(map[string]*Directory),
							Files:            make(map[string]File),
						}
						currentDir.ChildDirectories[dirName] = newDir

						parentDir = currentDir
						currentDir = newDir
					}
				}
			}
		} else if isFileInput(command) {
			fileSize, err := strconv.Atoi(commandComponents[0])
			fileName := commandComponents[1]

			if err != nil {
				log.Println("error converting file size to int ", err)
			}

			file := File{
				Name: fileName,
				Size: fileSize,
			}

			currentDir.Files[fileName] = file
		}
	}

	return rootDir
}

type Directory struct {
	Name             string
	Parent           *Directory
	ChildDirectories map[string]*Directory
	Files            map[string]File
}

type File struct {
	Name string
	Size int
}

func isFileInput(input string) bool {
	if _, err := strconv.Atoi(strings.Fields(input)[0]); err == nil {
		return true
	}
	return false
}

//totalDirSize := 0
//currentDirSize := 0
//for _, input := range inputs {
//
//	//if its a file
//	if input[0] != '$' && input[0] != 'd' {
//		fileSize, err := strconv.Atoi(strings.Fields(input)[0])
//		if err != nil{
//			log.Println(err)
//		}
//
//		currentDirSize += fileSize
//	}else{
//		if currentDirSize >
//		totalDirSize += currentDirSize
//		currentDirSize = 0
//	}
//}
