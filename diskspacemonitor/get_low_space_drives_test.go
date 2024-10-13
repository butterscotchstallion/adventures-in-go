package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLowSpaceDrives(t *testing.T) {
	parts, _ := disk.Partitions(true)
	lowSpaceDrives := GetLowDiskSpaceDrives(parts, 90.0)

	fmt.Printf("lowSpaceDrives: %+v\n", lowSpaceDrives)
	assert.Equal(t, 1, len(lowSpaceDrives))
}
