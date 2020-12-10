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
	sort.Ints(input)
	m := part1(input)
	fmt.Println("Part 1 Solution:", m[1]*m[3])
	fmt.Println("Part 2 Solution:", part2(input))
}

func part2(baseInput []int) (count int) {
	input := make([]int, len(baseInput))
	copy(input, baseInput)
	input = append([]int{0}, input...)
	input = append(input, input[len(input)-1]+3)

	exists := make(map[int]int)
	for i, v := range input {
		exists[v] = i
	}

	permutations := map[int]int{len(input) - 1: 1}
	for i := len(input) - 2; i >= 0; i-- {
		sum := 0
		for diff := 1; diff <= 3; diff++ {
			if pos, ok := exists[input[i]+diff]; ok {
				sum += permutations[pos]
			}
		}
		permutations[i] = sum
	}
	count = permutations[0]
	return
}

func part1(baseInput []int) map[int]int {
	input := copyArray(baseInput)
	m := map[int]int{}
	current := 0
	i := 0

	for i < len(input) {
		difference := input[i] - current
		_, found := m[difference]
		if found {
			m[difference]++
		} else {
			m[difference] = 1
		}
		current = input[i]
		i++
	}
	if _, ok := m[3]; ok {
		m[3]++
	} else {
		m[3] = 1
	}

	return m
}

func copyArray(arr []int) (copyArr []int) {
	copyArr = make([]int, len(arr))
	copy(copyArr, arr)
	return
}

func readInput(delimiter string) []int {
	input, _ := ioutil.ReadFile("day_10.txt")
	inputArray := strings.Split(string(input), delimiter)
	inputArrayNum := make([]int, len(inputArray))
	for i, val := range inputArray {
		num, _ := strconv.Atoi(val)
		inputArrayNum[i] = num
	}
	return inputArrayNum
}
