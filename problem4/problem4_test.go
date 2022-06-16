package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	var testcases = []struct {
		testcase string
		input    string
		output   bool
	}{
		{
			testcase: "Case 1",
			input:    "civic",
			output:   true,
		},
		{
			testcase: "Case 2",
			input:    "katak",
			output:   true,
		},
		{
			testcase: "Case 3",
			input:    "kasur rusak",
			output:   true,
		},
		{
			testcase: "Case 4",
			input:    "kupu-kupa",
			output:   false,
		},
		{
			testcase: "Case 5",
			input:    "lion",
			output:   false,
		},
	}
	for _, value := range testcases {
		result := Palindrome(value.input)
		t.Run(value.testcase, func(t *testing.T) {
			assert.Equal(t, value.output, result)
		})
	}
}
