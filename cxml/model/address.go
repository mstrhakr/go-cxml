package model

import "encoding/xml"

// ─── Basic address elements ──────────────────────────────────────────────────

// CountryCode is a telephone dialing code.
type CountryCode struct {
	XMLName        xml.Name `xml:"CountryCode"`
	IsoCountryCode string   `xml:"isoCountryCode,attr,omitempty"`
	Value          string   `xml:",chardata"`
}

// AreaOrCityCode is the area/city prefix of a telephone number.
type AreaOrCityCode struct {
	XMLName xml.Name `xml:"AreaOrCityCode"`
	Value   string   `xml:",chardata"`
}

// PhoneNumber is the subscriber portion of a telephone number.
type PhoneNumber struct {
	XMLName xml.Name `xml:"Number"`
	Value   string   `xml:",chardata"`
}

// Extension is a telephone extension number.
type Extension struct {
	XMLName xml.Name `xml:"Extension"`
	Value   string   `xml:",chardata"`
}

// TelephoneNumber is a structured phone number.
type TelephoneNumber struct {
	XMLName        xml.Name        `xml:"TelephoneNumber"`
	CountryCode    *CountryCode    `xml:"CountryCode,omitempty"`
	AreaOrCityCode *AreaOrCityCode `xml:"AreaOrCityCode,omitempty"`
	Number         *PhoneNumber    `xml:"Number,omitempty"`
	Extension      *Extension      `xml:"Extension,omitempty"`
}

// Phone wraps a TelephoneNumber with a name attribute.
type Phone struct {
	XMLName         xml.Name         `xml:"Phone"`
	Name            string           `xml:"name,attr,omitempty"`
	TelephoneNumber *TelephoneNumber `xml:"TelephoneNumber"`
}

// Fax is a fax contact (can be a telephone number, URL, or email).
type Fax struct {
	XMLName         xml.Name         `xml:"Fax"`
	Name            string           `xml:"name,attr,omitempty"`
	TelephoneNumber *TelephoneNumber `xml:"TelephoneNumber,omitempty"`
	URL             *URL             `xml:"URL,omitempty"`
	Email           *Email           `xml:"Email,omitempty"`
}

// Email holds an email address.
type Email struct {
	XMLName       xml.Name `xml:"Email"`
	Name          string   `xml:"name,attr,omitempty"`
	PreferredLang string   `xml:"preferredLang,attr,omitempty"`
	Value         string   `xml:",chardata"`
}

// ─── Address elements ────────────────────────────────────────────────────────

// DeliverTo is a named care-of delivery addressee.
type DeliverTo struct {
	XMLName xml.Name `xml:"DeliverTo"`
	Value   string   `xml:",chardata"`
}

// Street is one line of a street address.
type Street struct {
	XMLName xml.Name `xml:"Street"`
	Value   string   `xml:",chardata"`
}

// City holds a city name.
type City struct {
	XMLName  xml.Name `xml:"City"`
	CityCode string   `xml:"cityCode,attr,omitempty"`
	Value    string   `xml:",chardata"`
}

// State holds a state or province code/name.
type State struct {
	XMLName      xml.Name `xml:"State"`
	IsoStateCode string   `xml:"isoStateCode,attr,omitempty"`
	Value        string   `xml:",chardata"`
}

// Municipality is a municipality within a city.
type Municipality struct {
	XMLName          xml.Name `xml:"Municipality"`
	MunicipalityCode string   `xml:"municipalityCode,attr,omitempty"`
	Value            string   `xml:",chardata"`
}

// PostalCode is a postal/ZIP code.
type PostalCode struct {
	XMLName xml.Name `xml:"PostalCode"`
	Value   string   `xml:",chardata"`
}

// Country holds a country code and optional name.
type Country struct {
	XMLName        xml.Name `xml:"Country"`
	IsoCountryCode string   `xml:"isoCountryCode,attr,omitempty"`
	Value          string   `xml:",chardata"`
}

// PostalAddress is a structured postal delivery address.
type PostalAddress struct {
	XMLName      xml.Name      `xml:"PostalAddress"`
	Name         string        `xml:"name,attr,omitempty"`
	DeliverTo    []*DeliverTo  `xml:"DeliverTo,omitempty"`
	Streets      []*Street     `xml:"Street"`
	City         *City         `xml:"City,omitempty"`
	Municipality *Municipality `xml:"Municipality,omitempty"`
	State        *State        `xml:"State,omitempty"`
	PostalCode   *PostalCode   `xml:"PostalCode,omitempty"`
	Country      *Country      `xml:"Country,omitempty"`
	Extrinsic    []*Extrinsic  `xml:"Extrinsic,omitempty"`
}

