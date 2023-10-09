package timeseries

// FIXME: In future maybe this can be changed to time.Time but this is a good start
type Timeseries map[string]float64

func New() Timeseries {
	t := make(map[string]float64)
	return t
}
