package main

import (
	"os"

	"github.com/getlantern/systray"
	"go.uber.org/zap"
)

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

func SetSystemTrayIcon(logger *zap.SugaredLogger) {
	logger.Debug("Setting system tray icon")

	var onReadyCallback = func() {
		onReady(logger)
	}
	systray.Run(onReadyCallback, onExit)
}
