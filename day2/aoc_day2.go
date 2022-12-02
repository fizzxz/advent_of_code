package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	countPoints, lines := 0, 0
	for scanner.Scan() {
		scannerTextRemoveSpace := strings.ReplaceAll(scanner.Text(), " ", "")
		opponent := string(scannerTextRemoveSpace[0:1])
		me := string(scannerTextRemoveSpace[1:])
		countPoints += response(opponent, me)
		lines++
		fmt.Println(scannerTextRemoveSpace)

	}
	fmt.Println(countPoints)
	fmt.Print(lines)
}

//A - X = ROCK
//B - Y = PAPER
//C - Z = SCISSORS
func response(opponent, me string) int {
	switch opponent {
	case "A":
		if me == "X" {
			return 3 + 1
		}
		if me == "Y" {
			return 6 + 2
		}
		if me == "Z" {
			return 0 + 3
		}
	case "B":
		if me == "X" {
			return 0 + 1
		}
		if me == "Y" {
			return 3 + 2
		}
		if me == "Z" {
			return 6 + 3
		}
	case "C":
		if me == "X" {
			return 6 + 1
		}
		if me == "Y" {
			return 0 + 2
		}
		if me == "Z" {
			return 3 + 3
		}
	}

	return 0

}
