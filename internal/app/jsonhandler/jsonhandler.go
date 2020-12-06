package jsonhandler

import (
	"encoding/json"
	"fmt"
	"goencoderplayground/internal/pkg/utils"
)

// BasicStructuredJSON test data structure for basic.json
// Generated using https://transform.tools/json-to-go
type BasicStructuredJSON struct {
	Widget struct {
		Debug  string `json:"debug"`
		Window struct {
			Title  string `json:"title"`
			Name   string `json:"name"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"window"`
		Image struct {
			Src       string `json:"src"`
			Name      string `json:"name"`
			HOffset   int    `json:"hOffset"`
			VOffset   int    `json:"vOffset"`
			Alignment string `json:"alignment"`
		} `json:"image"`
		Text struct {
			Data      string `json:"data"`
			Size      int    `json:"size"`
			Style     string `json:"style"`
			Name      string `json:"name"`
			HOffset   int    `json:"hOffset"`
			VOffset   int    `json:"vOffset"`
			Alignment string `json:"alignment"`
			OnMouseUp string `json:"onMouseUp"`
		} `json:"text"`
	} `json:"widget"`
}

//ref: https://golang.org/pkg/encoding/json/

//UnmarshalJSONData converts bytevalue to map
func UnmarshalJSONData(path string) (interface{}, error) {
	byteValue, err := utils.ReadRawFileContents(path)

	if err != nil {
		return nil, err
	}

	var data interface{}

	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, err
	}

	return data, nil
}

//UnmarshalStructuredJSONData converts bytevalue to a struct
func UnmarshalStructuredJSONData(path string) (BasicStructuredJSON, error) {
	byteValue, err := utils.ReadRawFileContents(path)

	if err != nil {
		return BasicStructuredJSON{}, err
	}

	retStruct := BasicStructuredJSON{}

	if err := json.Unmarshal(byteValue, &retStruct); err != nil {
		return BasicStructuredJSON{}, err
	}

	return retStruct, nil
}

//PrettyPrintJSON pretty prints an unmarshalled object
func PrettyPrintJSON(rawData interface{}) error {
	prettyData, err := json.MarshalIndent(rawData, "", "  ")

	if err != nil {
		return err
	}

	fmt.Println(string(prettyData))
	return nil
}
