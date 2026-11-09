package didweb_test

import (
	"github.com/reiver/go-didweb"

	"testing"
)

func TestSplit(t *testing.T) {

	tests := []struct{
		Path string
		ExpectedHost string
		ExpectedPath string
		ExpectedOK   bool
	}{
		{
			Path: "",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},



		{
			Path: "d",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},
		{
			Path: "di",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},
		{
			Path: "did",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},
		{
			Path: "did:",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},
		{
			Path: "did:w",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},
		{
			Path: "did:we",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},
		{
			Path: "did:web",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},
		{
			Path: "did:web:",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},



		{
			Path: "did:web:e",
			ExpectedHost: "e",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:ex",
			ExpectedHost: "ex",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:ex",
			ExpectedHost: "ex",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:exa",
			ExpectedHost: "exa",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:exam",
			ExpectedHost: "exam",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:examp",
			ExpectedHost: "examp",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:exampl",
			ExpectedHost: "exampl",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:example",
			ExpectedHost: "example",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:example.",
			ExpectedHost: "example.",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:example.c",
			ExpectedHost: "example.c",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:example.co",
			ExpectedHost: "example.co",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:example.com",
			ExpectedHost: "example.com",
			ExpectedPath: "",
			ExpectedOK: true,
		},
		{
			Path: "did:web:example.com:",
			ExpectedHost: "",
			ExpectedPath: "",
			ExpectedOK: false,
		},
		{
			Path: "did:web:example.com:~",
			ExpectedHost: "example.com",
			ExpectedPath: "~",
			ExpectedOK: true,
		},
		{
			Path: "did:web:example.com:~joeblow",
			ExpectedHost: "example.com",
			ExpectedPath: "~joeblow",
			ExpectedOK: true,
		},
		{
			Path: "did:web:example.com:~joeblow:over-9000.gif",
			ExpectedHost: "example.com",
			ExpectedPath: "~joeblow:over-9000.gif",
			ExpectedOK: true,
		},
	}

	for testNumber, test := range tests {

		actualHost, actualPath, actualOK := didweb.Split(test.Path)

		{
			actual := actualOK
			expected := test.ExpectedOK

			if expected != actual {
				t.Errorf("For test #%d, the actual 'ok' for splitting the DID-Web (did:web) was not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("PATH: %q", test.Path )
				t.Logf("ACTUAL-HOST: %q", actualHost)
				t.Logf("ACTUAL-PATH: %q", actualPath)
				continue
			}
		}

		{
			actual := actualHost
			expected := test.ExpectedHost

			if expected != actual {
				t.Errorf("For test #%d, the actual DID-Web (did:web) 'host' was not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("PATH: %q", test.Path )
				continue
			}
		}

		{
			actual := actualPath
			expected := test.ExpectedPath

			if expected != actual {
				t.Errorf("For test #%d, the actual DID-Web (did:web) 'path' was not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("PATH: %q", test.Path )
				continue
			}
		}
	}
}
