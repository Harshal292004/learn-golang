package main

import "fmt"

type user struct {
	name   string
	number int
}
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 || mToSend.recipient.number == 0 {
		return false
	}
	return true

}
func mainn() {
	sender := user{"sgsg", 798}
	recipient := user{"dsjjf", 909}

	mToSend := messageToSend{"djsnfoids", sender, recipient}
	var boolean = canSendMessage(mToSend)
	fmt.Printf("%v is the new", boolean)
}
