package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readInput("\n\n")

	count := 0
	for _, members := range input {
		a := readGroupAnswersPart1(members)
		count += a
	}
	fmt.Println("Part 1 Solution:", count)

	count = 0
	for _, members := range input {
		a := readGroupAnswersPart2(members)
		count += a
	}
	fmt.Println("Part 2 Solution:", count)
}

func readGroupAnswersPart1(members string) int {
	set := map[string]bool{}
	memberAnswers := strings.Split(members, "\n")
	for _, memberAnswer := range memberAnswers {
		for _, answer := range memberAnswer {
			answerLetter := string(answer)
			if _, found := set[answerLetter]; !found {
				set[answerLetter] = true
			}
		}
	}
	return len(set)
}

func readGroupAnswersPart2(members string) int {
	m := map[string]int{}
	count := 0
	memberAnswers := strings.Split(members, "\n")

	if len(memberAnswers) == 1 {
		return len(memberAnswers[0])
	}

	for _, memberAnswer := range memberAnswers {
		for _, answer := range memberAnswer {
			answerLetter := string(answer)
			_, found := m[answerLetter]
			if !found {
				m[answerLetter] = 1
			} else {
				m[answerLetter] = m[answerLetter] + 1
			}
		}
	}

	for _, val := range m {
		if val == len(memberAnswers) {
			count++
		}
	}
	return count
}

func readInput(delimiter string) []string {
	input, _ := ioutil.ReadFile("day_6.txt")
	inputArray := strings.Split(string(input), delimiter)
	return inputArray
}
