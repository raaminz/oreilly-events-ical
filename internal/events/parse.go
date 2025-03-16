package events

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func ParseJSON(data []byte) ([]*OreillyEvent, error) {
	events := gjson.Get(string(data), "data.liveEvents.results")
	if !events.Exists() {
		return nil, errors.New("error parsing json, data.liveEvents.results")
	}

	allEvents := []*OreillyEvent{}
	for _, e := range events.Array() {
		param := convertParam{}
		param.title = e.Get("title").Str
		param.desc = e.Get("shortDescription").Str
		param.topic = joinStrFromArray(e.Get("topics").Array(), "name")
		param.contributor = joinStrFromArray(e.Get("contributors").Array(),
			"fullName")
		param.level = joinStrFromArray(e.Get("contentLevels").Array(), "")
		param.sessions = getSessions(e.Get("sessions").Array())

		oreillyEvents, err := convertToOreillyEvents(&param)
		if err != nil {
			log.Printf("error parsing an event. error %v %v\n", err, e.Raw)
			continue
		}
		allEvents = append(allEvents, oreillyEvents...)
	}

	return allEvents, nil
}

type session struct {
	ourn      string
	startTime string
	endTime   string
}

type convertParam struct {
	title       string
	desc        string
	sessions    []*session
	topic       string
	contributor string
	level       string
}

func convertToOreillyEvents(param *convertParam) ([]*OreillyEvent, error) {
	if param.title == "" {
		return nil, errors.New("no title for the event")
	}
	events := []*OreillyEvent{}
	for _, p := range param.sessions {
		if p.ourn == "" {
			return nil, errors.New("no ourn for the event")
		}
		startTime, err := parseTime(p.startTime)
		if err != nil {
			return nil, err
		}
		endTime, err := parseTime(p.endTime)
		if err != nil {
			return nil, err
		}
		events = append(events, &OreillyEvent{
			ID:          p.ourn,
			Topics:      param.topic,
			Title:       param.title,
			Description: strings.TrimSpace(param.desc),
			StartTime:   startTime,
			EndTime:     endTime,
			Instructors: param.contributor,
			Levels:      param.level})
	}
	return events, nil
}

func getSessions(array []gjson.Result) []*session {
	sessions := []*session{}
	for _, v := range array {
		sessions = append(sessions, &session{
			ourn:      v.Get("ourn").Str,
			startTime: v.Get("startTime").Str,
			endTime:   v.Get("endTime").Str,
		})
	}
	return sessions
}

func joinStrFromArray(array []gjson.Result, key string) string {
	result := []string{}
	for _, v := range array {
		if key == "" && v.Exists() {
			result = append(result, v.Str)
		}
		if v.Get(key).Exists() {
			result = append(result, v.Get(key).Str)
		}
	}
	return strings.Join(result, ", ")
}

func parseTime(input string) (time.Time, error) {
	const layout = "2006-01-02T15:04:05Z07:00"
	return time.Parse(layout, input)
}
