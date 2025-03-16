package ical_test

import (
	"testing"
	"time"

	"ramin.tech/oreilly-events-ical/internal/events"
	"ramin.tech/oreilly-events-ical/internal/ical"
)

const expected = `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//arran4//Golang ICS Library
METHOD:PUBLISH
BEGIN:VEVENT
UID:github.com/raaminz/oreilly-events-ical/1@123
STATUS:CONFIRMED
CREATED:20221226T000000Z
DTSTAMP:20221226T000000Z
DTSTART:20221226T000000Z
DTEND:20221226T040000Z
SUMMARY:Course Title
DESCRIPTION:Topic: programming\n\nHow to make a complicated
  system\n\nLevel: beginner\nInstructor: Somebody
END:VEVENT
END:VCALENDAR
`

func TestSerialize(t *testing.T) {
	event := &events.OreillyEvent{}
	event.ID = "123"
	event.Title = "Course Title"
	event.Description = "How to make a complicated system"
	event.Levels = "beginner"
	event.Instructors = "Somebody"
	event.Topics = "programming"

	now := time.Date(2022, 12, 25, 0, 0, 0, 0, time.UTC)
	event.StartTime = now.Add(time.Hour * 24)
	event.EndTime = now.Add(time.Hour * 28)

	cal := ical.NewIcal()
	cal.AddEvent(event)
	got := cal.Serialize()
	if got != expected {
		t.Errorf("want %s got %s", expected, got)
	}

}
