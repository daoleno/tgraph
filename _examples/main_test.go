package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestResolveNumber(t *testing.T)  {
	records := [][]string{
		{"86.1", "67.1", "66.7", "62.9", "62.3"},
		{"13.9", "32.9", "33.3", "37.1", "37.3"}}
	expectedRecords := [][]float64{{86.1, 13.9}, {67.1, 32.9}, {66.7, 33.3}, {62.9, 37.1}, {62.3, 37.3}}

	numbers := resolveNumber(records)
	assert.Equal(t, expectedRecords, numbers)
}
