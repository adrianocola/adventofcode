package main

import (
	"aoc2015/utils"
)

var up rune = '('

func main() {
	input := utils.ReadInput("01/input.txt")

	utils.CheckResult("01", "1", func() int { return step1(input) })
	utils.CheckResult("01", "2", func() int { return step2(input) })
}

func step1(input string) int {
	var counter int
	for _, v := range input {
		if v == up {
			counter++
		} else {
			counter--
		}
	}

	return counter
}

func step2(input string) int {
	var counter int
	for i, v := range input {
		if v == up {
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
