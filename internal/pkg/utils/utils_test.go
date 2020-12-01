package utils_test

import (
	"goencoderplayground/internal/pkg/utils"
	"os"
	"strings"
	"testing"
)

func TestGetProjectRootDir(t *testing.T) {
	rootDir, err := utils.GetProjectRootDir()

	if err != nil {
		t.Errorf(
			"getProjectRootDir returned error %q",
			err)
	}

	splitDir := strings.Split(rootDir, string(os.PathSeparator))

	if len(splitDir) < 1 {
		t.Errorf(
			"getProjectRootDir() split path length of %q for %q less than 1",
			len(splitDir),
			rootDir)
	}

	if splitDir[len(splitDir)-1] != "go-encoder-playground" {
		t.Errorf(
			"getProjectRootDir() base dir of %q does not equal %q",
			splitDir[len(splitDir)-1],
			"go-encoder-playground")
	}
}

func TestGetTestFilePositive(t *testing.T) {
	testFileSlice, err := utils.GetTestFile(".json")

	if err != nil {
		t.Errorf("GetTestFile() unable to find test file error: %q",
			testFileSlice)
	}

	if len(testFileSlice) < 1 {
		t.Errorf("GetTestFile() less than 1 match, expected 1 got %q",
			len(testFileSlice))
	}

	for _, value := range testFileSlice {
		if len(value) < 4 {
			t.Errorf("GetTestFile() file length less than 4 characters, got %q",
				len(value))
		}

		if value[len(value)-4:] != "json" {
			t.Errorf("GetTestFile() returned a file %q  %q",
				value[len(value)-4:],
				"json")
		}
	}
}

func TestGetTestFileNegative(t *testing.T) {
	testFileSlice, err := utils.GetTestFile("randomFileThatShouldNotExist")

	if err != nil {
		t.Errorf("GetTestFile() unable to find test file error: %q",
			testFileSlice)
	}

	if len(testFileSlice) > 0 {
		t.Errorf("GetTestFile() less than 1 match, expected 1 got %q",
			len(testFileSlice))
	}
}
