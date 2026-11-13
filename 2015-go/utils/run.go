package utils

import (
	"fmt"
	"strconv"
	"time"
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

func convertResultToString[T int | string](result T) string {
	switch v := any(result).(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	default:
		return ""
	}
}

func RunPart[T int | string](part int, runner func() T) {
	start := time.Now()
	result := runner()
	elapsed := time.Since(start).Milliseconds()

	day := GetExecDay()

	expected := ReadResult(day, part)
	resultStr := convertResultToString(result)

	prefix := fmt.Sprintf("Day %s Part %d:", day, part)
	var suffix string

	if resultStr == expected {
		suffix = Green + "✔︎" + Reset + Faint + " (" + resultStr + ")"
	} else {
		suffix = Red + "X" + Reset + Faint + " (got " + resultStr + ", expected: " + expected + ")"
	}

	fmt.Printf("%s %s %dms\n", prefix, suffix, elapsed)
}
