package main

import (
	"fmt"
	human "github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/disk"
	"math"
)

func getLowDiskSpaceDrives(diskPartitions []disk.PartitionStat, lowSpacePercentageThreshold float64) []string {
	formatter := "%-14s %7s %7s %7s %4s %s\n"
	fmt.Printf(formatter, "Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on")

	lowSpaceDrives := make([]string, 0)
	for _, p := range diskPartitions {
		device := p.Mountpoint
		stat, _ := disk.Usage(device)

		if stat.Total == 0 {
			continue
		}

		usedPercent := fmt.Sprintf("%2.f%%", stat.UsedPercent)
		fmt.Printf(
			formatter,
			stat.Fstype,
			human.Bytes(stat.Total),
			human.Bytes(stat.Used),
			human.Bytes(stat.Free),
			usedPercent,
			p.Mountpoint,
		)

		roundedUsedPercent := math.Round(stat.UsedPercent)
		if roundedUsedPercent >= lowSpacePercentageThreshold {
			lowSpaceDrives = append(lowSpaceDrives, device)
		} else {
			fmt.Printf("%v is at %v%%\n", device, roundedUsedPercent)
		}
	}
	return lowSpaceDrives
}

func main() {
	diskPartitions, _ := disk.Partitions(true)
	lowSpaceDrives := getLowDiskSpaceDrives(diskPartitions, 90.0)
	for key, drive := range lowSpaceDrives {
		fmt.Printf("%v: has low disk space (%v%% usage)\n", key, drive)
	}
}
