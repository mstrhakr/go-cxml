package model

import "encoding/xml"

type PunchOutOrderMessage struct {
	XMLName                    xml.Name                    `xml:"PunchOutOrderMessage"`
	BuyerCookie                string                      `xml:"BuyerCookie,omitempty"`
	PunchOutOrderMessageHeader *PunchOutOrderMessageHeader `xml:"PunchOutOrderMessageHeader,omitempty"`
	ItemIn                     []ItemIn                    `xml:"ItemIn,omitempty"`
}

type PunchOutOrderMessageHeader struct {
	XMLName                xml.Name           `xml:"PunchOutOrderMessageHeader"`
	PunchOutOrderMessageID string             `xml:"payloadID,attr,omitempty"`
	Operation              string             `xml:"operation,attr,omitempty"`
	OperationAllowed       string             `xml:"operationAllowed,attr,omitempty"`
	Total                  *Money             `xml:"Total>Money,omitempty"`
	ShipTo                 *ShipTo            `xml:"ShipTo,omitempty"`
	Shipping               *Shipping          `xml:"Shipping,omitempty"`
	Tax                    *Tax               `xml:"Tax,omitempty"`
	SupplierOrderInfo      *SupplierOrderInfo `xml:"SupplierOrderInfo,omitempty"`
}

type ItemIn struct {
	XMLName            xml.Name        `xml:"ItemIn"`
	Quantity           float64         `xml:"quantity,attr,omitempty"`
	LineNumber         int             `xml:"lineNumber,attr,omitempty"`
	ParentLineNumber   string          `xml:"parentLineNumber,attr,omitempty"`
	ItemType           string          `xml:"itemType,attr,omitempty"`
	CompositeItemType  string          `xml:"compositeItemType,attr,omitempty"`
	ItemClassification string          `xml:"itemClassification,attr,omitempty"`
	ItemCategory       string          `xml:"itemCategory,attr,omitempty"`
	ItemID             *ItemID         `xml:"ItemID,omitempty"`
	ItemDetail         *ItemDetail     `xml:"ItemDetail,omitempty"`
	ShipTo             *ShipTo         `xml:"ShipTo,omitempty"`
	Shipping           *Shipping       `xml:"Shipping,omitempty"`
	Tax                *Tax            `xml:"Tax,omitempty"`
	SpendDetail        *SpendDetail    `xml:"SpendDetail,omitempty"`
	Distribution       []*Distribution `xml:"Distribution,omitempty"`
	Contact            []*Contact      `xml:"Contact,omitempty"`
}

// PunchOutSetupRequest opens a PunchOut session.
type PunchOutSetupRequest struct {
	XMLName         xml.Name         `xml:"PunchOutSetupRequest"`
	Operation       string           `xml:"operation,attr,omitempty"`
	BuyerCookie     string           `xml:"BuyerCookie,omitempty"`
	Extrinsic       []*Extrinsic     `xml:"Extrinsic,omitempty"`
	BrowserFormPost *BrowserFormPost `xml:"BrowserFormPost,omitempty"`
	Contact         []*Contact       `xml:"Contact,omitempty"`
	SupplierSetup   *SupplierSetup   `xml:"SupplierSetup,omitempty"`
	ShipTo          *ShipTo          `xml:"ShipTo,omitempty"`
	SelectedItem    *SelectedItem    `xml:"SelectedItem,omitempty"`
	ItemOut         []*ItemOut       `xml:"ItemOut,omitempty"`
}

type BrowserFormPost struct {
	XMLName xml.Name `xml:"BrowserFormPost"`
	URL     *URL     `xml:"URL,omitempty"`
}

type SupplierSetup struct {
	XMLName xml.Name `xml:"SupplierSetup"`
	URL     *URL     `xml:"URL,omitempty"`
}

type SelectedItem struct {
	XMLName    xml.Name    `xml:"SelectedItem"`
	ItemID     *ItemID     `xml:"ItemID,omitempty"`
	ItemDetail *ItemDetail `xml:"ItemDetail,omitempty"`
}

// PunchOutSetupResponse returns the PunchOut start page.
type PunchOutSetupResponse struct {
	XMLName   xml.Name   `xml:"PunchOutSetupResponse"`
	StartPage *StartPage `xml:"StartPage,omitempty"`
}

type StartPage struct {
	XMLName xml.Name `xml:"StartPage"`
	URL     *URL     `xml:"URL,omitempty"`
}

func (p *PunchOutOrderMessage) RequestPayloadName() string {
	return "PunchOutOrderMessage"
}

func (o *OrderRequest) RequestPayloadName() string {
	return "OrderRequest"
}
