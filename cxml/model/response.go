package model

import "encoding/xml"

type Response struct {
	XMLName xml.Name `xml:"Response"`
	Status  *Status  `xml:"Status,omitempty"`
}

// Status is used for both Response and top-level success/failure.
type Status struct {
	XMLName xml.Name `xml:"Status"`
	Code    string   `xml:"code,attr,omitempty"`
	Text    string   `xml:"text,attr,omitempty"`
}
