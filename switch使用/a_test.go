package switch使用

import (
	"testing"
	"fmt"
)

func TestA(t *testing.T)  {
	a()
}


func a() {
	switch {
	case true:
		fmt.Println("进入第一个")
		fallthrough
	case false:
		fmt.Println("第二个不应该进来")
	case true:
		fmt.Println("第三个不知道")
	}
}