package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var calories []int
	calories = getSliceOfCals(scanner, calories)
	sort.Ints(calories)
	fmt.Println(calories[len(calories)-1])
	fmt.Println(calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3])

}

func getSliceOfCals(scanner *bufio.Scanner, calories []int) []int {
	var calCount int
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			calories = append(calories, calCount)
			calCount = 0
		} else {
			intVar, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			calCount += intVar
		}
	}
	return calories
}
