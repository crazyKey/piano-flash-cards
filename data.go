package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
)

var (
	notes       []Note
	currentNote *Note
)

type Note struct {
	Name  string `json:"name"`
	Clef  string `json:"clef"`
	Image string `json:"image"`
	Key   int    `json:"key"`
}

func loadNotes() {
	// Open notes json
	jNotes, err := os.Open("notes.json")
	must(err)

	// Close file
	defer jNotes.Close()

	// Read file
	bNotes, _ := ioutil.ReadAll(jNotes)

	// Unmarshall
	err = json.Unmarshal(bNotes, &notes)
	must(err)
}

func randomNote() {
	// Pick a random note
	currentNote = &notes[rand.Intn(len(notes)-1)]
}

func checkNote(n int) string {
	// Note is correct
	if currentNote == nil || n == currentNote.Key {
		return ""
	}

	// Lower or higher hint
	if currentNote.Key < n {
		return "lower"
	}
	return "higher"
}
