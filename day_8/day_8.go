package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	input := readInput("\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) (acc int) {
	acc, _ = executeProgram(input)
	return
}

func part2(input []string) (acc int){
	acc = 0
	for index, instruction := range input {
		operation := instruction[:3]
		switch operation {
		case "nop":
			newInstruction := strings.Replace(instruction, "nop", "jmp", 1)
			input[index] = newInstruction
			calculatedAcc, finished := executeProgram(input)

			if finished {
				acc = calculatedAcc
				return
			}

			input[index] = instruction
		case "jmp":
			newInstruction := strings.Replace(instruction, "jmp", "nop", 1)
			input[index] = newInstruction
			calculatedAcc, finished := executeProgram(input)

			if finished {
				acc = calculatedAcc
				return
			}

			input[index] = instruction
		default:
			continue
		}
	}
	return
}

func executeProgram(input []string) (acc int, finished bool) {
	acc = 0
	finished = false
	index := 0
	visited := map[int]bool{}

	for index < len(input) {
		if _, found := visited[index]; found {
			// Found loop
			return
		}
		visited[index] = true
		index, acc = executeInstruction(input[index], index, acc)
	}

	finished = true
	return
}

func executeInstruction(instruction string, currentIndex, currentAcc int) (index int, acc int) {
	i := strings.Split(instruction, " ")
	operation := i[0]
	argument, _ := strconv.Atoi(i[1])
	switch operation {
	case "acc":
		acc = currentAcc + argument
		index = currentIndex + 1
	case "jmp":
		acc = currentAcc
		index = currentIndex + argument
	case "nop":
		acc = currentAcc
		index = currentIndex + 1
	}

	return
}

func readInput(delimiter string) []string {
	input, _ := ioutil.ReadFile("day_8.txt")
	inputArray := strings.Split(string(input), delimiter)
	return inputArray
}
