package app

import "fmt"

const (
	Success = iota
	Error
	AuthError
	LoginLimitError
	InvalidCredentials
	AdminTokenReadError
)

var MsgFlags = map[int]string{
	Success:             "success",
	Error:               "fail",
	AuthError:           "认证失败",
	LoginLimitError:     "登录次数超过限制，请稍后再试",
	InvalidCredentials:  "用户名或密码错误",
	AdminTokenReadError: "获取admin token失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}

// AppError 自定义应用错误类型
type AppError struct {
	Code    int
	Message string
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	return e.Message
}

// NewError 创建新的 AppError 实例
func NewError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// NewErrorf 创建带格式化消息的 AppError 实例
func NewErrorf(code int, format string, args ...interface{}) *AppError {
	return &AppError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// WrapError 包装现有错误
func WrapError(code int, message string, err error) *AppError {
	if err == nil {
		return NewError(code, message)
	}
	return &AppError{
		Code:    code,
		Message: fmt.Sprintf("%s: %v", message, err),
	}
}
