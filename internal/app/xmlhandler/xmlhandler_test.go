package xmlhandler_test

import (
	"goencoderplayground/internal/app/xmlhandler"
	"goencoderplayground/internal/pkg/utils"
	"testing"
)

func TestUnmarshalXMLDataPositive(t *testing.T) {
	testFileName := "basic.xml"

	testFilesSlice, err := utils.GetTestFile(testFileName)

	if err != nil {
		t.Errorf("GetTestFile() unable to find test file error: %q",
			testFilesSlice)
	}

	if len(testFilesSlice) < 1 {
		t.Errorf("GetTestFile() less than 1 match, expected 1 got %q",
			len(testFilesSlice))
	}

	_, err := xmlhandler.UnmarshalXMLData(testFilesSlice[0])

	if err != nil {
		t.Errorf("UnmarshalXMLData() unmarshal file %q, got error: %q",
			testFileName,
			testFilesSlice)
	}

}
