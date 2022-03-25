package controller

type MStatus int

const (
	CodeSuccess         MStatus = 1001
	CodeInvaildParams   MStatus = 1002
	CodeUserExisted     MStatus = 1003
	CodeUserNotExisted  MStatus = 1004
	CodeInvaildPassword MStatus = 1005
	CodeSeverBase       MStatus = 1006
	CodeCertifiedFailed MStatus = 1007
)

var msgStrings = map[MStatus]string{
	CodeSuccess:         "success",
	CodeInvaildParams:   "请求参数错误",
	CodeUserExisted:     "用户已存在",
	CodeUserNotExisted:  "用户不存在",
	CodeInvaildPassword: "用户名或密码错误",
	CodeSeverBase:       "服务繁忙",
	CodeCertifiedFailed: "密码错误",
}

func (s MStatus) msg() string {
	if str, ok := msgStrings[s]; ok {
		return str
	}
	return ""
}
