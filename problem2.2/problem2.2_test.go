package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaktorBilanganDesc(t *testing.T) {
	var testcases = []struct {
		testcase string
		input    int
		output   string
	}{
		{
			testcase: "Case 1",
			input:    6,
			output:   "6\n3\n2\n1\n",
		},
		{
			testcase: "Case 2",
			input:    20,
			output:   "20\n10\n5\n4\n2\n1\n",
		},
	}
	for _, value := range testcases {
		result := FaktorBilanganDesc(value.input)
		t.Run(value.testcase, func(t *testing.T) {
			assert.Equal(t, value.output, result)
		})
	}
}
