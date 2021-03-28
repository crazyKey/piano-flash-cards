package main

import (
	"math/rand"
	"time"
)

func main() {
	// Seed rand
	rand.Seed(time.Now().UnixNano())

	// Setup MIDI
	setupMIDI()

	// Start MIDI listener
	startMIDI()

	// Load notes
	loadNotes()

	// Display notes
	startUI()

	// Stop MIDI
	stopMIDI()
}

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
