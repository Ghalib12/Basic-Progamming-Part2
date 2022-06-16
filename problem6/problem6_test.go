package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullPrime(t *testing.T) {
	var testcases = []struct {
		testcase string
		input    int
		output   bool
	}{
		{
			testcase: "Case 1",
			input:    2,
			output:   true,
		},
		{
			testcase: "Case 2",
			input:    3,
			output:   true,
		},
		{
			testcase: "Case 3",
			input:    11,
			output:   false,
		},
		{
			testcase: "Case 4",
			input:    13,
			output:   false,
		},
		{
			testcase: "Case 5",
			input:    23,
			output:   true,
		},
		{
			testcase: "Case 6",
			input:    29,
			output:   false,
		},
	}
	for _, value := range testcases {
		result := FullPrima(value.input)
		t.Run(value.testcase, func(t *testing.T) {
			assert.Equal(t, value.output, result)
		})
	}
}
