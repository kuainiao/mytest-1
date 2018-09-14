package 结构体转化map

import (
	"testing"
	"fmt"
)

func TestA(t *testing.T) {
	a()
}


func a() {
	a := &OauthAccountData{
		Sns:     "sina",
		Uid:     "iamUid",
		Guid:    1111,
		Appid:   "iamAppid",
		UnionId: "iamUnionid",
	}
	b := a.Serialize()
	fmt.Printf("%T\n", b)
	fmt.Printf("%+v\n", b)
}

type OauthAccountData struct {
	Sns     string
	Uid     string
	Guid    int64
	Appid   string
	UnionId string
}

func (t *OauthAccountData) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"Sns":     t.Sns,
		"Uid":     t.Uid,
		"Guid":    t.Guid,
		"Appid":   t.Appid,
		"UnionId": t.UnionId,
	}
}