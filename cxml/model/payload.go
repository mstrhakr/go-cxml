package model

type RequestPayload interface {
	RequestPayloadName() string
}

type ResponsePayload interface {
	ResponsePayloadName() string
}

type MessagePayload interface {
	MessagePayloadName() string
}
