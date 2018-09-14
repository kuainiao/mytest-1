package newInterface

import "fmt"

type SmsService interface {
	Send()
}

type Sms struct {
	
}

func NewSms() *Sms{
	return &Sms{}
}

func (s *Sms)Send()  {
	fmt.Println("this is sms send print")
}
