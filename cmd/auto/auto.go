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

	c := timeseries.Chart("first chart", t)
	html, _ := timeseries.NewHTML("Title", "No story yet")
	html.Charts = append(html.Charts, c)
	fmt.Println("html is", html)

}
