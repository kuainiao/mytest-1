package newInterface

import "fmt"

type TokenService interface {
	Send()
	Delete()
}

type Token struct {
}

func NewToken() *Token {
	return &Token{}
}

func (s *Token) Send() {
	fmt.Println("this is token send print")
}
func (s *Token) Delete() {
	fmt.Println("this is token delete print")
}
