package utils

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

//returns an absolute path to the project root dir
func getProjectRootDir() (string, error) {
	var _, b, _, _ = runtime.Caller(0)
	var basePath = filepath.Dir(b)
	splitPath := strings.Split(basePath, string(os.PathSeparator))

	tempDir := "/"

	for _, dir := range splitPath {

		tempDir = path.Join(tempDir, dir)

		if dir == "go-encoder-playground" {
			break
		}
	}

	baseDir := strings.Split(tempDir, string(os.PathSeparator))

	if len(baseDir) < 1 {
		return "", errors.New("Unable to find root dir")
	}

	if baseDir[len(baseDir)-1] != "go-encoder-playground" {
		return "", errors.New("Could not find project root directory")
	}

	return tempDir, nil
}

//GetTestFile returns an array of absolute paths
func GetTestFile(partialOrFullName string) ([]string, error) {
	var ruleFiles []string

	projectRootDir, pathErr := getProjectRootDir()

	if pathErr != nil {
		return nil, pathErr
	}

	ruleDir := path.Join(projectRootDir, "test/")

	err := filepath.Walk(ruleDir, func(path string, info os.FileInfo, err error) error {

		if strings.Contains(filepath.Base(path), partialOrFullName) {
			ruleFiles = append(ruleFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ruleFiles, nil
}

//ContainsKey checks if a value present in a slice
func ContainsKey(haystack []string, needle string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}
	return false
}
