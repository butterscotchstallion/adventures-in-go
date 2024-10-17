package main

import (
	"math"

	human "github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/disk"
	"go.uber.org/zap"
)

type deviceInfo struct {
	MountPoint       string
	DiskUsagePercent float64
	DiskTotal        string
}

func GetLowDiskSpaceDrives(logger *zap.SugaredLogger, diskPartitions []disk.PartitionStat, lowSpacePercentageThreshold float64) []deviceInfo {
	// formatter := "%-14s %7s %7s %7s %4s %s\n"
	// fmt.Printf(formatter, "Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on")

	lowSpaceDrives := make([]deviceInfo, 0)

	for _, partitionStat := range diskPartitions {
		device := partitionStat.Mountpoint
		stat, _ := disk.Usage(device)

		if stat.Total == 0 {
			continue
		}

		/*
			usedPercent := fmt.Sprintf("%2.f%%", stat.UsedPercent)
			logger.Debugf(
				formatter,
				stat.Fstype,
				human.Bytes(stat.Total),
				human.Bytes(stat.Used),
				human.Bytes(stat.Free),
				usedPercent,
				partitionStat.Mountpoint,
			)
		*/

		roundedUsedPercent := math.Round(stat.UsedPercent)
		if roundedUsedPercent >= lowSpacePercentageThreshold {
			info := deviceInfo{
				MountPoint:       partitionStat.Mountpoint,
				DiskUsagePercent: roundedUsedPercent,
				DiskTotal:        human.Bytes(stat.Total),
			}
			lowSpaceDrives = append(lowSpaceDrives, info)
		} else {
			logger.Debugf("%v %v%% usage", device, roundedUsedPercent)
		}
	}

	return lowSpaceDrives
}
