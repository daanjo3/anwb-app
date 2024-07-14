package main

import (
	"bytes"
	"net/http"
)

func GetAnwbData() ([]byte, error) {
	resp, err := http.Get("https://api.anwb.nl/v2/incidents?apikey=***REMOVED***&polylines=true&polylineBounds=true&totals=true")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
