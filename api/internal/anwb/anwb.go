package anwb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UploadedAt primitive.DateTime `json:"_uploaded_at,omitempty" bson:"_uploaded_at,omitempty"`
	Success    bool               `json:"success,omitempty" bson:"success,omitempty"`
	Roads      []Road             `json:"roads,omitempty" bson:"roads,omitempty"`
}

type IndexEntry struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	UploadedAt primitive.DateTime `json:"_uploaded_at,omitempty"`
}

func (doc *Document) AsIndexEntry() IndexEntry {
	return IndexEntry{
		Id:         doc.Id,
		UploadedAt: doc.UploadedAt,
	}
}

type Road struct {
	Road     string        `json:"road,omitempty" bson:"road,omitempty"`
	RoadType string        `json:"type,omitempty" bson:"type,omitempty"`
	Segments []RoadSegment `json:"segments,omitempty" bson:"segments,omitempty"`
}

type RoadSegment struct {
	Start     string     `json:"start,omitempty" bson:"start,omitempty"`
	End       string     `json:"end,omitempty" bson:"end,omitempty"`
	Jams      []RoadInfo `json:"jams,omitempty" bson:"jams,omitempty"`
	Radars    []RoadInfo `json:"radars,omitempty" bson:"radars,omitempty"`
	Roadworks []RoadInfo `json:"roadworks,omitempty" bson:"roadworks,omitempty"`
}

type RoadInfo struct {
	Id            int      `json:"id,omitempty" bson:"id,omitempty"`
	Road          string   `json:"road,omitempty" bson:"road,omitempty"`
	SegmentId     int      `json:"segmentId,omitempty" bson:"segmentId,omitempty"`
	CodeDirection int      `json:"codeDirection,omitempty" bson:"codeDirection,omitempty"`
	Type          string   `json:"type,omitempty" bson:"type,omitempty"`
	Afrc          int      `json:"afrc,omitempty" bson:"afrc,omitempty"`
	Category      string   `json:"category,omitempty" bson:"category,omitempty"`
	IncidentType  string   `json:"incidentType,omitempty" bson:"incidentType,omitempty"`
	From          string   `json:"from,omitempty" bson:"from,omitempty"`
	FromLoc       Location `json:"fromLoc,omitempty" bson:"fromLoc,omitempty"`
	To            string   `json:"to,omitempty" bson:"to,omitempty"`
	ToLoc         Location `json:"toLoc,omitempty" bson:"toLoc,omitempty"`
	Polyline      string   `json:"polyline,omitempty" bson:"polyline,omitempty"`
	HM            float32  `json:"HM,omitempty" bson:"HM,omitempty"`
	Bounds        struct {
		SouthWest Location `json:"southWest,omitempty" bson:"southWest,omitempty"`
		NorthEast Location `json:"northEast,omitempty" bson:"northEast,omitempty"`
	} `json:"bounds,omitempty" bson:"bounds,omitempty"`
	Events []struct {
		AlertC int    `json:"alertC,omitempty" bson:"alertC,omitempty"`
		Text   string `json:"text,omitempty" bson:"text,omitempty"`
	} `json:"events,omitempty" bson:"events,omitempty"`
	Distance int    `json:"distance,omitempty" bson:"distance,omitempty"`
	Delay    int    `json:"delay,omitempty" bson:"delay,omitempty"`
	Start    string `json:"start,omitempty" bson:"start,omitempty"` // UTC timestamp
	Stop     string `json:"stop,omitempty" bson:"stop,omitempty"`   // UTC timestamps
	Reason   string `json:"reason,omitempty" bson:"reason,omitempty"`
}

type Location struct {
	Lat float32 `json:"lat,omitempty" bson:"lat,omitempty"`
	Lon float32 `json:"lon,omitempty" bson:"lon,omitempty"`
}

type Totals struct {
	A     TotalEntry `json:"a,omitempty" bson:"a,omitempty"`
	N     TotalEntry `json:"n,omitempty" bson:"n,omitempty"`
	Other TotalEntry `json:"other,omitempty" bson:"other,omitempty"`
	All   TotalEntry `json:"all,omitempty" bson:"all,omitempty"`
}

type TotalEntry struct {
	Distance int `json:"distance,omitempty" bson:"distance,omitempty"`
	Delay    int `json:"delay,omitempty" bson:"delay,omitempty"`
	Count    int `json:"count,omitempty" bson:"count,omitempty"`
}

var requestUrl *string

func getRequestUrl() string {
	if requestUrl != nil {
		return *requestUrl
	}
	apiKey := os.Getenv("ANWB_API_KEY")
	if apiKey == "" {
		log.Fatal("No API key provided")
	}
	url := fmt.Sprintf("https://api.anwb.nl/v2/incidents?apikey=%s&polylines=true&polylineBounds=true&totals=true", apiKey)
	requestUrl = &url
	return *requestUrl
}

func Get() (Document, error) {
	resp, err := http.Get(getRequestUrl())
	if err != nil {
		return Document{}, err
	}
	defer resp.Body.Close()
	var document Document
	if err := json.NewDecoder(resp.Body).Decode(&document); err != nil {
		return Document{}, err
	}
	return document, nil
}
