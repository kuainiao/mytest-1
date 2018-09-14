package _interface

import (
	"testing"
	"fmt"
)

func TestB(t *testing.T) {
	a2()
}

func a2() {
	var i Humaner

	var a AAA0 = "wo"
	i = a
	i.sayHi()

}

type Humaner interface {
	sayHi()
}

type AAA0 string

func (a AAA0) sayHi() {
	fmt.Println("AAA0.sayHi=", a)
}
