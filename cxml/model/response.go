package model

import "encoding/xml"

type Response struct {
	XMLName               xml.Name               `xml:"Response"`
	Status                *Status                `xml:"Status,omitempty"`
	OrderResponse         *OrderResponse         `xml:"OrderResponse,omitempty"`
	ProfileResponse       *ProfileResponse       `xml:"ProfileResponse,omitempty"`
	PunchOutSetupResponse *PunchOutSetupResponse `xml:"PunchOutSetupResponse,omitempty"`
}

func (r *Response) PayloadType() string {
	switch {
	case r.OrderResponse != nil:
		return "OrderResponse"
	case r.ProfileResponse != nil:
		return "ProfileResponse"
	case r.PunchOutSetupResponse != nil:
		return "PunchOutSetupResponse"
	case r.Status != nil:
		return "Status"
	default:
		return ""
	}
}

type ProfileResponse struct {
	XMLName       xml.Name       `xml:"ProfileResponse"`
	EffectiveDate string         `xml:"effectiveDate,attr,omitempty"`
	LastRefresh   string         `xml:"lastRefresh,attr,omitempty"`
	Option        []*Option      `xml:"Option,omitempty"`
	Transaction   []*Transaction `xml:"Transaction,omitempty"`
}

type Option struct {
	XMLName xml.Name `xml:"Option"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Transaction struct {
	XMLName     xml.Name  `xml:"Transaction"`
	RequestName string    `xml:"requestName,attr,omitempty"`
	URL         *URL      `xml:"URL,omitempty"`
	Option      []*Option `xml:"Option,omitempty"`
}

// OrderResponse represents a cXML order response.
type OrderResponse struct {
	XMLName             xml.Name             `xml:"OrderResponse"`
	OrderResponseHeader *OrderResponseHeader `xml:"OrderResponseHeader,omitempty"`
}

type OrderResponseHeader struct {
	XMLName   xml.Name `xml:"OrderResponseHeader"`
	OrderID   string   `xml:"orderID,attr,omitempty"`
	OrderDate string   `xml:"orderDate,attr,omitempty"`
	Status    *Status  `xml:"Status,omitempty"`
}

// Status is used for both Response and top-level success/failure.
type Status struct {
	XMLName xml.Name `xml:"Status"`
	Code    string   `xml:"code,attr,omitempty"`
	Text    string   `xml:"text,attr,omitempty"`
}
