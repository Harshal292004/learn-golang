package main

import (
	"fmt"
	"notifications/interfaces" // Adjust the import path based on your module name
)

func main() {
	// Example instances of your notification types
	directMessage := interfaces.DirectMessage{
		SenderUsername: "user123",
		MessageContent: "Hello!",
		PriorityLevel:  1,
		IsUrgent:       true,
	}

	groupMessage := interfaces.GroupMessage{
		GroupName:      "Dev Team",
		MessageContent: "Meeting at 3 PM.",
		PriorityLevel:  2,
	}

	systemAlert := interfaces.SystemAlert{
		AlertCode:      "ALERT001",
		MessageContent: "System Maintenance Scheduled.",
	}

	// Process each notification
	notifications := []interfaces.Notifications{directMessage, groupMessage, systemAlert}

	for _, n := range notifications {
		msg, importance := interfaces.ProcessNotification(n)
		fmt.Println(msg)
		fmt.Println("Importance Level:", importance)
	}
}
