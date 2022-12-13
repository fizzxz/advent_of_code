package main

//based off of https://github.com/lynerist/Advent-of-code-2022-golang/tree/main/Day_07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	isFile   bool
	subDirs  map[string]*node
	upperDir *node
}

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var currentDirectory *node
	var currentDirectoryPartTwo *node

	dirs1 := []*node{}
	dirs2 := []*node{}
	for sc.Scan() {
		currentDirectory, dirs1 = directorySearchPartOne(sc.Text(), currentDirectory, dirs1)
		currentDirectoryPartTwo, dirs2 = directorySearchPartTwo(sc.Text(), currentDirectoryPartTwo, dirs2)
	}
	partOne(currentDirectory, dirs1)
	PartTwo(currentDirectoryPartTwo, dirs2)

}

func partOne(currentDirectory *node, dirs1 []*node) {

	var totalSize int

	for _, dir := range dirs1 {
		size := calcSize(*dir)
		if size <= 100000 {
			totalSize += size
		}
	}

	fmt.Println(totalSize)
}

func directorySearchPartOne(sc string, currentDirectory *node, dirs []*node) (*node, []*node) {
	line := strings.Fields(sc)
	if len(line) > 2 {
		if line[2] == ".." {
			currentDirectory = currentDirectory.upperDir
		} else if line[2] == "/" {
			currentDirectory = &node{"/", 0, false, make(map[string]*node), nil}
		} else {
			currentDirectory = currentDirectory.subDirs[line[2]]
		}
	} else if line[0] == "dir" {
		currentDirectory.subDirs[line[1]] = &node{line[1], 0, false, make(map[string]*node), currentDirectory}
		dirs = append(dirs, currentDirectory.subDirs[line[1]])
	} else if line[0] != "$" {
		size, _ := strconv.Atoi(line[0])
		currentDirectory.subDirs[line[1]] = &node{line[1], size, true, nil, currentDirectory}
	}
	return currentDirectory, dirs
}

func PartTwo(currentDirectory *node, dirs []*node) {
	toFree := 30000000 - (70000000 - calcSize(*dirs[0]))
	var smallestEnaugthSize int = calcSize(*dirs[0])

	for _, dir := range dirs {
		size := calcSize(*dir)
		if size > toFree && size-toFree < smallestEnaugthSize-toFree {
			smallestEnaugthSize = size
		}
	}

	fmt.Println(smallestEnaugthSize)
}

func directorySearchPartTwo(sc string, currentDirectory2 *node, dirs []*node) (*node, []*node) {
	line := strings.Fields(sc)
	if len(line) > 2 {
		if line[2] == ".." {
			currentDirectory2 = currentDirectory2.upperDir
		} else if line[2] == "/" {
			currentDirectory2 = &node{"/", 0, false, make(map[string]*node), nil}
			dirs = append(dirs, currentDirectory2)
		} else {
			currentDirectory2 = currentDirectory2.subDirs[line[2]]
		}
	} else if line[0] == "dir" {
		currentDirectory2.subDirs[line[1]] = &node{line[1], 0, false, make(map[string]*node), currentDirectory2}
		dirs = append(dirs, currentDirectory2.subDirs[line[1]])
	} else if line[0] != "$" {
		size, _ := strconv.Atoi(line[0])
		currentDirectory2.subDirs[line[1]] = &node{line[1], size, true, nil, currentDirectory2}
	}
	return currentDirectory2, dirs
}

func calcSize(root node) (size int) {
	if root.isFile {
		return root.size
	}
	for _, d := range root.subDirs {
		size += calcSize(*d)
	}
	return
}
