package timeseries

import (
	"sort"
	"time"
)

// FIXME: In future maybe this can be changed to time.Time but this is a good start
type Timeseries map[time.Time]float64

func (t Timeseries) Add(strTime string, pt float64) error {
	var parsedTime time.Time
	var err error

	// List of layout formats to try
	layoutFormats := []string{
		"2006-Jan-02",
		"2006-01-02",
		"2006/01/02",
		"Jan-02-2006",
		"01-02-2006",
		"01/02/2006",
	}

	// Try parsing the time with each layout format until successful
	for _, layout := range layoutFormats {
		parsedTime, err = time.Parse(layout, strTime)
		if err == nil {
			break // Exit loop if parsing successful
		}
	}

	// If all layout formats failed to parse the time
	if err != nil {
		return err
	}

	t[parsedTime] = pt
	return nil
}

func (t Timeseries) SortedCoordinates() ([]time.Time, []float64) {
	// Extract keys (timestamps) from the map
	var timestamps []time.Time
	for timestamp := range t {
		timestamps = append(timestamps, timestamp)
	}

	// Sort the timestamps
	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i].Before(timestamps[j])
	})

	// Get the corresponding values (points) for the sorted timestamps
	var points []float64
	for _, timestamp := range timestamps {
		points = append(points, t[timestamp])
	}

	return timestamps, points
}

func New() Timeseries {
	t := make(map[time.Time]float64)
	return t
}
