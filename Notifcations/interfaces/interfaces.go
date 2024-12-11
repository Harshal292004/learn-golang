// interfaces.go
package interfaces

import "fmt"

type Notifications interface {
	Importance() int
}

type DirectMessage struct {
	SenderUsername string
	MessageContent string
	PriorityLevel  int
	IsUrgent       bool
}

type GroupMessage struct {
	GroupName      string
	MessageContent string
	PriorityLevel  int
}

type SystemAlert struct {
	AlertCode      string
	MessageContent string
}

func (d DirectMessage) Importance() int {
	if d.IsUrgent {
		return 50
	}
	return d.PriorityLevel
}

func (g GroupMessage) Importance() int {
	return g.PriorityLevel
}

func (s SystemAlert) Importance() int {
	return 100
}

func ProcessNotification(n Notifications) (string, int) {
	switch v := n.(type) {
	case DirectMessage:
		msg := fmt.Sprintf("Is a %v notification and the priority level is %d", v, v.Importance())
		return msg, v.Importance()
	case GroupMessage:
		msg := fmt.Sprintf("Is a %v notification and the priority level is %d", v, v.Importance())
		return msg, v.Importance()
	case SystemAlert:
		msg := fmt.Sprintf("Is a %v notification and the importance level is %d", v, v.Importance())
		return msg, v.Importance()
	default:
		return "", 0
	}
}
