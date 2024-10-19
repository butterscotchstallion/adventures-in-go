package main

import (
	"log"
	"strconv"

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
	window.Resize(fyne.NewSize(400, 300))

	// Settings area title
	settingsHeaderText := widget.NewLabel("Settings")
	settingsHeaderRow := container.New(
		layout.NewVBoxLayout(),
		settingsHeaderText,
	)

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
	hours := make([]string, 25)
	for hour := 0; hour <= 24; hour++ {
		hours[hour] = strconv.Itoa(hour)
	}
	hourSelect := widget.NewSelect(hours, func(value string) {
		log.Println("Hour select set to", value)
	})
	minutes := make([]string, 61)
	for minute := 0; minute <= 60; minute++ {
		minutes[minute] = strconv.Itoa(minute)
	}
	minuteSelect := widget.NewSelect(minutes, func(value string) {
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
		settingsHeaderRow,
		content,
		bottomGrid,
	))
	window.ShowAndRun()
}
