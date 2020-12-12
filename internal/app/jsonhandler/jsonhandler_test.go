package jsonhandler_test

import (
	"errors"
	"fmt"
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

	for k, _ := range widget.(map[string]interface{}) {
		if !utils.ContainsKey(expectedKeys, k) {
			t.Errorf("Unable to find expected value %q in expected keys %q",
				k,
				expectedKeys)
		}
	}
}

func TestUnmarshalStructedJSONDataPositive(t *testing.T) {
	file := "basic.json"
	err := validateStructuredDataFormat(file)

	if err != nil {
		t.Errorf("UnmarshalStructedJSONData() error handinling file %q, error %q",
			file,
			err)
	}
}

func TestUnmarshalStructedJSONDataNegative(t *testing.T) {
	negativeFiles, err := utils.GetTestFile("negative")

	if len(negativeFiles) < 1 {
		t.Errorf("GetTestFile() returned %q results, expected at least 1",
			len(negativeFiles))
	}

	if err != nil {
		t.Errorf("GetTestFile() error handinling file %q, error %q",
			negativeFiles,
			err)
	}

	for _, file := range negativeFiles {
		err := validateStructuredDataFormat(file)

		if err == nil {
			t.Errorf("UnmarshalStructedJSONData() error handinling file %q, error %q",
				file,
				err)
		}
	}

}

func validateStructuredDataFormat(testFileName string) error {
	testFilesSlice, err := utils.GetTestFile(testFileName)

	if err != nil {
		errorText := fmt.Sprintf("GetTestFile() unable to find test file error: %q",
			testFilesSlice)

		return errors.New(errorText)
	}

	if len(testFilesSlice) < 1 {
		errorText := fmt.Sprintf("GetTestFile() less than 1 match, expected 1 got %q",
			len(testFilesSlice))

		return errors.New(errorText)
	}

	widgetStructure, err := jsonhandler.UnmarshalStructuredJSONData(testFilesSlice[0])

	if err != nil {
		errorText := fmt.Sprintf("UnmarshalStructuredJSONData() unmarshal file %q, got error: %q",
			testFileName,
			testFilesSlice)

		return errors.New(errorText)
	}

	if widgetStructure.Widget.Debug == "" {
		errorText := fmt.Sprintf("UnmarshalStructuredJSONData() unable to find property %q for file %q",
			"widgetStructure.Widget.Debug",
			testFileName)

		return errors.New(errorText)
	}

	return nil
}

func validateUnstructuredDataFormat(testFileName string) error {
	//ToDo
	return nil
}
