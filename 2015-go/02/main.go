package main

import (
	"aoc2015/utils"
	"slices"
	"strings"
)

func main() {
	input := utils.ReadInputLines("02/input.txt")

	utils.CheckResult("02", "1", func() int { return step1(input) })
	utils.CheckResult("02", "2", func() int { return step2(input) })
}

func step1(boxesLines []string) int {
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

func step2(boxesLines []string) int {
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
