package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/go-co-op/gocron/v2"
	"github.com/shirou/gopsutil/v4/disk"
	"go.uber.org/zap"
)

const appName string = "Disk Space Monitor"

func CheckLowSpaceAndNotify() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	fmt.Println(appName)

	// Get low space devices
	diskPartitions, _ := disk.Partitions(true)
	lowSpaceDrives := GetLowDiskSpaceDrives(diskPartitions, 90.0)
	messages := make([]string, 0)
	for _, statInfo := range lowSpaceDrives {
		message := fmt.Sprintf("%v has low disk space (%v%% usage)\n", statInfo.MountPoint, statInfo.DiskUsagePercent)
		sugar.Info(message)
		messages = append(messages, message)
	}

	// Show notifications, if any low space drives
	if len(messages) > 0 {
		sugar.Debugf("Showing %v notifications", len(messages))
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
	var onReadyCallback = func() {
		onReady(logger)
	}
	systray.Run(onReadyCallback, onExit)
}

func scheduleSpaceCheck(logger *zap.SugaredLogger) {
	scheduler, _ := gocron.NewScheduler()
	ScheduleSpaceCheck(scheduler, logger)
}

func main() {
	sugar := zap.Must(zap.NewDevelopment()).Sugar()
	defer func(sugar *zap.SugaredLogger) {
		err := sugar.Sync()
		if err != nil {
			sugar.Fatal(err)
		}
	}(sugar)

	go scheduleSpaceCheck(sugar)
	setSystemTrayIcon(sugar)
}
