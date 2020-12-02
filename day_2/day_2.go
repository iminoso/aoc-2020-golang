package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readInput("\n")
	fmt.Println("Valid Passwords Part 1", part1(input))
	fmt.Println("Valid Passwords Part 2", part2(input))
}

func part1(input []string) (validPasswords int) {
	validPasswords = 0
	for _, line := range input {
		minCount, maxCount, targetChar, password := parseLine(line)
		targetCount := strings.Count(password, targetChar)
		if targetCount >= minCount && targetCount <= maxCount {
			validPasswords++
		}
	}
	return
}

func part2(input []string) (validPasswords int) {
	validPasswords = 0
	for _, line := range input {
		firstIndex, secondIndex, targetChar, password := parseLine(line)
		firstIndex--
		secondIndex--
		firstChar := string(password[firstIndex])
		secondChar := string(password[secondIndex])
		if xor(firstChar == targetChar, secondChar == targetChar) {
			validPasswords++
		}
	}
	return
}

func xor(a bool, b bool) bool {
	return ( a || b ) && !( a && b )
}

func parseLine(line string) (minCount int, maxCount int, targetChar string, password string) {
	r := regexp.MustCompile(`([0-9]+)\-([0-9]+) ([a-z])\: ([a-z]*)`)
	allSubmatches := r.FindAllStringSubmatch(line, -1)
	minCount, _ = strconv.Atoi(allSubmatches[0][1])
	maxCount, _ = strconv.Atoi(allSubmatches[0][2])
	targetChar = allSubmatches[0][3]
	password = allSubmatches[0][4]
	return
}

func readInput(delimiter string) []string {
	input, _ := ioutil.ReadFile("day_2.txt")
	inputArray := strings.Split(string(input), delimiter)
	return inputArray
}
