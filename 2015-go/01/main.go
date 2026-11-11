package main

import (
	"aoc2015/utils"
)

var up rune = '('

func main() {
	input := utils.ReadInput("01/input.txt")
	expectedResult1 := utils.ReadResult("01/result1.txt")
	result1 := solution1(input)
	utils.CheckResult("01", "1", result1, expectedResult1)

	expectedResult2 := utils.ReadResult("01/result2.txt")
	result2 := solution2(input)
	utils.CheckResult("01", "2", result2, expectedResult2)
}

func solution1(input string) int {
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

func solution2(input string) int {
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
