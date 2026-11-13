package main

import (
	"aoc2015/utils"
)

func main() {
	input := utils.ReadInput()

	utils.RunPart(1, func() int { return part1(input) })
	utils.RunPart(2, func() int { return part2(input) })
}

func part1(input string) int {
	var counter int
	for _, v := range input {
		if v == '(' {
			counter++
		} else {
			counter--
		}
	}

	return counter
}

func part2(input string) int {
	var counter int
	for i, v := range input {
		if v == '(' {
			counter++
		} else {
			counter--
		}

		if counter <= -1 {
			return i + 1
		}
	}

	return 0
}
