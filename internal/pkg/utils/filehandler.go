package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//ReadRawFileContents reads a file from an arbitrary path and returns the contents
func ReadRawFileContents(path string) ([]byte, error) {

	if !fileExists(path) {
		return nil, errors.New("File does not exist")
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
