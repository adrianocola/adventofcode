package main

import (
	"aoc2015/utils"
	"slices"
	"strings"
)

func main() {
	input := utils.ReadInputLines()

	utils.RunPart(1, func() int { return part1(input) })
	utils.RunPart(2, func() int { return part2(input) })
}

func part1(boxesLines []string) int {
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

func part2(boxesLines []string) int {
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
