package 时间加减

import (
	"testing"
	"time"
	"fmt"
)

func TestA(t *testing.T) {
	a()
}

func a() {
	aa:=1*time.Minute
	fmt.Printf("%T\n",aa)

	a := time.Now().Add(time.Duration(-1000 * time.Minute))
	fmt.Println("一千分钟前：", a)

	b := time.Now()
	fmt.Println("现在：", b)
}
