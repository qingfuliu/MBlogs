package controller

type MStatus int

const (
	CodeSuccess MStatus = 1000 + iota
	CodeInvaildParams
	CodeUserExisted
	CodeUserNotExisted
	CodeInvaildPassword
	CodeSeverBase
	CodeCertifiedFailed
)

var MsgStrings = map[MStatus]string{
	CodeSuccess:         "success",
	CodeInvaildParams:   "请求参数错误",
	CodeUserExisted:     "用户已存在",
	CodeUserNotExisted:  "用户不存在",
	CodeInvaildPassword: "用户名或密码错误",
	CodeSeverBase:       "服务繁忙",
	CodeCertifiedFailed: "密码错误",
}

func (s MStatus) msg() string {
	if str, ok := MsgStrings[s]; ok {
		return str
	}
	return ""
}
