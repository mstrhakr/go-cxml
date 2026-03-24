package model

import "encoding/xml"

type Message struct {
	XMLName xml.Name        `xml:"Message"`
	Subject string          `xml:"Subject,omitempty"`
	Payload *PayloadWrapper `xml:"Payload,omitempty"`
}

type PayloadWrapper struct {
	XMLName xml.Name `xml:"Payload"`
	Content string   `xml:",innerxml"`
}
