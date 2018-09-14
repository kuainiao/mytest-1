package 测试time类型

import (
	"testing"
	"time"
	"fmt"
)

func TestA(t *testing.T) {
	a()
	b()
}

func a() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	a := GetExpireDuration()
	fmt.Println(a)
	b := GetExpireTime()
	c := time.Unix(b, 0)
	fmt.Println(c.Format("2006-01-02 15:04:05"))

}


func b() {
	fmt.Println(time.Now().Unix())
	a := GetExpireDuration()
	fmt.Println(a)
	b := GetExpireTime()
	fmt.Println(b)

}


// redis 失效时间 返回类型 Duration
func GetExpireDuration() time.Duration {
	return 24 * time.Hour
}

// redis 失效时间 返回类型 int64 失效截止时间
func GetExpireTime() int64 {
	return time.Now().Add(GetExpireDuration()).Unix()
}
