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
	passportDataList := buildPassportData(input)
	fmt.Println("Part 1 Solution:", part1(passportDataList))
	fmt.Println("Part 2 Solution:", part2(passportDataList))
}

func part1(passportDataList []map[string]string) (validCount int) {
	validCount = 0
	for _, passportData := range passportDataList {
		if validPassportDataPart1(passportData) {
			validCount++
		}
	}
	return
}

func part2(passportDataList []map[string]string) (validCount int) {
	validCount = 0
	for _, passportData := range passportDataList {
		if validPassportDataPart2(passportData) {
			validCount++
		}
	}
	return
}

func buildPassportData(input []string) (passportDataList []map[string]string) {
	passportDataList = []map[string]string{}
	passportData := map[string]string{}
	for _, line := range input {
		if line == "" {
			passportDataList = append(passportDataList, passportData)
			passportData = make(map[string]string)
		} else {
			fieldsList := strings.Split(line, " ")
			for _, fields := range fieldsList {
				field := strings.Split(fields, ":")
				key := field[0]
				val := field[1]
				passportData[key] = val
			}
		}
	}
	if len(passportData) > 0 {
		passportDataList = append(passportDataList, passportData)
	}
	return
}

func validPassportDataPart1(passportData map[string]string) (valid bool) {
	valid = false

	if len(passportData) == 8 {
		valid = true
	}

	_, found := passportData["cid"]
	if len(passportData) == 7 && !found {
		valid = true
	}

	return
}

func validPassportDataPart2(passportData map[string]string) (valid bool) {
	valid = false

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if _, found := passportData["byr"]; !found {
		return
	}
	byr, err := strconv.Atoi(passportData["byr"])
	if err != nil || byr < 1920 || byr > 2002 {
		return
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if _, found := passportData["iyr"]; !found {
		return
	}
	iyr, err := strconv.Atoi(passportData["iyr"])
	if err != nil || iyr < 2010 || iyr > 2020 {
		return
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if _, found := passportData["eyr"]; !found {
		return
	}
	eyr, err := strconv.Atoi(passportData["eyr"])
	if err != nil || eyr < 2020 || eyr > 2030 {
		return
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	if _, found := passportData["hgt"]; !found {
		return
	}
	hgt := passportData["hgt"]
	hgtUnit := hgt[len(hgt)-2:]
	if hgtUnit != "in" && hgtUnit != "cm" {
		return
	}
	hgtValStr := hgt[0 : len(hgt)-2]
	hgtVal, err := strconv.Atoi(hgtValStr)
	if err != nil {
		return
	}
	if hgtUnit == "cm" && (hgtVal < 150 || hgtVal > 193) {
		return
	}
	if hgtUnit == "in" && (hgtVal < 59 || hgtVal > 76) {
		return
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if _, found := passportData["hcl"]; !found {
		return
	}
	hcl := passportData["hcl"]
	if string(hcl[0]) != "#" {
		return
	}
	found, err := regexp.MatchString("#[0-9a-f]{6}", hcl)
	if err != nil || !found {
		return
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if _, found := passportData["ecl"]; !found {
		return
	}
	eclValues := map[string]int{"amb": 1, "blu": 1, "brn": 1, "gry": 1, "grn": 1, "hzl": 1, "oth": 1}
	if _, found := eclValues[passportData["ecl"]]; !found {
		return
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if _, found := passportData["pid"]; !found {
		return
	}
	if len(passportData["pid"]) != 9 {
		return
	}
	_, err = strconv.Atoi(passportData["pid"])
	if err != nil {
		return
	}

	valid = true
	return
}

func readInput(delimiter string) []string {
	input, _ := ioutil.ReadFile("day_4.txt")
	inputArray := strings.Split(string(input), delimiter)
	return inputArray
}
