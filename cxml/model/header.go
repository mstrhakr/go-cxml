package model

import "encoding/xml"

type Party struct {
	Credential    *Credential    `xml:"Credential,omitempty"`
	Identity      string         `xml:"Identity,omitempty"`
	Contact       *Contact       `xml:"Contact,omitempty"`
	Name          *Name          `xml:"Name,omitempty"`
	PostalAddress *PostalAddress `xml:"PostalAddress,omitempty"`
}

type Sender struct {
	Credential *Credential `xml:"Credential,omitempty"`
	UserAgent  string      `xml:"UserAgent,omitempty"`
}

type Credential struct {
	XMLName      xml.Name `xml:"Credential"`
	Domain       string   `xml:"domain,attr,omitempty"`
	Identity     string   `xml:"Identity,omitempty"`
	SharedSecret string   `xml:"SharedSecret,omitempty"`
}

type Name struct {
	XMLName xml.Name `xml:"Name"`
	Value   string   `xml:",chardata"`
}

type Contact struct {
	XMLName xml.Name `xml:"Contact"`
	Name    string   `xml:"Name,omitempty"`
	Email   string   `xml:"Email,omitempty"`
}

type PostalAddress struct {
	XMLName    xml.Name `xml:"PostalAddress"`
	Street1    string   `xml:"Street1,omitempty"`
	City       string   `xml:"City,omitempty"`
	State      string   `xml:"State,omitempty"`
	PostalCode string   `xml:"PostalCode,omitempty"`
	Country    *Country `xml:"Country,omitempty"`
}

type Country struct {
	XMLName xml.Name `xml:"Country"`
	Code    string   `xml:"isoCountryCode,attr,omitempty"`
	Name    string   `xml:",chardata"`
}
