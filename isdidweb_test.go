package didweb_test

import (
	"github.com/reiver/go-didweb"

	"testing"
)

func TestIsDIDWeb(t *testing.T) {

	tests := []struct{
		DID string
		Expected bool
	}{
		{
			DID: "",
			Expected: false,
		},



		{
			DID: "apple",
			Expected: false,
		},
		{
			DID: "banana",
			Expected: false,
		},
		{
			DID: "cherry",
			Expected: false,
		},



		{
			DID: "did:once:one",
			Expected: false,
		},
		{
			DID: "did:twice:two",
			Expected: false,
		},
		{
			DID: "did:thrice:three",
			Expected: false,
		},
		{
			DID: "did:fource:four",
			Expected: false,
		},



		{
			DID: "did:plc:",
			Expected: false,
		},
		{
			DID: "did:plc:abcde12345",
			Expected: false,
		},



		{
			DID: "did:web",
			Expected: false,
		},
		{
			DID: "did:web:",
			Expected: false,
		},



		{
			DID: "did:web:once",
			Expected: true,
		},
		{
			DID: "did:web:twice",
			Expected: true,
		},
		{
			DID: "did:web:thrice",
			Expected: true,
		},
		{
			DID: "did:web:fource",
			Expected: true,
		},



		{
			DID: "did:web:example.com",
			Expected: true,
		},
		{
			DID: "diD:web:example.com",
			Expected: true,
		},
		{
			DID: "dId:web:example.com",
			Expected: true,
		},
		{
			DID: "dID:web:example.com",
			Expected: true,
		},
		{
			DID: "Did:web:example.com",
			Expected: true,
		},
		{
			DID: "DiD:web:example.com",
			Expected: true,
		},
		{
			DID: "DId:web:example.com",
			Expected: true,
		},
		{
			DID: "DID:web:example.com",
			Expected: true,
		},



		{
			DID: "did:web:example.com:users:joeblow",
			Expected: true,
		},
		{
			DID: "diD:web:example.com:users:joeblow",
			Expected: true,
		},
		{
			DID: "dId:web:example.com:users:joeblow",
			Expected: true,
		},
		{
			DID: "dID:web:example.com:users:joeblow",
			Expected: true,
		},
		{
			DID: "Did:web:example.com:users:joeblow",
			Expected: true,
		},
		{
			DID: "DiD:web:example.com:users:joeblow",
			Expected: true,
		},
		{
			DID: "DId:web:example.com:users:joeblow",
			Expected: true,
		},
		{
			DID: "DID:web:example.com:users:joeblow",
			Expected: true,
		},



		{
			DID: "did:weB:example.com",
			Expected: false,
		},
		{
			DID: "did:wEb:example.com",
			Expected: false,
		},
		{
			DID: "did:wEB:example.com",
			Expected: false,
		},
		{
			DID: "did:Web:example.com",
			Expected: false,
		},
		{
			DID: "did:WeB:example.com",
			Expected: false,
		},
		{
			DID: "did:WEb:example.com",
			Expected: false,
		},
		{
			DID: "did:WEB:example.com",
			Expected: false,
		},



		{
			DID: "did:weB:example.com:users:joeblow",
			Expected: false,
		},
		{
			DID: "did:wEb:example.com:users:joeblow",
			Expected: false,
		},
		{
			DID: "did:wEB:example.com:users:joeblow",
			Expected: false,
		},
		{
			DID: "did:Web:example.com:users:joeblow",
			Expected: false,
		},
		{
			DID: "did:WeB:example.com:users:joeblow",
			Expected: false,
		},
		{
			DID: "did:WEb:example.com:users:joeblow",
			Expected: false,
		},
		{
			DID: "did:WEB:example.com:users:joeblow",
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := didweb.IsDIDWeb(test.DID)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, whether the actual DID was a DID-Web or not was not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("DID: %q", test.DID)
			continue
		}
	}
}
