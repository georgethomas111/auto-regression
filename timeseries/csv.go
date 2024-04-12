package timeseries

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

func ReadCSV(d []byte) (Timeseries, error) {
	t := New()
	r := csv.NewReader(strings.NewReader(string(d)))
	labels := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if labels {
			// Nothing can be done
			labels = false
			continue
		}

		if err != nil {
			return nil, err
		}

		y, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}

		err = t.Add(record[0], y)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}
