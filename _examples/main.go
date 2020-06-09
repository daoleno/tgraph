package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/daoleno/tgraph"
)

var (
	title   string
	width   float64
	color   string
	stacked bool
	tick    string
)

var flagSet *flag.FlagSet

func init() {
	flagSet = flag.NewFlagSet("", flag.ExitOnError)
	flagSet.StringVar(&title, "title", "", "title of graph")
	flagSet.Float64Var(&width, "width", 50, "width of graph in characters default:50")
	flagSet.StringVar(&color, "color", "", "support red,blue,green,magenta,yellow,black,cyan")
	flagSet.BoolVar(&stacked, "stacked", false, "stacked bar graph default:false")
	flagSet.StringVar(&tick, "custom-tick", "▇", "custom tick")
}

func main() {
	usage := `
Usage: go run main.go data.csv
`
	flagSet.Parse(os.Args[2:])

	if os.Args[1] == "help" {
		fmt.Println(usage)
		return
	}

	records, err := readCSV(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	labels := records[0]
	data := resolveNumber(records[1:])
	colors := strings.Split(color, ",")

	tgraph.Chart(title, labels, data, colors, width, stacked, tick)
}

func readCSV(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func resolveNumber(records [][]string) [][]float64 {
	var data [][]float64
	for i, r := range records {
		for j, n := range r {
			nfloat64, err := strconv.ParseFloat(n, 64)
			if err != nil {
				panic(err)
			}

			if i == 0 {
				data = append(data, []float64{nfloat64})
				continue
			}
			data[j] = append(data[j], nfloat64)
		}
	}
	return data
}
