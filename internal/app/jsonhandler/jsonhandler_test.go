package jsonhandler_test

import (
	"goencoderplayground/internal/app/jsonhandler"
	"goencoderplayground/internal/pkg/utils"
	"testing"
)

func TestUnmarshalJSONDataPositive(t *testing.T) {

	testFileName := "basic.json"

	testFilesSlice, err := utils.GetTestFile(testFileName)

	if err != nil {
		t.Errorf("GetTestFile() unable to find test file error: %q",
			testFilesSlice)
	}

	if len(testFilesSlice) < 1 {
		t.Errorf("GetTestFile() less than 1 match, expected 1 got %q",
			len(testFilesSlice))
	}

	json, err := jsonhandler.UnmarshalJSONData(testFilesSlice[0])

	if err != nil {
		t.Errorf("UnmarshalJSONData() unmarshal file %q, got error: %q",
			testFileName,
			testFilesSlice)
	}

	widget := json.(map[string]interface{})["widget"]
	widgetImage := widget.(map[string]interface{})["image"]
	widgetImageName := widgetImage.(map[string]interface{})["name"]

	expectedValue := "sun1"

	if widgetImageName != expectedValue {
		t.Errorf("Unable to find expected value %q in json response %q",
			expectedValue,
			widgetImageName)
	}

	expectedKeys := []string{
		"window",
		"image",
		"text",
		"debug",
	}

	for k, v := range widget.(map[string]interface{}) {
		if !containsKey(expectedKeys, k) {
			t.Errorf("Unable to find expected value %q in expected keys %q",
				k,
				expectedKeys)
		}
	}
}

func containsKey(haystack []string, needle string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}
	return false
}
