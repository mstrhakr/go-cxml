package model

import "encoding/xml"

// CXML is the root element of every cXML document.
type CXML struct {
	XMLName          xml.Name `xml:"cXML"`
	PayloadID        string   `xml:"payloadID,attr,omitempty"`
	Timestamp        string   `xml:"timestamp,attr,omitempty"`
	Version          string   `xml:"version,attr,omitempty"`
	SignatureVersion string   `xml:"signatureVersion,attr,omitempty"`
	Lang             string   `xml:"xml:lang,attr,omitempty"`

	From   *Party  `xml:"Header>From,omitempty"`
	To     *Party  `xml:"Header>To,omitempty"`
	Sender *Sender `xml:"Header>Sender,omitempty"`

	Request  *Request  `xml:"Request,omitempty"`
	Response *Response `xml:"Response,omitempty"`
	Message  *Message  `xml:"Message,omitempty"`

	Status *Status `xml:"Status,omitempty"`
}

func (c *CXML) GetPayloadType() string {
	switch {
	case c.Request != nil:
		return c.Request.PayloadType()
	case c.Response != nil:
		return c.Response.PayloadType()
	case c.Message != nil:
		return c.Message.PayloadType()
	default:
		return ""
	}
}

func (c *CXML) IsRequest() bool {
	return c.Request != nil
}

func (c *CXML) IsResponse() bool {
	return c.Response != nil
}

func (c *CXML) IsMessage() bool {
	return c.Message != nil
}
