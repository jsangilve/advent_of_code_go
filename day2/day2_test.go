package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRow(t *testing.T) {
	result := parseRow("1-3 b: cdefg")
	assert.Equal(t, &passwordPolicy{Min: 1, Max: 3, Char: "b", Password: "cdefg"}, result)
}
