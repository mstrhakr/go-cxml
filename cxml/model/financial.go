package model

import "encoding/xml"

// ─── Tax ─────────────────────────────────────────────────────────────────────

// Tax represents a tax charge with optional detail breakdown.
type Tax struct {
	XMLName             xml.Name             `xml:"Tax"`
	Money               *Money               `xml:"Money"`
	TaxAdjustmentAmount *TaxAdjustmentAmount `xml:"TaxAdjustmentAmount,omitempty"`
	Description         *Description         `xml:"Description,omitempty"`
	TaxDetail           []*TaxDetail         `xml:"TaxDetail,omitempty"`
	Distribution        []*Distribution      `xml:"Distribution,omitempty"`
	Extrinsic           []*Extrinsic         `xml:"Extrinsic,omitempty"`
}

// TaxDetail holds per-line tax detail.
type TaxDetail struct {
	XMLName                           xml.Name                           `xml:"TaxDetail"`
	TaxedElement                      string                             `xml:"taxedElement,attr,omitempty"`
	Purpose                           string                             `xml:"purpose,attr,omitempty"`
	Category                          string                             `xml:"category,attr"` // REQUIRED
	PercentageRate                    string                             `xml:"percentageRate,attr,omitempty"`
	IsVatRecoverable                  string                             `xml:"isVatRecoverable,attr,omitempty"` // (yes)
	TaxPointDate                      string                             `xml:"taxPointDate,attr,omitempty"`
	PaymentDate                       string                             `xml:"paymentDate,attr,omitempty"`
	IsTriangularTransaction           string                             `xml:"isTriangularTransaction,attr,omitempty"` // (yes)
	ExemptDetail                      string                             `xml:"exemptDetail,attr,omitempty"`            // (zeroRated|exempt)
	IsWithholdingTax                  string                             `xml:"isWithholdingTax,attr,omitempty"`        // (yes)
	TaxRateType                       string                             `xml:"taxRateType,attr,omitempty"`
	BasePercentageRate                string                             `xml:"basePercentageRate,attr,omitempty"`
	IsIncludedInPrice                 string                             `xml:"isIncludedInPrice,attr,omitempty"` // (yes)
	TaxableAmount                     *TaxableAmount                     `xml:"TaxableAmount,omitempty"`
	TaxAmount                         *TaxAmount                         `xml:"TaxAmount"`
	TaxLocation                       *TaxLocation                       `xml:"TaxLocation,omitempty"`
	TaxAdjustmentAmount               *TaxAdjustmentAmount               `xml:"TaxAdjustmentAmount,omitempty"`
	Description                       *Description                       `xml:"Description,omitempty"`
	TriangularTransactionLawReference *TriangularTransactionLawReference `xml:"TriangularTransactionLawReference,omitempty"`
}

// TaxAmount wraps the tax money amount.
type TaxAmount struct {
	XMLName xml.Name `xml:"TaxAmount"`
	Money   *Money   `xml:"Money"`
}

// TaxableAmount wraps the amount that tax is applied to.
type TaxableAmount struct {
	XMLName xml.Name `xml:"TaxableAmount"`
	Money   *Money   `xml:"Money"`
}

