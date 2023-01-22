package pkg

import (
	"encoding/json"
	"io"
	"time"

	"github.com/lvrach/testprof/internal/test2json"
	"golang.org/x/exp/slices"
)

type Timing struct {
	Package  string
	Test     string `json:",omitempty"`
	Duration time.Duration
}

type Splitter struct {
	Timings  []Timing
	Packages []string
	Parts    int
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

func (s *Splitter) List(part int) []string {
	slices.SortFunc(s.Timings, func(a, b Timing) bool {
		return a.Duration > b.Duration
	})

	var pgks []string
	for i, sample := range s.Timings {
		if part == i%s.Parts {
			pgks = append(pgks, sample.Package)
		}
	}

	return pgks
}
