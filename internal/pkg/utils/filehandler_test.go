package utils_test

import (
	"goencoderplayground/internal/pkg/utils"
	"testing"
)

func TestFileExists(t *testing.T) {
	filePath, err := utils.GetTestFile("basic.json")

	if err != nil {
		t.Errorf(
			"TestFileExists() Unable to get GetTestFile(), received error %q ",
			err)
	}

	if len(filePath) < 1 {
		t.Errorf(
			"TestFileExists() Unable to get GetTestFile(), found %q  file, expected %q",
			len(filePath),
			1)
	}

	pathExists := utils.FileExists(filePath[0])

	if !pathExists {
		t.Errorf(
			"TestFileExists() returned %t for file %q, expected %t",
			pathExists,
			filePath[0],
			true)
	}

	if utils.FileExists("") {
		t.Errorf(
			"TestFileExists() returned %t for file %q, expected %t",
			pathExists,
			filePath[0],
			false)
	}
}
