package util_test

import (
	"myapp/crossCutting/util"
	"testing"

	"gotest.tools/assert"
)

func TestIsNilEmptyOrWhiteSpace(t *testing.T) {

	testcases := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"     ", true},
		{"av", false},
	}

	for _, test := range testcases {
		got := util.IsNilEmptyOrWhiteSpace(test.input)
		assert.Equal(t, test.expected, got)
	}
}
