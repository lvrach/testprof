package test2json

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Record struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Output  string    `json:"Output"`
	Elapsed float64   `json:"Elapsed"`
}

func Parse(reader io.Reader) ([]Record, error) {
	var records []Record

	d := json.NewDecoder(reader)
	for {
		var v Record
		err := d.Decode(&v)
		if err == io.EOF {
			break // done decoding file
		}
		if err != nil {
			return nil, fmt.Errorf("parsing error: %w", err)
		}
		records = append(records, v)
	}

	return records, nil
}
