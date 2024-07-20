package service

import "github.com/daanjo3/anweb-app/api/internal/anwb"

func ListJams(document anwb.Document) []anwb.RoadEvent {
	return listEvents(document, func(rs anwb.RoadSegment) []anwb.RoadEvent { return rs.Jams })
}

func ListRoadWorks(document anwb.Document) []anwb.RoadEvent {
	return listEvents(document, func(rs anwb.RoadSegment) []anwb.RoadEvent { return rs.Roadworks })
}

func ListRadars(document anwb.Document) []anwb.RoadEvent {
	return listEvents(document, func(rs anwb.RoadSegment) []anwb.RoadEvent { return rs.Radars })
}

type infoLookup func(anwb.RoadSegment) []anwb.RoadEvent

func listEvents(document anwb.Document, lookup infoLookup) []anwb.RoadEvent {
	allRoadInfo := make([]anwb.RoadEvent, 0)
	for _, road := range document.Roads {
		for _, segment := range road.Segments {
			roadInfo := lookup(segment)
			if len(roadInfo) > 0 {
				allRoadInfo = append(allRoadInfo, roadInfo...)
			}
		}
	}
	return allRoadInfo
}
