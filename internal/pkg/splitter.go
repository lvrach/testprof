package pkg

import (
	"time"

	"golang.org/x/exp/slices"
)

type Splitter struct {
	Timings  []Timing
	Packages []string
	Parts    int
}

type packageDuration struct {
	Package  string
	Duration time.Duration
}

func (s *Splitter) List(part int) []string {
	durations := make(map[string]time.Duration)
	for _, t := range s.Timings {
		durations[t.Package] = t.Duration
	}

	pds := make([]packageDuration, len(s.Packages))
	for i := range pds {
		pds[i] = packageDuration{
			Package:  s.Packages[i],
			Duration: durations[s.Packages[i]],
		}
	}

	slices.SortStableFunc(pds, func(a, b packageDuration) bool {
		return a.Duration > b.Duration
	})

	var packages []string
	for i, p := range pds {
		if part == i%s.Parts {
			packages = append(packages, p.Package)
		}
	}

	return packages
}
