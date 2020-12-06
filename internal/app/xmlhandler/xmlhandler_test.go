package xmlhandler_test

import (
	"fmt"
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

	xmlData, err := xmlhandler.UnmarshalXMLData(testFilesSlice[0])

	if err != nil {
		t.Errorf("UnmarshalXMLData() unmarshal file %q, got error: %q",
			testFileName,
			testFilesSlice)
	}

	if xmlData.Children[1].Children[1].XMLName.Local != "width" {
		t.Errorf("UnmarshalXMLData() parmeter missmatch, could not find %q in %q",
			"width",
			xmlData.Children[1].Children[1])
	}

	if xmlData.Children[1].Children[1].Text != "500" {
		fmt.Println("yeet")
		t.Errorf("UnmarshalXMLData() parmeter missmatch, could not find %q in %q",
			"500",
			xmlData.Children[1].Children[1])
	}

	expectedKeys := []string{
		"window",
		"image",
		"text",
		"debug",
	}

	for _, v := range xmlData.Children {
		if !utils.ContainsKey(expectedKeys, v.XMLName.Local) {
			t.Errorf("UnmarshalXMLData() unable to find expected value %q in expected keys %q",
				v.XMLName.Local,
				expectedKeys)
		}
	}

}
