package utils

import (
	"log"
	"path"
	"runtime/debug"
)

func GetExecDay() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("No build info")
	}

	return path.Base(info.Path)
}
