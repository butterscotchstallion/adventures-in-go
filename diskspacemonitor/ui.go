package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/zap"
)

func InitUI(logger *zap.SugaredLogger) {
	logger.Debug("Initializing UI")

	a := app.New()
	w := a.NewWindow(appNameWithVersion)
	w.Resize(fyne.NewSize(600, 400))

	hello := widget.NewLabel("Hello world!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
