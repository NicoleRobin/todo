package e

var msgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",
}

func GetMsg(code int) string {
	if msg, ok := msgFlags[code]; ok {
		return msg
	}
	return msgFlags[ERROR]
}
