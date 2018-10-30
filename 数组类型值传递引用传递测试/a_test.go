package 数组类型值传递引用传递测试

import (
	"testing"
	"fmt"
)

func TestA(t *testing.T) {
	a()
}

func a() {

	param := []string{"1111", "2222", "3333"}
	b(param)
	fmt.Println(param)

	c(&param)
	fmt.Println(param)
}
func c(p *[]string) {
	*p = append(*p, "5555")
}

func b(p []string) {
	p = append(p, "4444")
}
