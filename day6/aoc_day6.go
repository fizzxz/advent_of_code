package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {

	fileData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := convertToStringSlice(string(fileData))
	starts := partOne(line)
	fmt.Println(starts)
	startsTwo := partTwo(line)
	fmt.Println(startsTwo)
}

func partOne(line []string) string {
	chars := ""

	for i := 0; i <= len(line)-4; i++ {
		mapChars := make(map[string]bool)
		for charnum := i; charnum < i+4; charnum++ {
			mapChars[line[charnum]] = true
		}
		if len(mapChars) == 4 {
			chars = strconv.Itoa(i + 4)
			return chars
		}
	}
	return chars
}

func partTwo(line []string) string {
	chars := ""

	for i := 0; i <= len(line)-14; i++ {
		mapChars := make(map[string]bool)
		for charnum := i; charnum < i+14; charnum++ {
			mapChars[line[charnum]] = true
		}
		if len(mapChars) == 14 {
			chars = strconv.Itoa(i + 14)
			return chars
		}
	}
	return chars
}

func convertToStringSlice(input string) []string {
	var output []string
	for i := 0; i < len(input); i++ {
		output = append(output, string(input[i]))
	}
	return output
}
