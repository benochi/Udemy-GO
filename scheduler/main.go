package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Go Store")

	textarea := widget.NewMultiLineEntry()

	// Load the file content into textarea when the application starts
	content, err := os.ReadFile("gostore.txt")
	if err != nil && !os.IsNotExist(err) { // It's okay if the file doesn't exist yet
		log.Fatal(err)
	}
	textarea.SetText(string(content))

	saveButton := widget.NewButton("Save", func() {
		err := os.WriteFile("gostore.txt", []byte(textarea.Text), 0644)
		if err != nil {
			log.Fatal(err)
		}
	})

	myWindow.SetContent(container.NewVBox(textarea, saveButton))
	myWindow.ShowAndRun()
}
