package main

import (
	"disk-space-monitor/disk_space"
	"fmt"
	"github.com/shirou/gopsutil/v4/disk"
)

func main() {
	diskPartitions, _ := disk.Partitions(true)
	lowSpaceDrives := disk_space.GetLowDiskSpaceDrives(diskPartitions, 90.0)
	for key, drive := range lowSpaceDrives {
		fmt.Printf("%v: has low disk space (%v%% usage)\n", key, drive)
	}
}
