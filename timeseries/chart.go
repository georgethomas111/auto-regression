package timeseries

import (
	"fmt"
	"html/template"
	"os"
)

type LineChart struct {
	Title string
	X     []string
	Y     []float64
}

type Charts struct {
	Title string
	Story string

	Charts []LineChart
}

func NewCharts(title string, d []byte) (*Charts, error) {
	return &Charts{}, nil
}

// Render renders the chart based on the input parameters
func (c *Charts) Render() {
	t, err := template.New("webpage").Parse(ChartTmp)
	err = t.Execute(os.Stdout, c)
	if err != nil {
		fmt.Println(" Error =", err.Error())
	}
}
