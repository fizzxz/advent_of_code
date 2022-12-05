package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count1 := 0
	for scanner.Scan() {

		line := scanner.Text()
		splitLine := strings.Split(line, ",")
		firstSeq := []int{}
		secondSeq := []int{}
		firstSeq = sequenceOfDigitsFromString(splitLine[0])
		secondSeq = sequenceOfDigitsFromString(splitLine[1])
		if isElementInSlice(firstSeq, secondSeq) || isElementInSlice(secondSeq, firstSeq) {
			count1++
		}
	}
	fmt.Println(count1)
	fmt.Println(" ")
	elapsed := time.Since(start)
	log.Printf("Task 1 took %s", elapsed)
}

func sequenceOfDigitsFromString(input string) []int {
	Digits := strings.Split(input, "-")
	numbers := []int{}
	x := 0
	for _, val := range Digits {
		x, _ = strconv.Atoi(val)
		numbers = append(numbers, x)
	}
	return sequenceOfDigitsFromTwoDigits(numbers)
}

func sequenceOfDigitsFromTwoDigits(input []int) []int {
	output := []int{}
	for i := input[0]; i <= input[1]; i++ {
		output = append(output, i)
	}

	return output
}

func isElementInSlice(searchIn []int, searchFrom []int) bool {
	count := 0
	for _, whatToLookFor := range searchFrom {
		if find(whatToLookFor, searchIn) {
			count++
		}
		if count == len(searchFrom) {
			return true
		}
	}
	return false
}

func find(input int, valuesToSearch []int) bool {
	for _, findVal := range valuesToSearch {
		if findVal == input {
			return true
		}
	}
	return false
}
