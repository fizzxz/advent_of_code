package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	fileData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fileData), "\n")

	stacks := initialStacks()

	for i := 10; i < len(lines); i++ {

		var moveAmount, from, dest int
		fmt.Sscanf(lines[i], "move %d from %d to %d", &moveAmount, &from, &dest)
		for i := 0; i < moveAmount; i++ {
			move(from-1, dest-1, stacks)
		}

	}
	fmt.Printf("top most letters for part1: ")
	topMostLetters(stacks)
	fmt.Println(" ")
	elapsed := time.Since(start)
	log.Printf("Task 1 took %s", elapsed)
	stacks = initialStacks()
	for i := 10; i < len(lines); i++ {
		var moveAmount, from, dest int
		fmt.Sscanf(lines[i], "move %d from %d to %d", &moveAmount, &from, &dest)
		movePartTwo(moveAmount, from-1, dest-1, stacks)
	}

	fmt.Printf("top most letters for part2: ")
	topMostLetters(stacks)
	fmt.Println(" ")
	elapsed = time.Since(start)
	log.Printf("Task 2 took %s", elapsed)
}

func initialStacks() [][]string {
	stacks := make([][]string, 9)
	stack := []string{}
	stack = append(stack, "GFVHPS")
	stack = append(stack, "GJFBVDZM")
	stack = append(stack, "GMLJN")
	stack = append(stack, "NGZVDWP")
	stack = append(stack, "VRCB")
	stack = append(stack, "VRSMPWLZ")
	stack = append(stack, "THP")
	stack = append(stack, "QRSNCHZV")
	stack = append(stack, "FLGPVQJ")

	for i, strings := range stack {
		for _, character := range strings {

			stacks[i] = append(stacks[i], string(character))

		}
	}
	return stacks
}

func move(from int, dest int, stack [][]string) {
	stack[dest] = append(stack[dest], stack[from][len(stack[from])-1])
	stack[from] = stack[from][:len(stack[from])-1]
}

func movePartTwo(numOfLetters int, from int, dest int, stack [][]string) {
	sumOfLetters := stack[from][len(stack[from])-numOfLetters:]
	for i := 0; i < numOfLetters; i++ {
		stack[dest] = append(stack[dest], sumOfLetters[i])
		stack[from] = stack[from][:len(stack[from])-1]
	}
}

func topMostLetters(strings [][]string) {
	for _, letter := range strings {
		fmt.Printf("%s", letter[len(letter)-1])
	}

}
