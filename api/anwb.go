package main

import (
	"bytes"
	"net/http"
)

type AnwbDoc struct {
	Id         string `json:"id" bson:"_id"`
	UploadedAt string
	Success    bool
	Roads      []Road
}

type Road struct {
	Road     string
	RoadType string `json:"type" bson:"type"`
	Segments []RoadSegment
}

type RoadSegment struct {
	Start     string
	End       string
	Jams      []RoadInfo
	Radars    []RoadInfo
	Roadworks []RoadInfo
}

type RoadInfo struct {
	Id            int
	Road          string
	SegmentId     int
	CodeDirection int
	Type          string
	Afrc          int
	Category      string
	IncidentType  string
	From          string
	FromLoc       Location
	To            string
	ToLoc         Location
	Polyline      string
	HM            float32
	Bounds        struct {
		SouthWest Location
		NorthEast Location
	}
	Events []struct {
		AlertC int
		Text   string
	}
	Distance int
	Delay    int
	Start    string // UTC timestamp
	Stop     string // UTC timestamps
	Reason   string
}

type Location struct {
	lat float32
	lon float32
}

type Totals struct {
	A     TotalEntry
	N     TotalEntry
	Other TotalEntry
	All   TotalEntry
}

type TotalEntry struct {
	Distance int
	Delay    int
	Count    int
}

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
