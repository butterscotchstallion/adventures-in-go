package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"go.uber.org/zap"
)

func InitUI(logger *zap.SugaredLogger) {
	logger.Debug("Initializing UI")
	go func() {
		window := new(app.Window)
		window.Option(app.Title(appNameWithVersion))
		window.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		err := run(window, logger)
		if err != nil {
			logger.Fatal(err)
		}

		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window, logger *zap.SugaredLogger) error {
	logger.Debug("Running app UI")
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define a large label with an appropriate text:
			title := material.Body1(theme, "Hello, I'm a Disk Space Monitor")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
