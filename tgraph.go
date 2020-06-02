package tgraph

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const (
	tick = `▇`
)

// Chart Handle the normalization of data and the printing of the graph.
func Chart(title string, labels []string, data [][]float64, colors []string, width float64) {
	// Find longest name
	maxLengthSlice := maxLengthSlice(labels)
	fmtLabelPrefix := "%" + strconv.Itoa(maxLengthSlice) + "s: "

	// Normalize data
	normData := Normalize(data, width)

	// Print data
	if len(title) != 0 {
		fmt.Println(title)
	}
	for i, label := range labels {
		fmt.Printf(fmtLabelPrefix, label)
		for j, d := range normData[i] {
			for idx := 0; idx < int(d); idx++ {
				if j == 1 {
					fmt.Print(tick)
					continue
				}
				fmt.Print(tick)
			}
		}
		fmt.Printf(" %.1f", data[i][0])
		fmt.Println()
	}
}

// Normalize the data and return it.
func Normalize(data [][]float64, width float64) [][]float64 {
	minData := findMinFloat64(data)
	// We offset by the minimum if there's a negative.
	var offsetData [][]float64
	if minData < 0 {
		minData = math.Abs(minData)
		for _, dd := range data {
			var dlist []float64
			for _, d := range dd {
				dlist = append(dlist, d+minData)
			}
			offsetData = append(offsetData, dlist)
		}
	} else {
		offsetData = data
	}

	minData = findMinFloat64(offsetData)
	maxData := findMaxFloat64(offsetData)

	if maxData < width {
		return offsetData
	}

	// maxData / width is the value for a single tick. normFactor is the
	// inverse of this value
	// If you divide a number to the value of single tick, you will find how
	// many ticks it does contain basically.
	normFactor := width / maxData
	var normData [][]float64
	for _, dd := range offsetData {
		var dlist []float64
		for _, d := range dd {
			dlist = append(dlist, d*normFactor)
		}
		normData = append(normData, dlist)
	}

	return normData
}

// Return the minimum value in list.
func minFloat64(v []float64) float64 {
	var m float64
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

// Return the minimum value in sublist of list.
func findMinFloat64(vv [][]float64) float64 {
	var m float64
	for i, v := range vv {
		e := minFloat64(v)
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

// Return the maximum value in list.
func maxFloat64(v []float64) float64 {
	var m float64
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

// Return the maximum value in sublist of list.
func findMaxFloat64(vv [][]float64) float64 {
	var m float64
	for i, v := range vv {
		e := maxFloat64(v)
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func maxLengthSlice(labels []string) int {
	if len(labels) == 0 {
		return 0
	}
	sort.Strings(labels)
	return len(labels[len(labels)-1])
}
