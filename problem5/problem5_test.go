package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPangkat(t *testing.T) {
	var testcases = []struct {
		testcase string
		input    []int
		output   int
	}{
		{
			testcase: "Case 1",
			input:    []int{2, 3},
			output:   8,
		},
		{
			testcase: "Case 2",
			input:    []int{5, 3},
			output:   125,
		},
		{
			testcase: "Case 3",
			input:    []int{10, 2},
			output:   100,
		},
		{
			testcase: "Case 4",
			input:    []int{2, 5},
			output:   32,
		},
		{
			testcase: "Case 5",
			input:    []int{7, 3},
			output:   343,
		},
	}
	for _, value := range testcases {
		result := Pangkat(value.input[0], value.input[1])
		t.Run(value.testcase, func(t *testing.T) {
			assert.Equal(t, value.output, result)
		})
	}
}
