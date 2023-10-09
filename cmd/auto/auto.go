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

	t, err := timeseries.Read(d)
	if err != nil {
		fmt.Println("Could not initialize chart ", err.Error())
		return
	}
	fmt.Println("Lets print t", t)
}
