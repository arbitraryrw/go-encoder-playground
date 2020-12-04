package main

import (
	"fmt"
	"goencoderplayground/internal/app/jsonhandler"
	"goencoderplayground/internal/app/xmlhandler"
	"goencoderplayground/internal/pkg/utils"
)

func init() {}

func main() {
	fmt.Println("go-encoder-playground running...")

	filePaths, _ := utils.GetTestFile(".json")

	for index, value := range filePaths {
		fmt.Printf("Test JSON file %b is %q\n", index, value)

		json, _ := jsonhandler.UnmarshalJSONData(value)

		jsonhandler.PrettyPrintJSON(json)
	}

	filePaths, _ = utils.GetTestFile(".xml")

	for index, value := range filePaths {
		fmt.Printf("Test XML file %b is %q\n", index, value)

		xml, _ := xmlhandler.UnmarshalXMLData(value)

		xmlhandler.PrettyPrintXML(xml)
	}

}
