package disk_space

import (
	"disk-space-monitor"
	"fmt"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLowSpaceDrives(t *testing.T) {
	parts, _ := disk.Partitions(true)
	lowSpaceDrives := disk_space_monitor.getLowDiskSpaceDrives(parts, 90.0)

	fmt.Printf("lowSpaceDrives: %+v\n", lowSpaceDrives)
	assert.Equal(t, 1, len(lowSpaceDrives))
}
