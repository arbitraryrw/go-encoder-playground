package xmlhandler

import (
	"encoding/xml"
	"fmt"
	"goencoderplayground/internal/pkg/utils"
)

//ref: https://golang.org/pkg/encoding/xml/
// Node is a structure for convinient parsing of XML
type Node struct {
	Attr     []xml.Attr
	XMLName  xml.Name
	Children []Node `xml:",any"`
	Text     string `xml:",chardata"`
}

// Widget is a structure for convinient parsing of XML
// generated using https://www.onlinetool.io/xmltogo/
type Widget struct {
	XMLName  xml.Name `xml:"widget"`
	Chardata string   `xml:",chardata"`
	Debug    string   `xml:"debug"`
	Window   struct {
		Text   string `xml:",chardata"`
		Title  string `xml:"title,attr"`
		Name   string `xml:"name"`
		Width  string `xml:"width"`
		Height string `xml:"height"`
	} `xml:"window"`
	Image struct {
		Text      string `xml:",chardata"`
		Src       string `xml:"src,attr"`
		Name      string `xml:"name,attr"`
		HOffset   string `xml:"hOffset"`
		VOffset   string `xml:"vOffset"`
		Alignment string `xml:"alignment"`
	} `xml:"image"`
	Text struct {
		Text      string `xml:",chardata"`
		Data      string `xml:"data,attr"`
		Size      string `xml:"size,attr"`
		Style     string `xml:"style,attr"`
		Name      string `xml:"name"`
		HOffset   string `xml:"hOffset"`
		VOffset   string `xml:"vOffset"`
		Alignment string `xml:"alignment"`
		OnMouseUp string `xml:"onMouseUp"`
	} `xml:"text"`
}

//UnmarshalXMLData converts bytevalue to map
func UnmarshalXMLData(path string) (Node, error) {
	byteValue, err := utils.ReadRawFileContents(path)

	if err != nil {
		return Node{}, err
	}

	data := Node{}

	if err := xml.Unmarshal(byteValue, &data); err != nil {
		return Node{}, err
	}

	return data, nil
}

//UnmarshalStructuredXMLData converts bytevalue to map
func UnmarshalStructuredXMLData(path string) (Widget, error) {
	byteValue, err := utils.ReadRawFileContents(path)

	if err != nil {
		return Widget{}, err
	}

	data := Widget{}

	if err := xml.Unmarshal(byteValue, &data); err != nil {
		return Widget{}, err
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
