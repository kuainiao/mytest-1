package _interface

import "fmt"

type Sms struct {
}

func NewSms() *Sms {
	return &Sms{}
}

func (s *Sms) Send() {
	fmt.Println("this is sms Send print")
}

func (s *Sms) Delete() {
	fmt.Println("this is sms Delete print")
}
