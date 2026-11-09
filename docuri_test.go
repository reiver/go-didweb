package didweb_test

import (
	"github.com/reiver/go-didweb"

	"testing"
)

func TestDocURI(t *testing.T) {

	tests := []struct{
		DID string
		Expected string
	}{
		{
			DID: "",
			Expected: "",
		},



		{
			DID: "apple",
			Expected: "",
		},
		{
			DID: "banana",
			Expected: "",
		},
		{
			DID: "cherry",
			Expected: "",
		},



		{
			DID: "did:once:one",
			Expected: "",
		},
		{
			DID: "did:twice:two",
			Expected: "",
		},
		{
			DID: "did:thrice:three",
			Expected: "",
		},
		{
			DID: "did:fource:four",
			Expected: "",
		},



		{
			DID: "did:plc:",
			Expected: "",
		},
		{
			DID: "did:plc:abcde12345",
			Expected: "",
		},



		{
			DID: "did:web:once",
			Expected: "https://once/.well-known/did.json",
		},
		{
			DID: "did:web:twice",
			Expected: "https://twice/.well-known/did.json",
		},
		{
			DID: "did:web:thrice",
			Expected: "https://thrice/.well-known/did.json",
		},
		{
			DID: "did:web:fource",
			Expected: "https://fource/.well-known/did.json",
		},



		{
			DID: "did:web:example.com",
			Expected: "https://example.com/.well-known/did.json",
		},
		{
			DID: "diD:web:example.com",
			Expected: "https://example.com/.well-known/did.json",
		},
		{
			DID: "dId:web:example.com",
			Expected: "https://example.com/.well-known/did.json",
		},
		{
			DID: "dID:web:example.com",
			Expected: "https://example.com/.well-known/did.json",
		},
		{
			DID: "Did:web:example.com",
			Expected: "https://example.com/.well-known/did.json",
		},
		{
			DID: "DiD:web:example.com",
			Expected: "https://example.com/.well-known/did.json",
		},
		{
			DID: "DId:web:example.com",
			Expected: "https://example.com/.well-known/did.json",
		},
		{
			DID: "DID:web:example.com",
			Expected: "https://example.com/.well-known/did.json",
		},



		{
			DID: "did:web:example.com:users:joeblow",
			Expected: "https://example.com/users/joeblow/did.json",
		},
		{
			DID: "diD:web:example.com:users:joeblow",
			Expected: "https://example.com/users/joeblow/did.json",
		},
		{
			DID: "dId:web:example.com:users:joeblow",
			Expected: "https://example.com/users/joeblow/did.json",
		},
		{
			DID: "dID:web:example.com:users:joeblow",
			Expected: "https://example.com/users/joeblow/did.json",
		},
		{
			DID: "Did:web:example.com:users:joeblow",
			Expected: "https://example.com/users/joeblow/did.json",
		},
		{
			DID: "DiD:web:example.com:users:joeblow",
			Expected: "https://example.com/users/joeblow/did.json",
		},
		{
			DID: "DId:web:example.com:users:joeblow",
			Expected: "https://example.com/users/joeblow/did.json",
		},
		{
			DID: "DID:web:example.com:users:joeblow",
			Expected: "https://example.com/users/joeblow/did.json",
		},



		{
			DID: "did:weB:example.com",
			Expected: "",
		},
		{
			DID: "did:wEb:example.com",
			Expected: "",
		},
		{
			DID: "did:wEB:example.com",
			Expected: "",
		},
		{
			DID: "did:Web:example.com",
			Expected: "",
		},
		{
			DID: "did:WeB:example.com",
			Expected: "",
		},
		{
			DID: "did:WEb:example.com",
			Expected: "",
		},
		{
			DID: "did:WEB:example.com",
			Expected: "",
		},



		{
			DID: "did:weB:example.com:users:joeblow",
			Expected: "",
		},
		{
			DID: "did:wEb:example.com:users:joeblow",
			Expected: "",
		},
		{
			DID: "did:wEB:example.com:users:joeblow",
			Expected: "",
		},
		{
			DID: "did:Web:example.com:users:joeblow",
			Expected: "",
		},
		{
			DID: "did:WeB:example.com:users:joeblow",
			Expected: "",
		},
		{
			DID: "did:WEb:example.com:users:joeblow",
			Expected: "",
		},
		{
			DID: "did:WEB:example.com:users:joeblow",
			Expected: "",
		},
	}

	for testNumber, test := range tests {

		actual := didweb.DocURI(test.DID)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual DID-Web (did:web) document URI was not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
