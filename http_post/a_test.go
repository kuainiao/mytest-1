package http_post

import (
	"testing"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
		)

func TestA(t *testing.T) {
	a()
}

func a() {
	client := &http.Client{}
	req := `{}`
	req_new := bytes.NewBuffer([]byte(req))
	request, _ := http.NewRequest("POST", "http://127.0.0.1:8080/v2/api/login/limit/oauth", req_new)
	request.Header.Set("Content-type", "application/json")


	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}


