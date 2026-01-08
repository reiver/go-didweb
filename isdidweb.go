package didweb

import (
	"strings"
)

// IsDIDWeb returns whether the passed value is a DID-Web or not.
func IsDIDWeb(did string) bool {
	const didPart string = "did:"
	const webPart string = "web:"

	// We do a "+ 1" because the domain-name part of a DID-Web cannot be empty.
	if len(did) < len(didPart) + len(webPart) + 1 {
		return false
	}

	// URL / URI / IRI schemes are case-insensitive.
	//
	// I.e., these DID-Web are the same:
	//
	//	• did:web:example.com
	//	• diD:web:example.com
	//	• dId:web:example.com
	//	• dID:web:example.com
	//	• Did:web:example.com
	//	• DiD:web:example.com
	//	• DId:web:example.com
	//	• DID:web:example.com
	if !strings.HasPrefix(strings.ToLower(did[:len(didPart)]), didPart) {
		return false
	}

	// DID methods (including the "web" DID method) must be lower-case.
	if !strings.HasPrefix(did[len(didPart):len(didPart)+len(webPart)], webPart) {
		return false
	}

	return true
}
