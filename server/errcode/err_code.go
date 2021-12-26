package errcode

type ErrCode int

//go:generate stringer -type ErrCode -linecomment -output code_string.go
const (
	ERR_CODE_OK             ErrCode = 0 // 请求成功
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 无效参数
	ERR_CODE_FAILED         ErrCode = 2 // 请求失败
	ERR_CODE_NO_RIGHTS      ErrCode = 3 // 没有权限
)
