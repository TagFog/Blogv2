package ctype

import "encoding/json"

type SignStatus int

// 登录类型 --先这样后面再扩展
const (
	SignQQ SignStatus = 1 //qq
)

func (s SignStatus) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string{
	var str string
	switch s{
	case SignQQ:
		str = "QQ"
	}
	return str
}