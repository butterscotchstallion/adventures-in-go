package main

import (
	"log"

	"gopkg.in/toast.v1"
)

func ShowNotification(title string, message string) {
	notification := toast.Notification{
		AppID:   "Disk Space Monitor",
		Title:   title,
		Message: message,
		// This file must exist (remove this line if it doesn't)
		Icon: "E:\\projects\\adventures-in-go\\diskspacemonitor\\disk_space_monitor_icon.png",
		Actions: []toast.Action{
			{"protocol", "Go to Disk Space Clean Up", "http://example.com"},
		},
		Duration: "long",
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
