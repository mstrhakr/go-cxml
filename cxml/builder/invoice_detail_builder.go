package builder

import "github.com/Depth8064/go-cxml/cxml/model"

// InvoiceDetailBuilder builds an outbound InvoiceDetailRequest cXML document.
type InvoiceDetailBuilder struct {
	builder *Builder
}

// NewInvoiceDetailBuilder returns a new InvoiceDetailBuilder.
func NewInvoiceDetailBuilder() *InvoiceDetailBuilder {
	return &InvoiceDetailBuilder{builder: New()}
}

func (b *InvoiceDetailBuilder) PayloadID(id string) *InvoiceDetailBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *InvoiceDetailBuilder) Timestamp(ts string) *InvoiceDetailBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *InvoiceDetailBuilder) Version(version string) *InvoiceDetailBuilder {
	b.builder.Version(version)
	return b
}

func (b *InvoiceDetailBuilder) From(party *model.Party) *InvoiceDetailBuilder {
	b.builder.From(party)
	return b
}

func (b *InvoiceDetailBuilder) To(party *model.Party) *InvoiceDetailBuilder {
	b.builder.To(party)
	return b
}

func (b *InvoiceDetailBuilder) Sender(sender *model.Sender) *InvoiceDetailBuilder {
	b.builder.Sender(sender)
	return b
}

// Request sets the InvoiceDetailRequest payload.
func (b *InvoiceDetailBuilder) Request(invoice *model.InvoiceDetailRequest) *InvoiceDetailBuilder {
	b.builder.Request(&model.Request{InvoiceDetailRequest: invoice})
	return b
}

func (b *InvoiceDetailBuilder) Build() *model.CXML {
	return b.builder.Build()
}
