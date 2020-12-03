package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readInput("\n")
	fmt.Println("Part 1 Solution: ", part1(input, 3, 1))
	part2Solution := part1(input, 1, 1) * part1(input, 3, 1) * part1(input, 5, 1) * part1(input, 7, 1) * part1(input, 1, 2)
	fmt.Println("Part 2 Solution: ", part2Solution)
}

func part1(input []string, colIncrement int, rowIncrement int) (treeCount int) {
	rowLength := len(input)
	colLength := len(input[0])
	row := 0
	col := 0
	treeCount = 0

	for row < rowLength {
		if string(input[row][col%colLength]) == "#" {
			treeCount++
		}
		row += rowIncrement
		col += colIncrement
	}
	return
}

func readInput(delimiter string) []string {
	input, _ := ioutil.ReadFile("day_3.txt")
	inputArray := strings.Split(string(input), delimiter)
	return inputArray
}
