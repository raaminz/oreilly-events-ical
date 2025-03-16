package events_test

import (
	_ "embed"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"ramin.tech/oreilly-events-ical/internal/events"
)

//go:embed testdata/events_simple.json
var examples []byte

func TestParseJSON(t *testing.T) {
	want := []*events.OreillyEvent{{
		ID:          "urn:orm:live-event-series:0636920090983:live-event:0642572011429:session:0642572011431",
		Topics:      "Software Architecture",
		Title:       "System Design Interview Boot Camp",
		Description: "Solve complex problems using a proven framework",
		StartTime:   time.Date(2025, 03, 04, 3, 00, 0, 0, time.UTC),
		EndTime:     time.Date(2025, 03, 04, 7, 00, 0, 0, time.UTC),
		Instructors: "Rohit Bhardwaj",
		Levels:      "intermediate",
	}, {
		ID:          "urn:orm:live-event-series:0636920090983:live-event:0642572011429:session:0642572011433",
		Topics:      "Software Architecture",
		Title:       "System Design Interview Boot Camp",
		Description: "Solve complex problems using a proven framework",
		StartTime:   time.Date(2025, 03, 05, 3, 00, 0, 0, time.UTC),
		EndTime:     time.Date(2025, 03, 05, 7, 00, 0, 0, time.UTC),
		Instructors: "Rohit Bhardwaj",
		Levels:      "intermediate",
	}, {
		ID:          "urn:orm:live-event-series:0636920080767:live-event:0642572011848:session:0642572011850",
		Topics:      "Software Architecture",
		Title:       "Architecture Decision Making by Example",
		Description: "A guide for architects and developers",
		StartTime:   time.Date(2025, 03, 04, 10, 00, 0, 0, time.UTC),
		EndTime:     time.Date(2025, 03, 04, 13, 00, 0, 0, time.UTC),
		Instructors: "Andrew Harmel-Law",
		Levels:      "intermediate",
	}}
	got, err := events.ParseJSON(examples)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot\n%v\nwant\n%v", msg(got), msg(want))
	}
}

func msg(events []*events.OreillyEvent) string {
	msg, _ := json.Marshal(events)
	return string(msg)
}
