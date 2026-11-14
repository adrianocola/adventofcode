package main

import (
	"aoc2015/utils"
	"crypto/md5"
	"fmt"
)

func main() {
	input := utils.ReadInput()

	utils.RunPart(1, func() int { return part1(input) })
	utils.RunPart(2, func() int { return part2(input) })
}

func part1(input string) int {
	var i int

	for {
		text := fmt.Sprintf("%s%d", input, i)
		hash := md5.Sum([]byte(text))

		if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
			return i
		}

		i++
	}
}

func part2(input string) int {
	var i int

	for {
		text := fmt.Sprintf("%s%d", input, i)
		hash := md5.Sum([]byte(text))

		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			return i
		}

		i++
	}
}
