package strings测试

import (
	"testing"
	"strings"
	"fmt"
)

func TestA(t *testing.T) {
	a()
}

func a() {
	a := strings.Index("qqqwer", "qqq")
	fmt.Println(a)
}
