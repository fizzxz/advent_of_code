package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rockPaperScissors struct {
	rock    map[string]int
	paper   map[string]int
	scissor map[string]int
}

type score struct {
	win  int
	draw int
	loss int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	countPointsStratGuide1 := 0
	countPointsStratGuide2 := 0
	for scanner.Scan() {
		scannerTextRemoveSpace := strings.ReplaceAll(scanner.Text(), " ", "")
		opponent := string(scannerTextRemoveSpace[0:1])
		me := string(scannerTextRemoveSpace[1:])
		countPointsStratGuide1 += responseStratGuide1(opponent, me)
		countPointsStratGuide2 += responseStratGuide2(opponent, me)
		fmt.Println(scannerTextRemoveSpace)
	}
	fmt.Println(countPointsStratGuide1)

	fmt.Println(countPointsStratGuide2)
}

func responseStratGuide1(opponent, me string) int {
	rps := initialiseRPS()
	scores := initialiseScore()
	switch opponent {
	case "A":
		if rock, ok := rps.rock[me]; ok {
			return scores.draw + rock
		}
		if paper, ok := rps.paper[me]; ok {
			return scores.win + paper
		}
		if scissor, ok := rps.scissor[me]; ok {
			return scores.loss + scissor
		}
	case "B":
		if rock, ok := rps.rock[me]; ok {
			return scores.loss + rock
		}
		if paper, ok := rps.paper[me]; ok {
			return scores.draw + paper
		}
		if scissor, ok := rps.scissor[me]; ok {
			return scores.win + scissor
		}
	case "C":
		if rock, ok := rps.rock[me]; ok {
			return scores.win + rock
		}
		if paper, ok := rps.paper[me]; ok {
			return scores.loss + paper
		}
		if scissor, ok := rps.scissor[me]; ok {
			return scores.draw + scissor
		}
	}
	return 0
}

func responseStratGuide2(opponent, me string) int {
	rps := initialiseRPS()
	scores := initialiseScore()
	switch opponent {
	case "A":
		if _, ok := rps.paper[me]; ok {
			return scores.draw + 1
		}
		if _, ok := rps.scissor[me]; ok {
			return scores.win + 2
		}
		if _, ok := rps.rock[me]; ok {
			return scores.loss + 3
		}
	case "B":
		if _, ok := rps.rock[me]; ok {
			return scores.loss + 1
		}
		if _, ok := rps.paper[me]; ok {
			return scores.draw + 2
		}
		if _, ok := rps.scissor[me]; ok {
			return scores.win + 3
		}
	case "C":
		if _, ok := rps.scissor[me]; ok {
			return scores.win + 1
		}
		if _, ok := rps.rock[me]; ok {
			return scores.loss + 2
		}
		if _, ok := rps.paper[me]; ok {
			return scores.draw + 3
		}
	}
	return 0
}

func initialiseRPS() rockPaperScissors {
	rocks := make(map[string]int)
	papers := make(map[string]int)
	scissors := make(map[string]int)
	rocks["A"], rocks["X"] = 1, 1
	papers["B"], papers["Y"] = 2, 2
	scissors["C"], scissors["Z"] = 3, 3
	rps := rockPaperScissors{
		rock:    rocks,
		paper:   papers,
		scissor: scissors,
	}
	return rps
}

func initialiseScore() score {
	scores := score{
		win:  6,
		draw: 3,
		loss: 0,
	}
	return scores
}
