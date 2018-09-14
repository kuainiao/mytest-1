package _interface

import "testing"

func TestC(t *testing.T) {
	a3()
}

func a3() {
	var sms MyService= NewSms()
	sms.Send()

	var token MyService=NewToken()
	token.Send()
}

type MyService interface {
	Send()
	Delete()
}
