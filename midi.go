package main

import (
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/rtmididrv"
)

var (
	in midi.In
)

func setupMIDI() {
	// MIDI driver
	drv, err := rtmididrv.New()
	must(err)

	// Close port at the end
	defer drv.Close()

	// Get all input ports
	ins, err := drv.Ins()
	must(err)

	// Take the first port
	in = ins[0]
}

func startMIDI() {
	// Open port
	must(in.Open())

	// Reader for MIDI input
	rd := reader.New(
		reader.NoLogger(),
		reader.Each(func(pos *reader.Position, msg midi.Message) {
			// Ignore everything but NoteOn
			_, ok := msg.(channel.NoteOn)
			if !ok {
				return
			}

			// Get keyboard key
			k := int(msg.(channel.NoteOn).Key())

			// Check with current note
			r := checkNote(k)

			// Give note position hint if wrong
			if r != "" {
				wrongKey(r)
				return
			}

			// Get next random note and display it
			randomNote()
			updateNoteImage()
		}),
	)

	// listen for MIDI
	err := rd.ListenTo(in)
	must(err)
}

func stopMIDI() {
	// Stop reading MIDI input
	err := in.StopListening()
	must(err)

	// Close port
	err = in.Close()
	must(err)
}
