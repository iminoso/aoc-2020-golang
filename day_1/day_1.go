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
	a, b := findSumPart1(input)
	fmt.Println("Part 1 Numbers:", a, b)
	fmt.Println("Part 1 Solution", a*b)
	a, b, c := findSumPart2(input)
	fmt.Println("Part 2 Numbers:", a, b, c)
	fmt.Println("Part 2 Solution", a*b*c)
}

func findSumPart1(input []int) (int, int) {
	m := map[int]bool{}
	target := 2020
	for _, value := range input {
		if len(m) == 0 {
			m[value] = true
		} else {
			if _, ok := m[target-value]; ok {
				return target - value, value
			}

			m[value] = true
		}
	}
	return 0, 0
}

func findSumPart2(baseInput []int) (int, int, int) {
	input := make([]int, len(baseInput))
	copy(input, baseInput)
	sort.Ints(input[:])
	target := 2020
	for i := 0; i < len(input)-2; i++ {
		j := i + 1
		k := len(input) - 1
		for j < k {
			if input[i]+input[j]+input[k] == target {
				return input[i], input[j], input[k]
			}

			if input[i]+input[j]+input[k] < target {
				j++
			} else {
				k--
			}

		}
	}
	return 0, 0, 0
}

func readInput(delimiter string) []int {
	input, _ := ioutil.ReadFile("day_1.txt")
	inputArray := strings.Split(string(input), delimiter)
	return sliceAtoi(inputArray)
}

func sliceAtoi(sa []string) []int {
	si := make([]int, len(sa))
	for key, val := range sa {
		i, _ := strconv.Atoi(val)
		si[key] = i
	}

	return si
}