// Address is a complete postal address with contact methods.
type Address struct {
	XMLName         xml.Name       `xml:"Address"`
	IsoCountryCode  string         `xml:"isoCountryCode,attr,omitempty"`
	AddressID       string         `xml:"addressID,attr,omitempty"`
	AddressIDDomain string         `xml:"addressIDDomain,attr,omitempty"`
	Name            *Name          `xml:"Name"`
	PostalAddress   *PostalAddress `xml:"PostalAddress,omitempty"`
	Email           *Email         `xml:"Email,omitempty"`
	Phone           *Phone         `xml:"Phone,omitempty"`
	Fax             *Fax           `xml:"Fax,omitempty"`
	URL             *URL           `xml:"URL,omitempty"`
}

// ─── Contact ──────────────────────────────────────────────────────────────────

// Contact identifies a business contact with optional address details.
type Contact struct {
	XMLName         xml.Name         `xml:"Contact"`
	Role            string           `xml:"role,attr,omitempty"`
	AddressID       string           `xml:"addressID,attr,omitempty"`
	AddressIDDomain string           `xml:"addressIDDomain,attr,omitempty"`
	Name            *Name            `xml:"Name,omitempty"`
	PostalAddress   []*PostalAddress `xml:"PostalAddress,omitempty"`
	Email           []*Email         `xml:"Email,omitempty"`
	Phone           []*Phone         `xml:"Phone,omitempty"`
	Fax             []*Fax           `xml:"Fax,omitempty"`
	URL             []*URL           `xml:"URL,omitempty"`
	IdReference     []*IdReference   `xml:"IdReference,omitempty"`
	Extrinsic       []*Extrinsic     `xml:"Extrinsic,omitempty"`
}

// PartnerContact combines a Contact with optional IdReferences.
type PartnerContact struct {
	XMLName     xml.Name       `xml:"PartnerContact"`
	Contact     *Contact       `xml:"Contact"`
	IdReference []*IdReference `xml:"IdReference,omitempty"`
}

// Issuer identifies the party that issued a document.
type Issuer struct {
	XMLName     xml.Name       `xml:"Issuer"`
	Contact     *Contact       `xml:"Contact,omitempty"`
	IdReference []*IdReference `xml:"IdReference,omitempty"`
}

// ─── Ship/Bill addresses ─────────────────────────────────────────────────────

// ShipTo is the delivery destination.
type ShipTo struct {
	XMLName              xml.Name                `xml:"ShipTo"`
	Address              *Address                `xml:"Address"`
	CarrierIdentifier    []*CarrierIdentifier    `xml:"CarrierIdentifier,omitempty"`
	TransportInformation []*TransportInformation `xml:"TransportInformation,omitempty"`
	IdReference          []*IdReference          `xml:"IdReference,omitempty"`
}

// BillTo is the billing address.
type BillTo struct {
	XMLName     xml.Name       `xml:"BillTo"`
	Address     *Address       `xml:"Address"`
	IdReference []*IdReference `xml:"IdReference,omitempty"`
}

// CarrierIdentifier identifies a shipping carrier.
type CarrierIdentifier struct {
	XMLName xml.Name `xml:"CarrierIdentifier"`
	Domain  string   `xml:"domain,attr"` // REQUIRED
	Value   string   `xml:",chardata"`
}

// TransportInformation holds optional routing and shipping instructions.
type TransportInformation struct {
	XMLName                xml.Name                `xml:"TransportInformation"`
	Route                  *Route                  `xml:"Route,omitempty"`
	ShippingContractNumber *ShippingContractNumber `xml:"ShippingContractNumber,omitempty"`
	ShippingInstructions   *ShippingInstructions   `xml:"ShippingInstructions,omitempty"`
}

// Route describes a transport route.
type Route struct {
	XMLName xml.Name `xml:"Route"`
	Content string   `xml:",innerxml"`
}

// ShippingContractNumber is a carrier contract reference.
type ShippingContractNumber struct {
	XMLName xml.Name `xml:"ShippingContractNumber"`
	Value   string   `xml:",chardata"`
}

// ShippingInstructions contains free-form shipping instructions.
type ShippingInstructions struct {
	XMLName xml.Name `xml:"ShippingInstructions"`
	Value   string   `xml:",chardata"`
}
