package model

import "encoding/xml"

// ─── Terms of delivery ────────────────────────────────────────────────────────

// TermsOfDelivery specifies the delivery terms (Incoterms).
type TermsOfDelivery struct {
	XMLName               xml.Name               `xml:"TermsOfDelivery"`
	TermsOfDeliveryCode   *TermsOfDeliveryCode   `xml:"TermsOfDeliveryCode"`
	ShippingPaymentMethod *ShippingPaymentMethod `xml:"ShippingPaymentMethod"`
	TransportTerms        *TransportTerms        `xml:"TransportTerms,omitempty"`
	Address               *Address               `xml:"Address,omitempty"`
	Comments              []*Comments            `xml:"Comments,omitempty"`
}

// TermsOfDeliveryCode is the Incoterms code (e.g. "EXW", "DDP").
type TermsOfDeliveryCode struct {
	XMLName xml.Name `xml:"TermsOfDeliveryCode"`
	Value   string   `xml:"value,attr,omitempty"`
	Text    string   `xml:",chardata"`
}

// ShippingPaymentMethod specifies who pays for shipping.
type ShippingPaymentMethod struct {
	XMLName xml.Name `xml:"ShippingPaymentMethod"`
	Value   string   `xml:"value,attr,omitempty"` // e.g. "prepaid", "collect"
	Text    string   `xml:",chardata"`
}

// TransportTerms holds additional transport arrangement details.
type TransportTerms struct {
	XMLName xml.Name `xml:"TransportTerms"`
	Value   string   `xml:",chardata"`
}

// ─── Terms of transport ───────────────────────────────────────────────────────

// TermsOfTransport describes equipment and sealing for transport.
type TermsOfTransport struct {
	XMLName                     xml.Name                     `xml:"TermsOfTransport"`
	SealID                      *SealID                      `xml:"SealID,omitempty"`
	SealingPartyCode            *SealingPartyCode            `xml:"SealingPartyCode,omitempty"`
	EquipmentIdentificationCode *EquipmentIdentificationCode `xml:"EquipmentIdentificationCode,omitempty"`
	TransportTerms              *TransportTerms              `xml:"TransportTerms,omitempty"`
	Dimension                   []*Dimension                 `xml:"Dimension,omitempty"`
	Extrinsic                   []*Extrinsic                 `xml:"Extrinsic,omitempty"`
}

// SealID is a container seal identifier.
type SealID struct {
	XMLName xml.Name `xml:"SealID"`
	Value   string   `xml:",chardata"`
}

// SealingPartyCode identifies the party responsible for sealing.
type SealingPartyCode struct {
	XMLName xml.Name `xml:"SealingPartyCode"`
	Value   string   `xml:",chardata"`
}

// EquipmentIdentificationCode identifies the transport equipment.
type EquipmentIdentificationCode struct {
	XMLName xml.Name `xml:"EquipmentIdentificationCode"`
	Value   string   `xml:",chardata"`
}

// ─── Dimensions ───────────────────────────────────────────────────────────────

// Dimension holds a measurement value with a unit of measure.
type Dimension struct {
	XMLName       xml.Name       `xml:"Dimension"`
	Type          string         `xml:"type,attr,omitempty"` // e.g. "weight", "volume", "length"
	Quantity      string         `xml:"quantity,attr,omitempty"`
	UnitOfMeasure *UnitOfMeasure `xml:"UnitOfMeasure,omitempty"`
}

// ─── Delivery timing ──────────────────────────────────────────────────────────

// DeliveryPeriod wraps a Period to specify a delivery window.
type DeliveryPeriod struct {
	XMLName xml.Name `xml:"DeliveryPeriod"`
	Period  *Period  `xml:"Period"`
}

// ─── Control keys ─────────────────────────────────────────────────────────────

// ControlKeys specifies operational controls for order processing.
type ControlKeys struct {
	XMLName            xml.Name            `xml:"ControlKeys"`
	OCInstruction      *OCInstruction      `xml:"OCInstruction,omitempty"`
	ASNInstruction     *ASNInstruction     `xml:"ASNInstruction,omitempty"`
	InvoiceInstruction *InvoiceInstruction `xml:"InvoiceInstruction,omitempty"`
	SESInstruction     *SESInstruction     `xml:"SESInstruction,omitempty"`
}

// OCInstruction controls whether order changes are allowed.
type OCInstruction struct {
	XMLName xml.Name `xml:"OCInstruction"`
	Value   string   `xml:"value,attr"` // (allowed|notAllowed|requiredBeforeASN) REQUIRED
	Lower   *Lower   `xml:"Lower,omitempty"`
	Upper   *Upper   `xml:"Upper,omitempty"`
}

// ASNInstruction controls advance ship notice requirements.
type ASNInstruction struct {
	XMLName xml.Name `xml:"ASNInstruction"`
	Value   string   `xml:"value,attr"` // (required|notRequired|optional) REQUIRED
}

// InvoiceInstruction controls invoice handling.
type InvoiceInstruction struct {
	XMLName          xml.Name `xml:"InvoiceInstruction"`
	Value            string   `xml:"value,attr"`                      // (invoiceRequired|invoiceNotRequired|invoiceForbidden|evaluated) REQUIRED
	VerificationType string   `xml:"verificationType,attr,omitempty"` // (2way|3way|4way)
}

// SESInstruction controls service entry sheet requirements.
type SESInstruction struct {
	XMLName xml.Name `xml:"SESInstruction"`
	Value   string   `xml:"value,attr"` // (required|notRequired|optional) REQUIRED
}

// ─── Packaging ────────────────────────────────────────────────────────────────

// Packaging describes physical packaging for shipment.
type Packaging struct {
	XMLName                              xml.Name                              `xml:"Packaging"`
	PackagingCode                        []*PackagingCode                      `xml:"PackagingCode,omitempty"`
	Dimension                            []*Dimension                          `xml:"Dimension,omitempty"`
	Description                          *Description                          `xml:"Description,omitempty"`
	PackagingLevelCode                   *PackagingLevelCode                   `xml:"PackagingLevelCode,omitempty"`
	ShippingContainerSerialCode          *ShippingContainerSerialCode          `xml:"ShippingContainerSerialCode,omitempty"`
	ShippingContainerSerialCodeReference *ShippingContainerSerialCodeReference `xml:"ShippingContainerSerialCodeReference,omitempty"`
	AssetInfo                            []*AssetInfo                          `xml:"AssetInfo,omitempty"`
}

// PackagingCode is a packaging type code.
type PackagingCode struct {
	XMLName xml.Name `xml:"PackagingCode"`
	Value   string   `xml:",chardata"`
}

// PackagingLevelCode is the level in a multi-level packaging hierarchy.
type PackagingLevelCode struct {
	XMLName xml.Name `xml:"PackagingLevelCode"`
	Value   string   `xml:",chardata"`
}

// ShippingContainerSerialCode is an SSCC barcode value.
type ShippingContainerSerialCode struct {
	XMLName xml.Name `xml:"ShippingContainerSerialCode"`
	Value   string   `xml:",chardata"`
}

// ShippingContainerSerialCodeReference refers to another container.
type ShippingContainerSerialCodeReference struct {
	XMLName xml.Name `xml:"ShippingContainerSerialCodeReference"`
	Value   string   `xml:",chardata"`
}
