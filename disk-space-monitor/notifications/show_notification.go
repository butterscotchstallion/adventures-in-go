package notifications

import (
	"gopkg.in/toast.v1"
	"log"
)

func ShowNotification(title string, message string) {
	notification := toast.Notification{
		AppID:   "Disk Space Monitor",
		Title:   title,
		Message: message,
		//Icon:    "go.png", // This file must exist (remove this line if it doesn't)
		/*
			Actions: []toast.Action{

				{"protocol", "I'm a button", ""},
				{"protocol", "Me too!", ""},
			},
		*/
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
