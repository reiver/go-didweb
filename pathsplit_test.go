package didweb_test

import (
	"github.com/reiver/go-didweb"

	"reflect"

	"testing"
)

func TestPathSplit(t *testing.T) {

	tests := []struct{
		Path string
		ExpectedParts []string
		ExpectedOK bool
	}{
		{
			Path: "",
			ExpectedParts: nil,
			ExpectedOK: true,
		},



		{
			Path: "%",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%0",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%1",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%2",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%3",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%4",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%5",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%6",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%7",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%8",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%9",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%A",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%a",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%B",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%b",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%C",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%c",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%D",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%d",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%E",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%e",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%F",
			ExpectedParts: nil,
			ExpectedOK: false,
		},
		{
			Path: "%f",
			ExpectedParts: nil,
			ExpectedOK: false,
		},



		{
			Path: "%7E",
			ExpectedParts: []string{"~"},
			ExpectedOK: true,
		},
		{
			Path: "%7e",
			ExpectedParts: []string{"~"},
			ExpectedOK: true,
		},
		{
			Path: "~",
			ExpectedParts: []string{"~"},
			ExpectedOK: true,
		},



		{
			Path: "%7Ejoeblow",
			ExpectedParts: []string{"~joeblow"},
			ExpectedOK: true,
		},
		{
			Path: "%7ejoeblow",
			ExpectedParts: []string{"~joeblow"},
			ExpectedOK: true,
		},
		{
			Path: "~joeblow",
			ExpectedParts: []string{"~joeblow"},
			ExpectedOK: true,
		},



		{
			Path: "%7Ejoeblow:over-9000.gif",
			ExpectedParts: []string{"~joeblow", "over-9000.gif"},
			ExpectedOK: true,
		},
		{
			Path: "%7ejoeblow:over-9000.gif",
			ExpectedParts: []string{"~joeblow", "over-9000.gif"},
			ExpectedOK: true,
		},
		{
			Path: "~joeblow:over-9000.gif",
			ExpectedParts: []string{"~joeblow", "over-9000.gif"},
			ExpectedOK: true,
		},



		{
			Path:                   "apple",
			ExpectedParts: []string{"apple"},
			ExpectedOK: true,
		},
		{
			Path:                   "banana",
			ExpectedParts: []string{"banana"},
			ExpectedOK: true,
		},
		{
			Path:                   "cherry",
			ExpectedParts: []string{"cherry"},
			ExpectedOK: true,
		},



		{
			Path:                   "Apple",
			ExpectedParts: []string{"Apple"},
			ExpectedOK: true,
		},
		{
			Path:                   "Banana",
			ExpectedParts: []string{"Banana"},
			ExpectedOK: true,
		},
		{
			Path:                   "Cherry",
			ExpectedParts: []string{"Cherry"},
			ExpectedOK: true,
		},



		{
			Path:                   "APPLE",
			ExpectedParts: []string{"APPLE"},
			ExpectedOK: true,
		},
		{
			Path:                   "BANANA",
			ExpectedParts: []string{"BANANA"},
			ExpectedOK: true,
		},
		{
			Path:                   "CHERRY",
			ExpectedParts: []string{"CHERRY"},
			ExpectedOK: true,
		},



		{
			Path:                   "once",
			ExpectedParts: []string{"once"},
			ExpectedOK: true,
		},
		{
			Path:                   "once:TWICE",
			ExpectedParts: []string{"once", "TWICE"},
			ExpectedOK: true,
		},
		{
			Path:                   "once:TWICE:Thrice",
			ExpectedParts: []string{"once", "TWICE", "Thrice"},
			ExpectedOK: true,
		},
		{
			Path:                   "once:TWICE:Thrice:fOuRcE",
			ExpectedParts: []string{"once", "TWICE", "Thrice", "fOuRcE"},
			ExpectedOK: true,
		},



		{
			Path:                   "users:joeblow",
			ExpectedParts: []string{"users", "joeblow"},
			ExpectedOK: true,
		},

	}

	for testNumber, test := range tests {

		actualParts, actualOK := didweb.PathSplit(test.Path)

		{
			actual := actualOK
			expected := test.ExpectedOK

			if expected != actual {
				t.Errorf("For test #%d, the actual 'ok' for splitting the DID-Web (did:web) path was not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("PATH: %q", test.Path )
				continue
			}
		}

		{
			actual := actualParts
			expected := test.ExpectedParts

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual DID-Web (did:web) document URI was not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
