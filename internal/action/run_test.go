package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type OutputStringTestCase struct {
	Interface interface{}
	Expected  string
	Message   string
}

func TestOutputString(t *testing.T) {
	testCases := []OutputStringTestCase{
		{
			Message:   "Terraform string type",
			Interface: "foo",
			Expected:  "foo",
		},
		{
			Message:   "Terraform int type",
			Interface: int(1),
			Expected:  "1",
		},
		{
			Message:   "Terraform float type",
			Interface: float64(3.14),
			Expected:  "3.14",
		},
		{
			Message:   "Terraform boolean type",
			Interface: true,
			Expected:  "true",
		},
		{
			Message: "Terraform map or object type",
			Interface: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
				},
			},
			Expected: `{"foo":{"bar":"baz"}}`,
		},
		{
			Message:   "Terraform list or set type",
			Interface: []interface{}{"foo", "bar", "baz"},
			Expected:  `["foo","bar","baz"]`,
		},
		{
			Message:   "Terraform null type",
			Interface: nil,
			Expected:  "",
		},
		{
			Message: "Multiline string",
			Interface: `
				multi
				line
				string`,
			Expected: `
				multi
				line
				string`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Message, func(t *testing.T) {
			actual, err := OutputString(tc.Interface)
			require.NoError(t, err)

			assert.Equal(t, tc.Expected, actual)
		})
	}
}
