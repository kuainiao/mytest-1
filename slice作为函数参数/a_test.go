package slice作为函数参数

import (
	"testing"
	"fmt"
)

func TestA(t *testing.T) {
	a()
}

func a() {
	aaa := []int{1, 2, 3, 4}
	t0(aaa)
	fmt.Printf("%+v\n", aaa)
	t1(&aaa)
	fmt.Printf("%+v\n", aaa)
	t2(&aaa)
	fmt.Printf("%+v\n", aaa)
}

func t0(a []int) {
	a = append(a, 5)
	a = append(a, 6)
}
func t1(a *[]int) {
	*a = append(*a, 7)
	*a = append(*a, 8)
}

func t2(a *[]int) {
	tmp := 1
	*a = append((*a)[:tmp], (*a)[tmp+1:]...)
}
