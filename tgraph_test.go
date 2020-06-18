package tgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFloat64(t *testing.T) {
	data := []float64{3.14, 6.27, 43.4, 68.9}
	m := minFloat64(data)
	assert.Equal(t, 3.14, m)
}

func TestMax(t *testing.T) {
	data := []float64{3.14, 6.27, 43.4, 68.9}
	m := maxFloat64(data)
	assert.Equal(t, 68.9, m)
}

func TestFindMinFloat64(t *testing.T) {
	data := [][]float64{{3.14, 6.28}, {6.27, 7.91}, {43.4, 4.33}, {68.9, 33.7}}
	m := findMinFloat64(data)
	assert.Equal(t, 3.14, m)
}

func TestFindMaxFloat64(t *testing.T) {
	data := [][]float64{{3.14, 6.28}, {6.27, 7.91}, {43.4, 4.33}, {68.9, 33.7}}
	m := findMaxFloat64(data)
	assert.Equal(t, 68.9, m)
}

func TestNormalize(t *testing.T) {
	data := [][]float64{{3.14, 6.28}, {6.27, 7.91}, {43.4, 4.33}, {68.9, 33.7}}
	expectedData := [][]float64{{2.27866473149492, 4.55732946298984}, {4.550072568940493, 5.7402031930333814},
		{31.494920174165454, 3.1422351233671986}, {50.0, 24.455732946298983}}
	result := Normalize(data, 50)
	assert.Equal(t, expectedData, result)
}

func TestChart(t *testing.T) {
	title := "eureka"
	labels1 := []string{"2020", "2021", "2020", "2023"}
	labels2 := []string{"testtest", "t", "testtf", "test"}

	data := [][]float64{{3.14, 6.28}, {6.27, 7.91}, {43.4, 4.33}, {68.9, 33.7}}
	colors := []string{"white", "red"}
	Chart(title, labels1, data, colors, 50, false, "|")
	Chart(title, labels2, data, colors, 50, true, "|")
}
