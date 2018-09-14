package 移位操作

import (
	"testing"
	"fmt"
)

func TestA(t *testing.T)  {
	a()
}

func a() {
	var a = 1 << 5
	fmt.Println(a)
}