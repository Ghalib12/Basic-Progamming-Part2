package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeNumber(t *testing.T) {
	var testcases = []struct {
		testcase string
		input    int
		output   bool
	}{
		{
			testcase: "Case 1",
			input:    11,
			output:   true,
		},
		{
			testcase: "Case 2",
			input:    13,
			output:   true,
		},
		{
			testcase: "Case 3",
			input:    17,
			output:   true,
		},
		{
			testcase: "Case 4",
			input:    20,
			output:   false,
		},
		{
			testcase: "Case 5",
			input:    35,
			output:   false,
		},
	}
	for _, value := range testcases {
		result := PrimeNumber(value.input)
		t.Run(value.testcase, func(t *testing.T) {
			assert.Equal(t, value.output, result)
		})
	}
}
