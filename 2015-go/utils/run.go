package utils

import (
	"fmt"
	"strconv"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Bold   = "\033[1m"
	Faint  = "\033[2m"
)

func CheckResult(day string, step string, result int, expectedResult int) {
	prefix := "Day " + day + " Step " + step + ":"
	var suffix string

	if result == expectedResult {
		suffix = Green + "✔︎" + Reset + Faint + " (" + strconv.Itoa(result) + ")"
	} else {
		suffix = Red + "X" + Reset + Faint + " (got " + strconv.Itoa(result) + ", expected: " + strconv.Itoa(expectedResult) + ")"
	}

	fmt.Println(prefix + " " + suffix)
}
