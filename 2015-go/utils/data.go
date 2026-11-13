package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filePath string) (string, error) {
	normalizedFilePath := strings.TrimSuffix(filePath, "/")
	data, err := os.ReadFile("../data/2015/" + normalizedFilePath)
	if err != nil {
		return "", err
	}

	// Normalize Windows line endings (\r\n) if needed, and trim a trailing newline
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.TrimSuffix(content, "\n")

	return content, nil
}

func readInputFile() string {
	day := GetExecDay()
	data, err := readFile(day + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ReadInput() string {
	return readInputFile()
}

func ReadInputLines() []string {
	data := readInputFile()

	return strings.Split(data, "\n")
}

func ReadResult(day string, part int) string {
	filePath := fmt.Sprintf("%s/part%d.txt", day, part)
	result, err := readFile(filePath)
	if err != nil {
		return "<missing>"
	}
	return result
}
