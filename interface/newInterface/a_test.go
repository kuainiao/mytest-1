package newInterface

import "testing"

func TestA(t *testing.T) {
	a1()
}

func a1() {
	var sms SmsService = NewSms()
	sms.Send()

	var token TokenService = NewToken()
	token.Send()
	token.Delete()
}
