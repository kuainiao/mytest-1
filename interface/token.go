package _interface

import "fmt"

type Token struct {
}

func NewToken() *Token {
	return &Token{}
}

func (s *Token) Send() {
	fmt.Println("this is token Send print")
}

func (s *Token) Delete() {
	fmt.Println("this is token Delete print")
}
