package utils

import (
	"os"
	"regexp"
)

func GetPath() string {
	const projectDirName = "golang-mpesa"

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}
