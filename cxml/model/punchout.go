package model

import "encoding/xml"

type PunchOutOrderMessage struct {
	XMLName                    xml.Name                    `xml:"PunchOutOrderMessage"`
	BuyerCookie                string                      `xml:"BuyerCookie,omitempty"`
	PunchOutOrderMessageHeader *PunchOutOrderMessageHeader `xml:"PunchOutOrderMessageHeader,omitempty"`
	ItemIn                     []ItemIn                    `xml:"ItemIn,omitempty"`
}

type PunchOutOrderMessageHeader struct {
	XMLName                xml.Name `xml:"PunchOutOrderMessageHeader"`
	PunchOutOrderMessageID string   `xml:"payloadID,attr,omitempty"`
	Operation              string   `xml:"operation,attr,omitempty"`
	Total                  *Money   `xml:"Total>Money,omitempty"`
}

type ItemIn struct {
	XMLName    xml.Name    `xml:"ItemIn"`
	Quantity   float64     `xml:"quantity,attr,omitempty"`
	LineNumber int         `xml:"lineNumber,attr,omitempty"`
	ItemID     *ItemID     `xml:"ItemID,omitempty"`
	ItemDetail *ItemDetail `xml:"ItemDetail,omitempty"`
}

// ItemID used in PunchOut messages.
type ItemID struct {
	XMLName                 xml.Name `xml:"ItemID"`
	SupplierPartID          string   `xml:"SupplierPartID,omitempty"`
	SupplierPartAuxiliaryID string   `xml:"SupplierPartAuxiliaryID,omitempty"`
}

func (p *PunchOutOrderMessage) RequestPayloadName() string {
	return "PunchOutOrderMessage"
}

func (o *OrderRequest) RequestPayloadName() string {
	return "OrderRequest"
}
