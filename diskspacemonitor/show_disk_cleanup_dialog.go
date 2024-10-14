package main

import (
	"fmt"
	"os/exec"
)

func ShowDiskCleanupDialog(mountPoint string) {
	// cleanmgr.exe /d c: sageset:1
	cmd := exec.Command(fmt.Sprintf("cleanmgr.exe /d %v", mountPoint))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
