package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaktorBilangan(t *testing.T) {
	var testcases = []struct {
		testcase string
		input    int
		output   string
	}{
		{
			testcase: "Case 1",
			input:    6,
			output:   "1\n2\n3\n6\n",
		},
		{
			testcase: "Case 2",
			input:    20,
			output:   "1\n2\n4\n5\n10\n20\n",
		},
	}
	for _, value := range testcases {
		result := FaktorBilangan(value.input)
		t.Run(value.testcase, func(t *testing.T) {
			assert.Equal(t, value.output, result)
		})
	}
}
