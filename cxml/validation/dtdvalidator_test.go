package validation

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDTDValidator_Valid(t *testing.T) {
    xml := []byte(`<?xml version="1.0"?>
<!DOCTYPE cXML SYSTEM "http://xml.cxml.org/schemas/cXML/1.2.014/cXML.dtd">
<cXML payloadID="abc"></cXML>`) 
    v := NewDTDValidator()
    err := v.Validate(xml)
    assert.NoError(t, err)
}

func TestDTDValidator_MissingDoctype(t *testing.T) {
    xml := []byte(`<?xml version="1.0"?><cXML payloadID="abc"></cXML>`)
    v := NewDTDValidator()
    err := v.Validate(xml)
    assert.Error(t, err)
}
