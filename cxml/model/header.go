package model

import "encoding/xml"

// ─── Header ───────────────────────────────────────────────────────────────────

// Header is the cXML document header containing routing credentials.
type Header struct {
	XMLName xml.Name `xml:"Header"`
	From    *From    `xml:"From"`
	To      *To      `xml:"To"`
	Sender  *Sender  `xml:"Sender"`
	Path    *Path    `xml:"Path,omitempty"`
}

// Party is a compatibility type used by existing builders/tests for Header>From and Header>To.
type Party struct {
	Credential    *Credential    `xml:"Credential,omitempty"`
	Identity      string         `xml:"Identity,omitempty"`
	Contact       *Contact       `xml:"Contact,omitempty"`
	Name          *Name          `xml:"Name,omitempty"`
	PostalAddress *PostalAddress `xml:"PostalAddress,omitempty"`
}

// From identifies the originating party.
type From struct {
	XMLName       xml.Name       `xml:"From"`
	Credential    *Credential    `xml:"Credential,omitempty"`
	Correspondent *Correspondent `xml:"Correspondent,omitempty"`
}

// PrimaryCredential returns the first credential or nil.
func (f *From) PrimaryCredential() *Credential {
	if f != nil {
		return f.Credential
	}
	return nil
}

// To identifies the recipient party.
type To struct {
	XMLName       xml.Name       `xml:"To"`
	Credential    *Credential    `xml:"Credential,omitempty"`
	Correspondent *Correspondent `xml:"Correspondent,omitempty"`
}

// PrimaryCredential returns the first credential or nil.
func (t *To) PrimaryCredential() *Credential {
	if t != nil {
		return t.Credential
	}
	return nil
}

// Sender identifies the sending system and its credential.
type Sender struct {
	XMLName       xml.Name       `xml:"Sender"`
	Credential    *Credential    `xml:"Credential,omitempty"`
	UserAgent     string         `xml:"UserAgent,omitempty"`
	Correspondent *Correspondent `xml:"Correspondent,omitempty"`
}

// PrimaryCredential returns the first credential or nil.
func (s *Sender) PrimaryCredential() *Credential {
	if s != nil {
		return s.Credential
	}
	return nil
}

// ─── Credential ───────────────────────────────────────────────────────────────

// Credential carries an identity and authentication secret.
type Credential struct {
	XMLName       xml.Name       `xml:"Credential"`
	Domain        string         `xml:"domain,attr"`            // REQUIRED
	Type          string         `xml:"type,attr,omitempty"`    // (marketplace)
	Identity      string         `xml:"Identity"`               // REQUIRED child element
	SharedSecret  string         `xml:"SharedSecret,omitempty"` // one of SharedSecret|CredentialMac
	CredentialMac *CredentialMac `xml:"CredentialMac,omitempty"`
}

// CredentialMac is a MAC-based authentication token.
type CredentialMac struct {
	XMLName        xml.Name `xml:"CredentialMac"`
	Type           string   `xml:"type,attr"`           // REQUIRED
	Algorithm      string   `xml:"algorithm,attr"`      // REQUIRED
	CreationDate   string   `xml:"creationDate,attr"`   // REQUIRED
	ExpirationDate string   `xml:"expirationDate,attr"` // REQUIRED
	Value          string   `xml:",chardata"`
}

// ─── Correspondent ────────────────────────────────────────────────────────────

// Correspondent represents an associated contact for a From/To/Sender party.
type Correspondent struct {
	XMLName           xml.Name     `xml:"Correspondent"`
	PreferredLanguage string       `xml:"preferredLanguage,attr,omitempty"`
	Contact           []*Contact   `xml:"Contact"`
	Routing           *Routing     `xml:"Routing,omitempty"`
	Extrinsic         []*Extrinsic `xml:"Extrinsic,omitempty"`
}

// Routing specifies a destination routing endpoint (e.g. Peppol).
type Routing struct {
	XMLName     xml.Name `xml:"Routing"`
	Destination string   `xml:"destination,attr"` // (peppol|fieldglass) REQUIRED
}
