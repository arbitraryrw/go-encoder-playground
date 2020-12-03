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

func TestReadRawFileContents(t *testing.T) {
	filePaths, err := utils.GetTestFile("basic.json")

	if err != nil {
		t.Errorf(
			"TestFileExists() Unable to get GetTestFile(), received error %q ",
			err)
	}

	if len(filePaths) < 1 {
		t.Errorf(
			"TestFileExists() Unable to get GetTestFile(), found %q  file, expected %q",
			len(filePaths),
			1)
	}

	fileHandle, err := utils.ReadRawFileContents(filePaths[0])

	if err != nil {
		t.Errorf(
			"ReadRawFileContents() Unable to get ReadRawFileContents(), received error %q ",
			err)
	}

	if len(fileHandle) < 1 {
		t.Errorf(
			"ReadRawFileContents(), byte length %q received for file %q. Expected >0 length.",
			len(fileHandle),
			filePaths[0])
	}
}
