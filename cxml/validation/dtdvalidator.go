package validation

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// DTDValidator validates cXML documents against DTD references.
type DTDValidator struct{}

func NewDTDValidator() *DTDValidator {
	return &DTDValidator{}
}

func (v *DTDValidator) Validate(xml []byte) error {
	if len(xml) == 0 {
		return errors.New("validation: empty document")
	}

	normalized := strings.ToLower(string(xml))

	if !strings.Contains(normalized, "<cxml") {
		return errors.New("validation: document does not contain cXML root")
	}

	if !strings.Contains(normalized, "<!doctype cxml") {
		return errors.New("validation: missing cXML doctype")
	}

	if dtddir := os.Getenv("CXML_DTD_DIR"); dtddir != "" {
		candidate := filepath.Join(dtddir, "1.2.069", "cXML.dtd")
		if _, err := os.Stat(candidate); err != nil {
			return errors.New("validation: local DTD file not found: " + candidate)
		}
	}

	return nil
}
