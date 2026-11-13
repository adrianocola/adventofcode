package utils

import (
	"fmt"
	"time"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	bold   = "\033[1m"
	faint  = "\033[2m"
)

func RunPart[T int | string](part int, runner func() T) {
	start := time.Now()
	result := runner()
	elapsed := time.Since(start)

	day := GetExecDay()

	expected := ReadResult(day, part)
	resultStr := fmt.Sprint(result)

	prefix := fmt.Sprintf("Day %s Part %d:", day, part)
	var suffix string

	if resultStr == expected {
		suffix = green + "✔︎" + reset + faint + " (" + resultStr + ")" + reset
	} else {
		suffix = red + "X" + reset + faint + " (got " + resultStr + ", expected: " + expected + ")" + reset
	}

	fmt.Printf("%s %s %v\n", prefix, suffix, elapsed)
}
