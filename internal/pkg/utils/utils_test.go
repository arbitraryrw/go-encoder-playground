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
