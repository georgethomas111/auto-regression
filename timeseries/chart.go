package timeseries

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

type LineChart struct {
	Title string
	X     []string
	Y     []float64
}

func Chart(title string, t Timeseries) LineChart {
	lineChart := LineChart{
		Title: title,
		X:     []string{},
		Y:     []float64{},
	}

	for x, y := range t {
		lineChart.X = append(lineChart.X, x)
		lineChart.Y = append(lineChart.Y, y)
	}

	return lineChart
}

type HTML struct {
	Title string
	Story string

	Charts []LineChart
}

func NewHTML(title string, story string) (*HTML, error) {
	return &HTML{
		Title:  title,
		Story:  story,
		Charts: []LineChart{},
	}, nil
}

// Render renders the chart based on the input parameters
func (h *HTML) Render(outName string) {
	var b []byte
	buf := bytes.NewBuffer(b)
	t, err := template.New("webpage").Parse(ChartTmp)
	err = t.Execute(buf, h)
	if err != nil {
		fmt.Println(" Error =", err.Error())
	}

	err = os.WriteFile(outName, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println(" Writing file Error =", err.Error())
	}
}
