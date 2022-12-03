package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		text := len(scanner.Text()) / 2

		fmt.Printf("Text: %d, Text Total %d", text, len(scanner.Text()))
		fmt.Println(" /")
	}
}
