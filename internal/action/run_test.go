package action

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type OutputStringTestCase struct {
	Interface interface{}
	Expected  string
	Message   string
	Error     error
}

func TestOutputString(t *testing.T) {
	testCases := []OutputStringTestCase{
		{
			Message:   "Test string output",
			Interface: "foo",
			Expected:  "foo",
		},
		{
			Message: "Test JSON string output",
			Interface: map[string]interface{}{
				"foo": "bar",
				"baz": "woz",
			},
			Expected: `{"baz":"woz","foo":"bar"}`,
		},
		{
			Message:   "Test slice string output",
			Interface: []string{"foo", "bar", "baz"},
			Expected:  `["foo","bar","baz"]`,
		},
		{
			Message:   "Test array string output",
			Interface: [2]int{1, 2},
			Expected:  `[1,2]`,
		},
		{
			Message:   "Test int",
			Interface: int(1),
			Expected:  "1",
		},
		{
			Message:   "Test int32",
			Interface: int32(1),
			Expected:  "1",
		},
		{
			Message:   "Test int64",
			Interface: int64(1),
			Expected:  "1",
		},
		{
			Message:   "Test uint",
			Interface: uint(1),
			Expected:  "1",
		},
		{
			Message:   "Test uint32",
			Interface: uint32(1),
			Expected:  "1",
		},
		{
			Message:   "Test uint64",
			Interface: uint64(1),
			Expected:  "1",
		},
		{
			Message:   "Test float32",
			Interface: float32(3.14),
			Expected:  "3.14",
		},
		{
			Message:   "Test float64",
			Interface: float64(3.14),
			Expected:  "3.14",
		},
		{
			Message:   "Test nil",
			Interface: nil,
			Expected:  "",
		},
		{
			Message:   "Test bool",
			Interface: true,
			Expected:  "true",
		},
		{
			Message: "Test struct",
			Interface: struct {
				Foo string
				Bar int `json:"bar"`
			}{
				Foo: "foo",
				Bar: 5,
			},
			Expected: `{"Foo":"foo","bar":5}`,
		},
		{
			Message: "Test multiline",
			Interface: `multi
		line
		string`,
			Expected: `multi
		line
		string`,
		},
		{
			Message:   "Test chan",
			Interface: make(chan int),
			Expected:  "",
		},
		{
			Message:   "Test func",
			Interface: func(string) string { return "" },
			Expected:  "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Message, func(t *testing.T) {
			actual, err := OutputString(tc.Interface)
			if tc.Error != nil {
				fmt.Println(err)
				assert.ErrorIs(t, err, tc.Error)
			} else {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
