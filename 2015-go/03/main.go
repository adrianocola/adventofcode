package main

import (
	"aoc2015/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput()

	utils.RunPart(1, func() int { return part1(input) })
	utils.RunPart(2, func() int { return part2(input) })
}

func getPosStr(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func part1(input string) int {
	houses := map[string]bool{}

	r := 1
	x, y := 0, 0
	houses[getPosStr(x, y)] = true
	for _, c := range input {
		switch c {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y--
		case 'v':
			y++
		}
		pos := getPosStr(x, y)
		if !houses[pos] {
			r++
			houses[pos] = true
		}
	}

	return r
}

type Pos struct {
	x int
	y int
}

func part2(input string) int {
	houses := map[string]bool{}

	r := 1
	posSanta := Pos{}
	posRobot := Pos{}
	houses[getPosStr(posSanta.x, posSanta.y)] = true
	for i, c := range input {
		currentPos := &posSanta
		if i%2 != 0 {
			currentPos = &posRobot
		}
		switch c {
		case '<':
			currentPos.x--
		case '>':
			currentPos.x++
		case '^':
			currentPos.y--
		case 'v':
			currentPos.y++
		}
		pos := getPosStr(currentPos.x, currentPos.y)
		if !houses[pos] {
			r++
			houses[pos] = true
		}
	}

	return r
}
