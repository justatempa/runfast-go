package app

const (
	Success = iota
	Error
)

var MsgFlags = map[int]string{
	Success: "success",
	Error:   "fail",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
