package conf

const (
	Success  = 0
	Error    = 1
	Download = 100
	ErrorMsg = "操作失败"
	Ok       = "ok"

	MissParamCode = 40002
	MissParamMsg  = "参数缺失"
	PassPortError = 4000
)

var accountErr = map[int]string{
	4000: "Passport服务异常，请稍后再试",
}

func GetAccountErr(code int) string {
	if msg, ok := accountErr[code]; ok {
		return msg
	}

	return "账号异常，请联系管理员"
}
