package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/nxadm/tail"
)

var logFile = "/home/boris/Code/golang/iis_monitor/test/test.log"

func main() {
	t, err := tail.TailFile(logFile, tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}

	var i = 0
	// var codes = make(map[string][]int)
	// сделать структуры, одно поле - код, второе map из значений, во время перезаписи, брать слайс в 10
	// последних значений, перезаписывать поле и возвращать в график
	type code struct {
		name         string
		currentValue int
		values       []int
	}

	type codes struct {
		code []code
	}

	allCodes := new(codes)
	for line := range t.Lines {
		splitedString := strings.Split(line.Text, " ")
		i++
		// codes[splitedString[len(splitedString)-1]]

		if i == 10 {

		}
	}

	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}
