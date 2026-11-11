package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) string {
	normalizedFilePath := strings.TrimSuffix(filePath, "/")
	data, err := os.ReadFile("../inputs/2015/" + normalizedFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Normalize Windows line endings (\r\n) if needed, and trim a trailing newline
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.TrimSuffix(content, "\n")

	return content
}

// Reads input string relative to this year input folder
func ReadInput(filePath string) string {
	return readFile(filePath)
}

// Reads input lines relative to this year input folder
func ReadInputLines(filePath string) []string {
	content := readFile(filePath)

	return strings.Split(content, "\n")
}

// Reads result relative to this year input folder
func ReadResult(filePath string) int {
	content := readFile(filePath)
	result, err := strconv.ParseInt(content, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	return int(result)
}
