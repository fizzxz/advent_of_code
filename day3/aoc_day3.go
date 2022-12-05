package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const abc = ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ."

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	prioritySum1 := 0
	prioritySum2 := 0

	var threeLinesGrouped []string
	for scanner.Scan() {
		line := scanner.Text()
		prioritySum1 = splitLine_FindCommonString_ReturnPrioritySum(line, prioritySum1)

		threeLinesGrouped = append(threeLinesGrouped, line)
		if len(threeLinesGrouped) == 3 {

			prioritySum2 = linesOfThree_FindCommonString_ReturnPrioritySum(threeLinesGrouped, prioritySum2)
			threeLinesGrouped = nil

		}
	}
	fmt.Printf("Priority sum for task part 1: %d", prioritySum1)
	fmt.Println(" ")
	fmt.Printf("Priority sum for task part 2: %d", prioritySum2)
	fmt.Println(" ")
	elapsed := time.Since(start)
	log.Printf("Task 1 and 2 took %s", elapsed)
}

func linesOfThree_FindCommonString_ReturnPrioritySum(Groupedlines []string, prioritySum int) int {
	line1 := Groupedlines[0]
	line2 := Groupedlines[1]
	line3 := Groupedlines[2]
	for i := range line1 {
		if strings.Contains(line2, string(line1[i])) &&
			strings.Contains(line3, string(line1[i])) {
			commonString := findCommonString(line1, i)
			prioritySum += lastIndexAny(abc, commonString)
			break
		}
	}
	return prioritySum
}

func splitLine_FindCommonString_ReturnPrioritySum(line string, prioritySum int) int {
	text := len(line) / 2
	textFirstHalf := line[:text]
	textSecondHalf := line[text:]

	for i := range textFirstHalf {
		if strings.Contains(textSecondHalf, string(textFirstHalf[i])) {
			commonString := findCommonString(textFirstHalf, i)
			prioritySum += lastIndexAny(abc, commonString)
			break
		}
	}
	return prioritySum
}

func findCommonString(strToLookInto string, indexOfStr int) string {
	return string(strToLookInto[lastIndexAny(strToLookInto, string(strToLookInto[indexOfStr]))])
}

func lastIndexAny(i, x string) int {
	return strings.LastIndexAny(i, x)
}
