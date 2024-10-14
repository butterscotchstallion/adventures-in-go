package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/go-co-op/gocron/v2"
	"github.com/shirou/gopsutil/v4/disk"
	"go.uber.org/zap"
	"os"
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

func getIconBytes() ([]byte, error) {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	iconFilePath := "disk_space_monitor_icon.ico"
	bytes, err := os.ReadFile(iconFilePath)
	if err != nil {
		sugar.Fatalf("Can't load icon file: %v", err)
	}
	return bytes, nil
}

func onReady() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	iconBytes, _ := getIconBytes()
	systray.SetIcon(iconBytes)
	systray.SetTitle(appName)
	systray.SetTooltip("Monitoring disk space...")
	mQuit := systray.AddMenuItem("Quit", "Quit and stop monitoring")
	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(iconBytes)

	sugar.Debug("Systray icon set")
}

func onExit() {
	os.Exit(0)
}

func setSystemTrayIcon() {
	systray.Run(onReady, onExit)
}

func main() {
	setSystemTrayIcon()

	scheduler, _ := gocron.NewScheduler()
	ScheduleSpaceCheck(scheduler)
}
