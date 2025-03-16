package ical

import (
	"fmt"
	"strings"

	ics "github.com/arran4/golang-ical"
	"ramin.tech/oreilly-events-ical/internal/events"
)

type ical struct {
	cal *ics.Calendar
}

func NewIcal() *ical {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)
	return &ical{cal}
}

func (i *ical) AddEvent(event *events.OreillyEvent) {
	e := i.cal.AddEvent(fmt.Sprintf("%s@%s",
		"github.com/raaminz/oreilly-events-ical/1", event.ID))
	e.SetStatus(ics.ObjectStatusConfirmed)
	e.SetCreatedTime(event.StartTime)
	e.SetDtStampTime(event.StartTime)
	e.SetStartAt(event.StartTime)
	e.SetEndAt(event.EndTime)
	e.SetSummary(event.Title)
	e.SetDescription(generateDescription(event))
}

func (i *ical) Serialize() string {
	return i.cal.Serialize()
}

func generateDescription(event *events.OreillyEvent) string {
	var builder strings.Builder
	if event.Topics != "" {
		builder.WriteString("Topic: ")
		builder.WriteString(event.Topics)
		builder.WriteString("\n\n")
	}
	if event.Description != "" {
		builder.WriteString(event.Description)
	}
	if event.Levels != "" {
		builder.WriteString("\n\n")
		builder.WriteString("Level: ")
		builder.WriteString(event.Levels)
		builder.WriteRune('\n')
	}
	if event.Instructors != "" {
		builder.WriteString("Instructor: ")
		builder.WriteString(event.Instructors)
	}
	return builder.String()
}
