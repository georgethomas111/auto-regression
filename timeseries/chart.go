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
	Index int
}

func Chart(title string, t Timeseries, index int) LineChart {
	lineChart := LineChart{
		Title: title,
		X:     []string{},
		Y:     []float64{},
		Index: index,
	}

	xt, y := t.SortedCoordinates()

	for _, x := range xt {
		timeStr := fmt.Sprintf("%s-%d", x.Month(), x.Day())
		lineChart.X = append(lineChart.X, timeStr)
	}
	lineChart.Y = y

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
