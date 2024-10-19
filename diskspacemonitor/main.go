package main

import (
	"fmt"

	"github.com/go-co-op/gocron/v2"
	"github.com/shirou/gopsutil/v4/disk"
	"go.uber.org/zap"
)

const appVersion string = "0.1.0"
const appName string = "Disk Space Monitor"

var appNameWithVersion = fmt.Sprintf("%v v%v", appName, appVersion)

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

func scheduleSpaceCheck(logger *zap.SugaredLogger) {
	logger.Debug("Scheduling space check")
	scheduler, _ := gocron.NewScheduler()
	ScheduleSpaceCheck(scheduler, logger)
}

func main() {
	sugar, err := CreateLogger()
	if err != nil {
		sugar.Fatal(err)
	}

	sugar.Debug("\n")
	sugar.Debugf("%v", appNameWithVersion)

	CheckLowSpaceAndNotify(sugar)
	go SetSystemTrayIcon(sugar)
	go scheduleSpaceCheck(sugar)
	InitUI(sugar)
}
