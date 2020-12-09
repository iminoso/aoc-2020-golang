package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := readInput("\n")
	preambleLen := 25
	invalidNum := findInvalid(input, preambleLen)
	fmt.Println("Part 1 Solution:", invalidNum)
	contiguousSet := findContiguousSet(input, invalidNum)
	sort.Ints(contiguousSet)
	fmt.Println("Part 2 Solution:", contiguousSet[0]+contiguousSet[len(contiguousSet)-1])
}

func findContiguousSet(input []int, target int) []int {
	var contiguousSet []int
	for keyI, valI := range input {
		contiguousSet = append(contiguousSet, valI)
		if keyI+1 < len(input) {
			for _, valJ := range input[keyI+1:] {
				currentSetSum := sumArray(contiguousSet)
				if currentSetSum == target {
					return contiguousSet
				}

				if currentSetSum < target {
					contiguousSet = append(contiguousSet, valJ)
				} else {
					contiguousSet = nil
					break
				}
			}
		}
	}
	return contiguousSet
}

func sumArray(input []int) (sum int) {
	sum = 0
	for _, val := range input {
		sum += val
	}
	return
}

func findInvalid(input []int, preambleLen int) (invalidNum int) {
	invalidNum = -1
	set := map[int]bool{}
	for key, val := range input {
		if key < preambleLen {
			set[val] = true
			continue
		}
		if !checkInvalidList(input[key-preambleLen:key], set, val) {
			invalidNum = val
			return
		}
		set[val] = true
		delete(set, input[key-preambleLen])
	}
	return
}

func checkInvalidList(input []int, set map[int]bool, target int) (valid bool) {
	valid = false

	for _, val := range input {
		if _, found := set[target-val]; found {
			valid = true
			return
		}
	}
	return
}

func readInput(delimiter string) []int {
	input, _ := ioutil.ReadFile("day_9.txt")
	inputArray := strings.Split(string(input), delimiter)
	inputArrayNum := make([]int, len(inputArray))
	for i, val := range inputArray {
		num, _ := strconv.Atoi(val)
		inputArrayNum[i] = num
	}
	return inputArrayNum
}
