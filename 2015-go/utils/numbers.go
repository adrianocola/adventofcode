package utils

import (
	"log"
	"strconv"
)

func ParseNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return num
}
