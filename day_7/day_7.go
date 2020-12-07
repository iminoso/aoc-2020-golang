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
	generateBags(input)
}

func generateBags(input []string) {
	r, _ := regexp.Compile(`([0-9]+) ([a-z, ]+)`)
	var bagIndex = map[string]map[string]int{}
	for _, ruleSet := range input {
		rule := strings.Split(ruleSet, "bags contain")
		bagKey := strings.TrimSpace(rule[0])
		containsBags := rule[1]
		bagMap := map[string]int{}
		for _, bag := range strings.Split(containsBags, `,`) {
			m := r.FindStringSubmatch(strings.TrimSpace(bag))
			if len(m) > 0 {
				val, _ := strconv.Atoi(m[1])
				key := filterBag(m[2])
				bagMap[key] = val
			}
		}
		bagIndex[bagKey] = bagMap
	}

	solutionP1 := 0
	for key := range bagIndex {
		if key != "shiny gold" {
			if dfsPart1(bagIndex, bagIndex[key], "shiny gold") > 0 {
				solutionP1++
			}
		}
	}

	solutionP2 := dfsPart2(bagIndex, bagIndex["shiny gold"])

	fmt.Println("Part 1 Solution:", solutionP1)
	fmt.Println("Part 2 Solution:", solutionP2)
}

func dfsPart1(bagIndex map[string]map[string]int, bagMap map[string]int, searchKey string) int {
	if len(bagMap) == 0 {
		return 0
	}

	count := 0
	for key := range bagMap {
		if key == searchKey {
			count++
		} else {
			count += dfsPart1(bagIndex, bagIndex[key], searchKey)
		}
	}
	return count
}

func dfsPart2(bagIndex map[string]map[string]int, bagMap map[string]int) int {
	if len(bagMap) == 0 {
		return 0
	}
	
	count := 0
	for key, val := range bagMap {
		count += val + val*dfsPart2(bagIndex, bagIndex[key])
	}
	return count
}

func filterBag(bag string) string {
	filteredBag := strings.Split(bag, "bag")[0]
	return strings.TrimSpace(filteredBag)
}

func readInput(delimiter string) []string {
	input, _ := ioutil.ReadFile("day_7.txt")
	inputArray := strings.Split(string(input), delimiter)
	return inputArray
}
