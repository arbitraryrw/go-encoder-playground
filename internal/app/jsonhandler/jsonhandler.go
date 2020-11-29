package jsonhandler

import (
	"encoding/json"
	"fmt"
	"goencoderplayground/internal/pkg/utils"
)

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

//PrettyPrintJSON pretty prints an unmarshalled object
func PrettyPrintJSON(rawData interface{}) error {
	prettyData, err := json.MarshalIndent(rawData, "", "  ")

	if err != nil {
		return err
	}

	fmt.Println(string(prettyData))
	return nil
}
