package http带hosts访问

import (
	"testing"
	"net/http"
	"fmt"
)

func TestA(t *testing.T) {
	a()
}
func a() {
	ip := "10.80.78.146"
	geturl := "test0.web-channel-spring.shank.ifeng.com"

	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%v/heartbeat", geturl), nil)
	req.Host = ip

	resp, _ := http.DefaultClient.Do(req)
	fmt.Println(resp.Body)
}
