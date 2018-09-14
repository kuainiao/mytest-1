package json转换结构体

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestA(t *testing.T)  {
	a()
}


func a() {
	body := `{"ok":true,"code":0,"error":"","data":
{"guid":"12345000141244","token":"04A84E4F104612B41CFC322102CECBC2","kickToken":"","username":"QiaQ","nickname":"","realNameStatus":1}}`
	var r Resp
	MustParseJSON(body, &r)
	fmt.Printf("%+v\n", r)
}

func MustParseJSON(body string, result interface{}) {
	err := json.Unmarshal([]byte(body), result)
	if err != nil {
		panic(err)
	}
}

type Resp struct {
	Ok   bool     `json:"ok"`
	Code int      `json:"code"`
	Err  string   `json:"error"`
	Data RespData `json:"data"`
}

type RespData struct {
	GUID           int64  `json:"guid,string"` // guid
	Token          string `json:"token"`       // token
	KickToken      string `json:"kickToken"`   // kicktoken
	UserName       string `json:"username"`    // 用户名(第三方登录用到)
	Nickname       string `json:"nickname"`
	RealnameStatus int8   `json:"realNameStatus"` // 实名认证状态
}
