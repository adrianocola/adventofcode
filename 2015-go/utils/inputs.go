package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) (string, error) {
	normalizedFilePath := strings.TrimSuffix(filePath, "/")
	data, err := os.ReadFile("../inputs/2015/" + normalizedFilePath)
	if err != nil {
		return "", err
	}

	// Normalize Windows line endings (\r\n) if needed, and trim a trailing newline
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.TrimSuffix(content, "\n")

	return content, err
}

// Reads input string relative to this year input folder
func ReadInput(filePath string) string {
	data, err := readFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Reads input lines relative to this year input folder
func ReadInputLines(filePath string) []string {
	data, err := readFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(data, "\n")
}

// Reads result relative to this year input folder
func ReadResult(filePath string) (int, error) {
	data, err := readFile(filePath)
	result, err := strconv.ParseInt(data, 10, 0)
	return int(result), err
}
