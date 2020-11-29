package xmlhandler

import (
	"encoding/xml"
	"fmt"
	"goencoderplayground/internal/pkg/utils"
)

//ref: https://golang.org/pkg/encoding/xml/

type node struct {
	Attr     []xml.Attr
	XMLName  xml.Name
	Children []node `xml:",any"`
	Text     string `xml:",chardata"`
}

//UnmarshalXMLData converts bytevalue to map
func UnmarshalXMLData(path string) (interface{}, error) {
	byteValue, err := utils.ReadRawFileContents(path)

	if err != nil {
		return nil, err
	}

	data := node{}

	if err := xml.Unmarshal(byteValue, &data); err != nil {
		return nil, err
	}

	return data, nil
}

//PrettyPrintXML pretty prints an unmarshalled object
func PrettyPrintXML(rawData interface{}) error {
	buf, err := xml.MarshalIndent(rawData, "", "\t")

	if err != nil {
		return err
	}

	fmt.Println(string(buf))
	return nil
}
