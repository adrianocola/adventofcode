package main

import (
	"aoc2015/utils"
	"slices"
	"strings"
)

func main() {
	input := utils.ReadInputLines("02/input.txt")

	expectedResult1 := utils.ReadResult("02/result1.txt")
	result1 := solution1(input)
	utils.CheckResult("02", "1", result1, expectedResult1)

	expectedResult2 := utils.ReadResult("02/result2.txt")
	result2 := solution2(input)
	utils.CheckResult("02", "2", result2, expectedResult2)
}

func solution1(boxesLines []string) int {
	var count int
	for _, v := range boxesLines {
		parts := strings.Split(v, "x")
		l, w, h := utils.ParseNumber(parts[0]), utils.ParseNumber(parts[1]), utils.ParseNumber(parts[2])
		lw := l * w
		wh := w * h
		hl := h * l

		smaller := min(lw, wh, hl)

		count += 2*lw + 2*wh + 2*hl + smaller
	}

	return count
}

func solution2(boxesLines []string) int {
	var count int
	for _, v := range boxesLines {
		parts := strings.Split(v, "x")
		l, w, h := utils.ParseNumber(parts[0]), utils.ParseNumber(parts[1]), utils.ParseNumber(parts[2])
		partsInt := []int{l, w, h}
		slices.Sort(partsInt)
		s1, s2 := partsInt[0], partsInt[1]

		count += 2*s1 + 2*s2 + l*w*h
	}

	return count
}
