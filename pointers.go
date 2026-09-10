package main

import (
	"fmt"
)

func (e *Email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type Email struct {
	message     string
	fromAddress string
	toAddress   string
}



func (e Email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}
