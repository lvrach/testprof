package pkg

import (
	"encoding/json"
	"io"
	"time"

	"github.com/lvrach/testprof/internal/test2json"
)

type Timing struct {
	Package  string
	Test     string `json:",omitempty"`
	Duration time.Duration
}

func Timings(rr []test2json.Record) []Timing {
	var timings []Timing
	for _, r := range rr {
		if r.Action == "pass" && r.Test == "" {
			timings = append(timings, Timing{
				Package:  r.Package,
				Test:     r.Test,
				Duration: time.Duration(r.Elapsed * float64(time.Second)),
			})
		}
	}
	return timings
}

func Parse(reader io.Reader) ([]Timing, error) {
	var timings []Timing
	err := json.NewDecoder(reader).Decode(&timings)
	return timings, err
}
