package _interface

import (
	"testing"
	"fmt"
)

func TestA(t *testing.T) {
	main()
}

func main() {
	var i Personer
	c := &a{}
	i = c
	fmt.Println(i.Send())

	b := &b{}
	fmt.Println(b.Send())
}

type Personer interface {
	Send() string
}

type a struct {
}
func (a *a) Send() string {
	return "this is a result"
}

type b struct {
}
func (b *b) Send() string {
	return "this is b result"
}
