package service

import (
	"os"
	"time"

	"github.com/justatempa/runfast-go/conf"
	"github.com/justatempa/runfast-go/pkg/app"
)

// AdminService 管理员服务
type AdminService struct {
	rateLimiter *app.RateLimiter
}

// NewAdminService 创建新的管理员服务
func NewAdminService() *AdminService {
	return &AdminService{
		rateLimiter: app.NewRateLimiter(3, time.Minute), // 1分钟内最多3次登录
	}
}

// Login 管理员登录
func (as *AdminService) Login(username, password string) (string, error) {
	// 限流检查
	if !as.rateLimiter.Allow(username) {
		return "", &app.AppError{Code: app.LoginLimitError, Message: app.GetMsg(app.LoginLimitError)}
	}

	// 获取管理员配置
	adminConfig := conf.GetAdminConfig()

	// 验证用户名
	if adminConfig.Username != username {
		return "", &app.AppError{Code: app.InvalidCredentials, Message: app.GetMsg(app.InvalidCredentials)}
	}

	// 验证密码
	// err := bcrypt.CompareHashAndPassword([]byte(adminConfig.Password), []byte(password))
	if !(adminConfig.Password == password) {
		return "", &app.AppError{Code: app.InvalidCredentials, Message: app.GetMsg(app.InvalidCredentials)}
	}

	// 读取admin token
	token, err := os.ReadFile("conf/admin_token.txt")
	if err != nil {
		return "", &app.AppError{Code: app.AdminTokenReadError, Message: app.GetMsg(app.AdminTokenReadError)}
	}

	return string(token), nil
}
