package main

import "ramin.tech/oreilly-events-ical/cmd"

func main() {
	err := cmd.Main()
	if err != nil {
		panic(err)
	}
}
