package tgraph

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/fatih/color"
)

const (
	defaultTick = `▇`
)

var colorMap = map[string]color.Attribute{
	"black":   color.FgHiBlack,
	"red":     color.FgHiRed,
	"green":   color.FgHiGreen,
	"yellow":  color.FgHiYellow,
	"blue":    color.FgHiBlue,
	"magenta": color.FgHiMagenta,
	"cyan":    color.FgHiCyan,
	"white":   color.FgHiWhite,
}

// Chart Handle the normalization of data and the printing of the graph.
func Chart(title string, labels []string, data [][]float64, colors []string, width float64, stacked bool, tick string) {
	// Set tick
	if len(tick) == 0 {
		tick = defaultTick
	}

	// Find longest name
	maxLengthSlice := maxLengthSlice(labels)
	fmtLabelPrefix := "%" + strconv.Itoa(maxLengthSlice) + "s"

	// Normalize data
	normData := Normalize(data, width)

	// Print data
	if len(title) != 0 {
		fmt.Println(title)
	}
	for i, label := range labels {
		var totalData float64
		for j, d := range normData[i] {
			totalData += data[i][j]
			if stacked {
				if j == 0 {
					fmt.Printf(fmtLabelPrefix+": ", label)
				}
				if len(colors) == len(normData[i]) {
					color.Set(colorMap[colors[j]])
				}
				for idx := 0; idx < int(d); idx++ {
					fmt.Print(tick)
				}
				color.Unset()
			} else {
				if j == 0 {
					fmt.Printf(fmtLabelPrefix+": ", label)
				} else {
					fmt.Printf(fmtLabelPrefix+"  ", "")
				}
				if len(colors) == len(normData[i]) {
					color.Set(colorMap[colors[j]])
				}
				for idx := 0; idx < int(d); idx++ {
					fmt.Print(tick)
				}
				color.Unset()
				fmt.Printf(" %.1f\n", data[i][j])
			}
		}
		if stacked {
			fmt.Printf(" %.1f\n", totalData)
		}
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
	sort.Sort(ByLength(labels))
	return len(labels[len(labels)-1])
}

// ByLength In order to sort by the length of the string
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
