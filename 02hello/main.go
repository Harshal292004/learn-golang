package main

import (
	"fmt"
	"reflect"
)

type rect struct {
	height int
	width  int
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Membership
	Name string
}

func (r rect) area() int {
	return r.width * r.height
}

func (u *User) assignLimit() { // Use pointer receiver
	if u.Type == "Premium" {
		u.MessageCharLimit = 100
	} else {
		u.MessageCharLimit = 30
	}
}

func main() {
	var r = rect{
		width:  5,
		height: 10,
	}
	mystruct := struct {
		brand string
		model string
	}{
		brand: "tesla",
		model: "model 3",
	}

	var u = User{
		Membership: Membership{
			Type:             "Premium",
			MessageCharLimit: 0,
		},
		Name: "Harshal",
	}

	u.assignLimit() // Now this will modify the original struct

	fmt.Printf("Limit assigned: %v\n", u.MessageCharLimit)

	typ := reflect.TypeOf(mystruct)

	fmt.Printf("Struct is %d bytes\n", typ.Size())

	fmt.Println(r.area())
}
