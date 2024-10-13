package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/disk"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	fmt.Println("Disk Space Monitor")

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
