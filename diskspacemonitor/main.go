package main

import (
	"fmt"
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/getlantern/systray"
	"github.com/go-co-op/gocron/v2"
	"github.com/shirou/gopsutil/v4/disk"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const appVersion string = "0.1.0"
const appName string = "Disk Space Monitor"

var appNameWithVersion = fmt.Sprintf("%v v%v", appName, appVersion)
var systrayIconSet = false

func CheckLowSpaceAndNotify(logger *zap.SugaredLogger) {
	// Get low space devices
	diskPartitions, _ := disk.Partitions(true)
	lowSpaceDrives := GetLowDiskSpaceDrives(logger, diskPartitions, 90.0)
	messages := make([]string, 0)
	for _, statInfo := range lowSpaceDrives {
		message := fmt.Sprintf("%v has low disk space (%v%% usage)\n", statInfo.MountPoint, statInfo.DiskUsagePercent)
		logger.Debug(message)
		messages = append(messages, message)
	}

	logger.Debugf("There are %v drives with low space", len(messages))

	// Show notifications, if any low space drives
	if len(messages) > 0 {
		logger.Debugf("Showing %v notifications", len(messages))
		for _, message := range messages {
			ShowNotification("Low disk space", message)
		}
	}
}

func getIconBytes(logger *zap.SugaredLogger) ([]byte, error) {
	iconFilePath := "disk_space_monitor_icon.ico"
	bytes, err := os.ReadFile(iconFilePath)
	if err != nil {
		logger.Fatalf("Can't load icon file: %v", err)
	}
	return bytes, nil
}

func onReady(logger *zap.SugaredLogger) {
	logger.Debug("Setting up systray properties")
	iconBytes, _ := getIconBytes(logger)
	systray.SetIcon(iconBytes)
	systray.SetTitle(appName)
	systray.SetTooltip("Monitoring disk space...")
	mQuit := systray.AddMenuItem("Quit", "Quit and stop monitoring")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	os.Exit(0)
}

func setSystemTrayIcon(logger *zap.SugaredLogger) {
	logger.Debug("Setting system tray icon")

	var onReadyCallback = func() {
		onReady(logger)
	}
	systray.Run(onReadyCallback, onExit)
}

func scheduleSpaceCheck(logger *zap.SugaredLogger) {
	logger.Debug("Scheduling space check")
	scheduler, _ := gocron.NewScheduler()
	ScheduleSpaceCheck(scheduler, logger)
}

func createLogger() (*zap.SugaredLogger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // Capitalize the log level names
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC timestamp format
		EncodeDuration: zapcore.SecondsDurationEncoder, // Duration in seconds
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Short caller (file and line)
	}

	// Set up lumberjack as a logger:
	logger := &lumberjack.Logger{
		Filename:   "./logs/dsm.log", // Or any other path
		MaxSize:    500,              // MB; after this size, a new log file is created
		MaxBackups: 3,                // Number of backups to keep
		MaxAge:     28,               // Days
		Compress:   true,             // Compress the backups using gzip
	}

	writeSyncer := zapcore.AddSync(logger)

	// Set up zap logger configuration:
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Using JSON encoder, but you can choose another
		writeSyncer,
		zapcore.DebugLevel,
	)

	loggerZap := zap.New(core)
	sugar := loggerZap.Sugar()
	defer func(sugar *zap.SugaredLogger) {
		err := sugar.Sync()
		if err != nil {
			return
		}
	}(sugar)

	return sugar, nil
}

func initUI(logger *zap.SugaredLogger) {
	logger.Debug("Initializing UI")
	go func() {
		window := new(app.Window)
		err := run(window, logger)
		if err != nil {
			logger.Fatal(err)
		}
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
			title := material.H1(theme, "Hello, I'm a Disk Space Monitor")

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

func main() {
	sugar, err := createLogger()
	if err != nil {
		sugar.Fatal(err)
	}

	sugar.Debugf("\n%v", appNameWithVersion)

	go setSystemTrayIcon(sugar)
	go scheduleSpaceCheck(sugar)
	initUI(sugar)
	CheckLowSpaceAndNotify(sugar)
}