// TaxLocation is the jurisdiction for the tax.
type TaxLocation struct {
	XMLName xml.Name `xml:"TaxLocation"`
	Lang    string   `xml:"xml:lang,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// TaxAdjustmentAmount is a tax adjustment value.
type TaxAdjustmentAmount struct {
	XMLName xml.Name `xml:"TaxAdjustmentAmount"`
	Money   *Money   `xml:"Money"`
}

// TaxAdjustmentDetail is a line-level tax adjustment detail.
type TaxAdjustmentDetail struct {
	XMLName xml.Name `xml:"TaxAdjustmentDetail"`
	Money   *Money   `xml:"Money"`
}

// TriangularTransactionLawReference is a legal reference for triangular transactions.
type TriangularTransactionLawReference struct {
	XMLName xml.Name `xml:"TriangularTransactionLawReference"`
	Lang    string   `xml:"xml:lang,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// ─── Shipping ─────────────────────────────────────────────────────────────────

// Shipping describes the shipping cost and method.
type Shipping struct {
	XMLName        xml.Name       `xml:"Shipping"`
	TrackingDomain string         `xml:"trackingDomain,attr,omitempty"`
	TrackingId     string         `xml:"trackingId,attr,omitempty"`
	Tracking       string         `xml:"tracking,attr,omitempty"`
	Money          *Money         `xml:"Money"`
	Description    *Description   `xml:"Description,omitempty"`
	Modifications  *Modifications `xml:"Modifications,omitempty"`
}

// ─── Payment ──────────────────────────────────────────────────────────────────

// Payment describes the payment instrument.
type Payment struct {
	XMLName      xml.Name      `xml:"Payment"`
	PCard        *PCard        `xml:"PCard,omitempty"`
	PaymentToken *PaymentToken `xml:"PaymentToken,omitempty"`
}

// PCard is a purchasing card (credit card) payment.
type PCard struct {
	XMLName       xml.Name       `xml:"PCard"`
	Number        string         `xml:"number,attr"`     // REQUIRED
	Expiration    string         `xml:"expiration,attr"` // REQUIRED
	Name          string         `xml:"name,attr,omitempty"`
	PostalAddress *PostalAddress `xml:"PostalAddress,omitempty"`
}

// PaymentToken is a tokenised payment instrument.
type PaymentToken struct {
	XMLName    xml.Name `xml:"PaymentToken"`
	Number     string   `xml:"number,attr"`     // REQUIRED
	Expiration string   `xml:"expiration,attr"` // REQUIRED
	Name       string   `xml:"name,attr,omitempty"`
	Token      string   `xml:"token,attr"` // REQUIRED
}

// PaymentTerm specifies payment terms including optional discount.
type PaymentTerm struct {
	XMLName           xml.Name     `xml:"PaymentTerm"`
	PayInNumberOfDays string       `xml:"payInNumberOfDays,attr"` // REQUIRED
	Discount          *Discount    `xml:"Discount,omitempty"`
	Extrinsic         []*Extrinsic `xml:"Extrinsic,omitempty"`
}

// PaymentTerms groups payment terms with a code.
type PaymentTerms struct {
	XMLName         xml.Name       `xml:"PaymentTerms"`
	PaymentTermCode string         `xml:"paymentTermCode,attr"` // REQUIRED
	PaymentTerm     []*PaymentTerm `xml:"PaymentTerm,omitempty"`
	Description     *Description   `xml:"Description,omitempty"`
	Extrinsic       []*Extrinsic   `xml:"Extrinsic,omitempty"`
}

// Discount is a price discount (percent or fixed amount).
type Discount struct {
	XMLName         xml.Name         `xml:"Discount"`
	DiscountPercent *DiscountPercent `xml:"DiscountPercent,omitempty"`
	DiscountAmount  *DiscountAmount  `xml:"DiscountAmount,omitempty"`
}

// DiscountPercent is a percentage-based discount.
type DiscountPercent struct {
	XMLName xml.Name `xml:"DiscountPercent"`
	Percent string   `xml:"percent,attr"` // REQUIRED
}

// DiscountAmount is a fixed monetary discount.
type DiscountAmount struct {
	XMLName xml.Name `xml:"DiscountAmount"`
	Money   *Money   `xml:"Money"`
}

// DiscountBasis is the base amount for a discount calculation.
type DiscountBasis struct {
	XMLName xml.Name `xml:"DiscountBasis"`
	Money   *Money   `xml:"Money"`
}

// ─── Cost allocation ──────────────────────────────────────────────────────────

// Distribution allocates cost across an accounting segment.
type Distribution struct {
	XMLName    xml.Name    `xml:"Distribution"`
	Accounting *Accounting `xml:"Accounting"`
	Charge     *Charge     `xml:"Charge"`
}

// Accounting holds cost center or GL account information.
type Accounting struct {
	XMLName           xml.Name             `xml:"Accounting"`
	Name              string               `xml:"name,attr"` // REQUIRED
	AccountingSegment []*AccountingSegment `xml:"AccountingSegment,omitempty"`
	Extrinsic         []*Extrinsic         `xml:"Extrinsic,omitempty"`
}

// AccountingSegment is a single accounting code segment.
type AccountingSegment struct {
	XMLName     xml.Name     `xml:"AccountingSegment"`
	ID          string       `xml:"id,attr"` // REQUIRED
	Name        *Name        `xml:"Name"`
	Description *Description `xml:"Description,omitempty"`
}

// Segment is a deprecated single accounting segment (older style).
type Segment struct {
	XMLName     xml.Name `xml:"Segment"`
	Type        string   `xml:"type,attr"`        // REQUIRED
	ID          string   `xml:"id,attr"`          // REQUIRED
	Description string   `xml:"description,attr"` // REQUIRED
}

// SpendDetail holds service or fee breakdown for indirect spend.
type SpendDetail struct {
	XMLName      xml.Name      `xml:"SpendDetail"`
	TravelDetail *TravelDetail `xml:"TravelDetail,omitempty"`
	FeeDetail    *FeeDetail    `xml:"FeeDetail,omitempty"`
	LaborDetail  *LaborDetail  `xml:"LaborDetail,omitempty"`
	Extrinsic    *Extrinsic    `xml:"Extrinsic,omitempty"`
}

// TravelDetail describes travel-related spend.
type TravelDetail struct {
	XMLName xml.Name `xml:"TravelDetail"`
	Content string   `xml:",innerxml"`
}

// FeeDetail describes fee-based spend.
type FeeDetail struct {
	XMLName     xml.Name    `xml:"FeeDetail"`
	IsRecurring string      `xml:"isRecurring,attr,omitempty"` // (yes)
	UnitRate    []*UnitRate `xml:"UnitRate"`
	Period      *Period     `xml:"Period,omitempty"`
}

// LaborDetail describes labor services spend.
type LaborDetail struct {
	XMLName xml.Name `xml:"LaborDetail"`
	Content string   `xml:",innerxml"`
}

// UnitRate is a price per time unit.
type UnitRate struct {
	XMLName       xml.Name       `xml:"UnitRate"`
	Money         *Money         `xml:"Money"`
	UnitOfMeasure *UnitOfMeasure `xml:"UnitOfMeasure"`
	Description   *Description   `xml:"Description,omitempty"`
}

// ─── Tolerances ───────────────────────────────────────────────────────────────

// Tolerances defines acceptable quantity, price and time variance ranges.
type Tolerances struct {
	XMLName           xml.Name           `xml:"Tolerances"`
	QuantityTolerance *QuantityTolerance `xml:"QuantityTolerance,omitempty"`
	PriceTolerance    *PriceTolerance    `xml:"PriceTolerance,omitempty"`
	TimeTolerance     *TimeTolerance     `xml:"TimeTolerance,omitempty"`
}

// QuantityTolerance sets the allowed quantity variance.
type QuantityTolerance struct {
	XMLName    xml.Name    `xml:"QuantityTolerance"`
	Percentage *Percentage `xml:"Percentage,omitempty"`
}

// PriceTolerance sets the allowed price variance.
type PriceTolerance struct {
	XMLName    xml.Name    `xml:"PriceTolerance"`
	Percentage *Percentage `xml:"Percentage,omitempty"`
	Money      *Money      `xml:"Money,omitempty"`
}

// TimeTolerance sets the allowed delivery date variance.
type TimeTolerance struct {
	XMLName xml.Name `xml:"TimeTolerance"`
	Limit   string   `xml:"limit,attr"`          // REQUIRED
	Type    string   `xml:"type,attr,omitempty"` // (minutes|hours|days|weeks)
}

// Upper is an upper bound for a tolerance.
type Upper struct {
	XMLName    xml.Name    `xml:"Upper"`
	Tolerances *Tolerances `xml:"Tolerances"`
}

// Lower is a lower bound for a tolerance.
type Lower struct {
	XMLName    xml.Name    `xml:"Lower"`
	Tolerances *Tolerances `xml:"Tolerances"`
}
