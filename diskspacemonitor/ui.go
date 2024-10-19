package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/zap"
)

func InitUI(logger *zap.SugaredLogger) {
	logger.Debug("Initializing UI")

	newApp := app.New()
	window := newApp.NewWindow(appNameWithVersion)
	window.Resize(fyne.NewSize(600, 400))

	// Bottom grid
	saveButton := widget.NewButton("Save Settings", func() {
		logger.Debug("Save button clicked")
	})
	cancelButton := widget.NewButton("Cancel", func() {
		logger.Debug("Cancel button clicked")
		window.Hide()
	})
	bottomGrid := container.New(
		layout.NewGridLayout(2),
		cancelButton,
		saveButton,
	)

	// Content
	hourSelect := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Hour select set to", value)
	})
	minuteSelect := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Minute select set to", value)
	})
	timeScheduleRow := container.New(layout.NewGridLayout(2), hourSelect, minuteSelect)
	timeScheduleLabelRow := container.New(
		layout.NewGridLayout(2),
		widget.NewLabel("Hour"),
		widget.NewLabel("Minute"),
	)
	content := container.New(
		layout.NewVBoxLayout(),
		timeScheduleLabelRow,
		timeScheduleRow,
		layout.NewSpacer(),
	)

	// Assemble window components
	window.SetContent(container.New(
		layout.NewVBoxLayout(),
		content,
		bottomGrid,
	))
	window.ShowAndRun()
}
