package main

import (
	"time"

	"gopkg.in/toast.v1"
)

func main() {
	var notification toast.Notification
	for {
		notification = toast.Notification{
			AppID:               "The 20 20 20 Rule",
			Title:               "Time to Rest Your Eyes",
			Message:             "It's been 20 minutes, look at something 20 feet away for 20 seconds.",
			Icon:                "",
			ActivationType:      "",
			ActivationArguments: "",
			Actions:             []toast.Action{{Type: "protocol", Label: "Dismiss", Arguments: ""}},
			Audio:               "",
			Loop:                false,
			Duration:            "long",
		}
		notification.Push()
		time.Sleep(time.Minute * 20)
	}
}
