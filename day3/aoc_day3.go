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
		text := len(scanner.Text()) / 2
		textFirstHalf := scanner.Text()[:text]
		textSecondHalf := scanner.Text()[text:]

		for i := range textFirstHalf {
			if strings.Contains(textSecondHalf, string(textFirstHalf[i])) {
				commonString := string(textFirstHalf[lastIndexAny(textFirstHalf, string(textFirstHalf[i]))])
				prioritySum += lastIndexAny(abc, commonString)
				break
			}
		}
	}
	fmt.Println(prioritySum)
}

func lastIndexAny(i, x string) int {
	return strings.LastIndexAny(i, x)
}
