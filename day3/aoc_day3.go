package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const abc = ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ."

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	prioritySum := 0

	for scanner.Scan() {
		prioritySum = splitLine_FindCommonString_ReturnPrioritySum(scanner, prioritySum)
	}
	fmt.Println(prioritySum)
}

func splitLine_FindCommonString_ReturnPrioritySum(scanner *bufio.Scanner, prioritySum int) int {
	text := len(scanner.Text()) / 2
	textFirstHalf := scanner.Text()[:text]
	textSecondHalf := scanner.Text()[text:]

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
