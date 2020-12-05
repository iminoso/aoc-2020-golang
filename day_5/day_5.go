package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	input := readInput("\n")
	solution(input)
}

func solution(input []string) {
	maxID := 0
	minID := math.MaxInt64
	var IDList [1024]bool
	for _, zone := range input {
		row := zone[0:7]
		rowVal := binarySpacePartition(row, 0, 127, "F", "B")
		col := zone[7:]
		colVal := binarySpacePartition(col, 0, 7, "L", "R")
		idVal := rowVal*8 + colVal
		if idVal > maxID {
			maxID = idVal
		}
		if idVal < minID {
			minID = idVal
		}
		IDList[idVal] = true
	}

	fmt.Println("Part 1 Solution", maxID)

	// Go through all existing ids until finding the empty id (ie. seat)
	i := minID
	for true {
		if !IDList[i] {
			fmt.Println("Part 2 Solution", i)
			break
		}
		i++
	}
	return
}

func binarySpacePartition(input string, start int, end int, low string, high string) int {
	if len(input) == 1 {
		if string(input[0]) == low {
			return start
		}

		if string(input[0]) == high {
			return end
		}
	}

	if string(input[0]) == low {
		return binarySpacePartition(input[1:], start, divideRoundDown(start+end), low, high)
	}
	return binarySpacePartition(input[1:], divideRoundUp(start+end), end, low, high)
}

func divideRoundUp(num int) int {
	return int(math.Ceil(float64(num) / float64(2)))
}

func divideRoundDown(num int) int {
	return int(math.Floor(float64(num) / float64(2)))
}

func readInput(delimiter string) []string {
	input, _ := ioutil.ReadFile("day_5.txt")
	inputArray := strings.Split(string(input), delimiter)
	return inputArray
}
