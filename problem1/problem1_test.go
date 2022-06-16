package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKonversiNilai(t *testing.T) {
	var testcases = []struct {
		testcase string
		input    int
		output   string
	}{
		{
			testcase: "test 70",
			input:    70,
			output:   "B+",
		},
		{
			testcase: "test 80",
			input:    80,
			output:   "A",
		},
	}
	for _, value := range testcases {
		result := KonversiNilai(value.input)
		t.Run(value.testcase, func(t *testing.T) {
			assert.Equal(t, value.output, result)
		})
	}
}
