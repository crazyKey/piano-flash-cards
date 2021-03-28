package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	uiApp    fyne.App
	uiWindow fyne.Window
	uiImage  *canvas.Image
	uiLabel  *widget.Label
)

func startUI() {
	// Create new app
	uiApp = app.New()

	// Create window
	uiWindow = uiApp.NewWindow("Piano Flash Cards")

	// Resize window
	uiWindow.Resize(fyne.NewSize(667, 826))

	// Initial label and image
	uiLabel = widget.NewLabel("Play any note to start")
	uiImage = &canvas.Image{FillMode: canvas.ImageFillOriginal}

	// Set window content
	uiWindow.SetContent(container.NewVBox(
		uiLabel,
		uiImage,
	))

	// Run app
	uiWindow.ShowAndRun()
}

func updateNoteImage() {
	// Update image file
	uiImage.File = fmt.Sprintf("images/%s/%s", currentNote.Clef, currentNote.Image)

	// Refresh image in canvas
	canvas.Refresh(uiImage)

	// Update label
	uiLabel.SetText("Play the correct key")
}

func wrongKey(i string) {
	// Update label
	uiLabel.SetText(fmt.Sprintf("That's wrong, the note is %s", i))
}
