package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/georgethomas111/auto-regression/timeseries"
)

func main() {
	csvInput := os.Args[1]

	d, err := ioutil.ReadFile(csvInput)
	if err != nil {
		fmt.Println("Could not read csvFile passed ", csvInput, " error = ", err.Error())
		return
	}

	t, err := timeseries.ReadCSV(d)
	if err != nil {
		fmt.Println("Could not initialize chart ", err.Error())
		return
	}

	chartIndex := 0
	c0 := timeseries.Chart("first chart", t, chartIndex)
	chartIndex = 1
	c1 := timeseries.Chart("second chart", t, chartIndex)
	html, _ := timeseries.NewHTML("Title", "No story yet")
	html.Charts = append(html.Charts, c0)
	html.Charts = append(html.Charts, c1)
	html.Render("./latest.html")
	timeseries.StartServer()
}
