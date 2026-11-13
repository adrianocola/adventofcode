package utils

import (
	"log"
	"runtime/debug"
	"strings"
)

func GetExecDay() string {
	info, ok := debug.ReadBuildInfo()
	if ok == false {
		log.Fatal("No build info")
	}

	splittedPath := strings.Split(info.Path, "/")
	day := splittedPath[len(splittedPath)-1]

	return day
}
