package cmd

import (
	"log"
	"os"

	"ramin.tech/oreilly-events-ical/internal/events"
	"ramin.tech/oreilly-events-ical/internal/ical"
)

func Main() error {
	data, err := os.ReadFile("events.json")
	if err != nil {
		return err
	}
	events, err := events.ParseJSON(data)
	if err != nil {
		return err
	}
	cal := ical.NewIcal()
	for _, e := range events {
		cal.AddEvent(e)
	}
	return writeToFile("oreilly-events.ics", cal.Serialize())
}

func writeToFile(fileName string, data string) error {
	log.Println("writing to file " + fileName)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
