package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readInput("\n")
	fillSeats := true
	for !validInput(input) {
		points := [][]int{}
		for r := range input {
			for c, val := range input[r] {
				if fillSeats {
					if val == "L" && seatShouldBeFilled(input, r, c) {
						points = append(points, []int{r, c})
					}
				} else {
					if val == "#" && seatShouldBeEmpty(input, r, c) {
						points = append(points, []int{r, c})
					}
				}
			}
		}

		if fillSeats {
			swapIteration(input, points, "#")
		} else {
			swapIteration(input, points, "L")
		}

		fillSeats = !fillSeats
		points = nil
	}

	fmt.Println("Part 1 Solution:", countFilled(input))
}

func countFilled(grid [][]string) (count int) {
	count = 0
	for r := range grid {
		for _, val := range grid[r] {
			if val == "#" {
				count++
			}
		}
	}
	return
}

func swapIteration(grid [][]string, points [][]int, replace string) {
	for _, point := range points {
		r := point[0]
		c := point[1]
		grid[r][c] = replace
	}
}

func validInput(grid [][]string) (valid bool) {
	valid = false
	for r := range grid {
		for c := range grid {
			if grid[r][c] == "L" && seatShouldBeFilled(grid, r, c) {
				return
			}

			if grid[r][c] == "#" && seatShouldBeEmpty(grid, r, c) {
				return
			}
		}
	}
	return true
}

func seatShouldBeFilled(grid [][]string, rIndex int, cIndex int) (valid bool) {
	valid = true
	if grid[rIndex][cIndex] != "L" {
		return false
	}

	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			adjacentR := rIndex + r
			validR := adjacentR >= 0 && adjacentR < len(grid)
			adjacentC := cIndex + c
			validC := adjacentC >= 0 && adjacentC < len(grid[0])

			if validR && validC && grid[adjacentR][adjacentC] == "#" {
				return false
			}
		}
	}
	return
}

func seatShouldBeEmpty(grid [][]string, rIndex int, cIndex int) (valid bool) {
	valid = false
	if grid[rIndex][cIndex] != "#" {
		return
	}

	count := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			adjacentR := rIndex + r
			validR := adjacentR >= 0 && adjacentR < len(grid)
			adjacentC := cIndex + c
			validC := adjacentC >= 0 && adjacentC < len(grid[0])

			if validR && validC && grid[adjacentR][adjacentC] == "#" {
				count++
				if count > 3 {
					return true
				}
			}
		}
	}
	return
}

func readInput(delimiter string) [][]string {
	input, _ := ioutil.ReadFile("day_11.txt")
	splitInput := strings.Split(string(input), delimiter)
	inputGrid := [][]string{}
	for _, row := range splitInput {
		s := strings.Split(row, "")
		inputGrid = append(inputGrid, s)
	}
	return inputGrid
}
