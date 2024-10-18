package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"go.uber.org/zap"
)

type C = layout.Context
type D = layout.Dimensions

func InitUI(logger *zap.SugaredLogger) {
	logger.Debug("Initializing UI")
	go func() {
		window := new(app.Window)
		window.Option(app.Title(appNameWithVersion))
		window.Option(app.Size(unit.Dp(600), unit.Dp(400)))
		// err := run(window, logger)
		err := draw(window)
		if err != nil {
			logger.Fatal(err)
		}

		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	var ops op.Ops
	var saveButton widget.Clickable
	th := material.NewTheme()
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx C) D {
						// ONE: First define margins around the button using layout.Inset ...
						margins := layout.Inset{
							Top:    unit.Dp(15),
							Bottom: unit.Dp(0),
							Right:  unit.Dp(15),
							Left:   unit.Dp(15),
						}
						// TWO: ... then we lay out those margins ...
						return margins.Layout(gtx,
							// THREE: ... and finally within the margins, we define and lay out the button
							func(gtx C) D {
								btn := material.Button(th, &saveButton, "Save Settings")
								return btn.Layout(gtx)
							},
						)
					},
				),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)
			e.Frame(gtx.Ops)

		// this is sent when the application is closed
		case app.DestroyEvent:
			return e.Err
		}
	}
}
