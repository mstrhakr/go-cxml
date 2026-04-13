package model

import "encoding/xml"

type Message struct {
	XMLName xml.Name        `xml:"Message"`
	Subject string          `xml:"Subject,omitempty"`
	Payload *PayloadWrapper `xml:"Payload,omitempty"`
}

func (m *Message) PayloadType() string {
	if m == nil {
		return ""
	}
	if m.Payload != nil {
		return "Payload"
	}
	return "Message"
}

type PayloadWrapper struct {
	XMLName xml.Name `xml:"Payload"`
	Content string   `xml:",innerxml"`
}
