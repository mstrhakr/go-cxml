package builder

import "github.com/Depth8064/go-cxml/cxml/model"

// ShipNoticeBuilder builds an outbound ShipNoticeRequest cXML document.
type ShipNoticeBuilder struct {
	builder *Builder
}

// NewShipNoticeBuilder returns a new ShipNoticeBuilder.
func NewShipNoticeBuilder() *ShipNoticeBuilder {
	return &ShipNoticeBuilder{builder: New()}
}

func (b *ShipNoticeBuilder) PayloadID(id string) *ShipNoticeBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *ShipNoticeBuilder) Timestamp(ts string) *ShipNoticeBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *ShipNoticeBuilder) Version(version string) *ShipNoticeBuilder {
	b.builder.Version(version)
	return b
}

func (b *ShipNoticeBuilder) From(party *model.Party) *ShipNoticeBuilder {
	b.builder.From(party)
	return b
}

func (b *ShipNoticeBuilder) To(party *model.Party) *ShipNoticeBuilder {
	b.builder.To(party)
	return b
}

func (b *ShipNoticeBuilder) Sender(sender *model.Sender) *ShipNoticeBuilder {
	b.builder.Sender(sender)
	return b
}

// Request sets the ShipNoticeRequest payload.
func (b *ShipNoticeBuilder) Request(shipNotice *model.ShipNoticeRequest) *ShipNoticeBuilder {
	b.builder.Request(&model.Request{ShipNoticeRequest: shipNotice})
	return b
}

func (b *ShipNoticeBuilder) Build() *model.CXML {
	return b.builder.Build()
}
