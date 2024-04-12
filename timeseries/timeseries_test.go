package timeseries

import (
	"fmt"
	"testing"
)

func TestTimeseries(t *testing.T) {
	tMap := New()
	err := tMap.Add("2009-Aug-29", 2.2)
	if err != nil {
		t.Error("Error adding ", err.Error())
	}

	err = tMap.Add("2023-Aug-29", 9.4)
	if err != nil {
		t.Error("Error adding ", err.Error())
	}

	x, y := tMap.SortedCoordinates()
	fmt.Println(x, y)
}
