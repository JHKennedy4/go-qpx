package qpx

import (
	"bytes"
	"encoding/json"
	"io"
)

type request struct {
	Request query `json:"request"`
}

type query struct {
	Passengers passengers   `json:"passengers"`
	Slice      [2]departure `json:"slice"`
	MaxPrice   string       `json:"maxPrice"`
	Solutions  int          `json:"solutions"`
}

type passengers struct {
	Adults int `json:"adultCount"`
}

type departure struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
}

func NewQuery(origin string, destination string, start_date string, end_date string) request {
	passenger_count := passengers{1}

	outbound := departure{origin, destination, start_date}
	inbound := departure{destination, origin, end_date}

	departures := [2]departure{outbound, inbound}

	q := query{passenger_count, departures, "USD5400", 3}

	r := request{q}

	return r
}

func (query *request) ToJSON() io.Reader {
	json, err := json.Marshal(*query)
	if err != nil {
		panic(err)
	}

	b := bytes.NewBuffer(json)

	return b
}

func (query *request) PrettyPrintJSON() io.Reader {
	json, err := json.MarshalIndent(*query, "", "   ")
	if err != nil {
		panic(err)
	}

	b := bytes.NewBuffer(json)

	return b
}
