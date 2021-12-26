package errcode

import "fmt"

type ErrCode int

// 定义错误码
const (
	ERR_CODE_OK             ErrCode = 0 // 请求成功
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 无效参数
	ERR_CODE_FAILED         ErrCode = 2 // 请求失败
	ERR_CODE_NO_RIGHTS      ErrCode = 3 // 没有权限
)

// 定义错误码与描述信息的映射
var mapErrDesc = map[ErrCode]string{
	ERR_CODE_OK:             "请求成功",
	ERR_CODE_INVALID_PARAMS: "无效参数",
	ERR_CODE_FAILED:         "请求失败",
	ERR_CODE_NO_RIGHTS:      "没有权限",
}

// 根据错误码返回描述信息
func getDescription(errCode ErrCode) string {
	if desc, exist := mapErrDesc[errCode]; exist {
		return desc
	}

	return fmt.Sprintf("error code: %d", errCode)
}

func (e ErrCode) String() string {
	return getDescription(e)
}
