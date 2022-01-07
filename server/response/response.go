package response

import (
	"monitor_system/errcode"
)

type H map[string]interface{}

func Response(code errcode.ErrCode) H {
	return H{"code": code, "msg": code.String()}
}

func ResponseWithData(code errcode.ErrCode, data interface{}) H {
	return H{"code": code, "msg": code.String(), "data": data}
}
