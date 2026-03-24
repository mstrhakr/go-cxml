package processor

import (
	"github.com/mstrhakr/go-cxml/cxml/handler"
	"github.com/mstrhakr/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type stubOrderHandler struct{}

func (h *stubOrderHandler) Handle(req *model.CXML) (*model.CXML, error) {
	resp := &model.CXML{
		PayloadID: req.PayloadID,
		Response:  &model.Response{Status: &model.Status{Code: "200", Text: "OK"}},
	}
	return resp, nil
}

func (h *stubOrderHandler) Name() string { return "OrderRequest" }

func TestProcessor_Process_OrderRequest(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&stubOrderHandler{})

	p := NewProcessor(reg)

	req := &model.CXML{PayloadID: "x1", Request: &model.Request{Payload: &model.OrderRequest{OrderRequestHeader: &model.OrderRequestHeader{OrderID: "PO"}}}}
	resp, err := p.Process(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "x1", resp.PayloadID)
	assert.Equal(t, "200", resp.Response.Status.Code)
}

func TestProcessor_Process_NoHandler(t *testing.T) {
	p := NewProcessor(handler.NewRegistry())
	_, err := p.Process(&model.CXML{Request: &model.Request{Payload: &model.OrderRequest{}}})
	assert.Error(t, err)
}
