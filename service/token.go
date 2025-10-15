package service

import (
	"github.com/justatempa/runfast-go/pkg/app"
)

// TokenManager 实例指针
var tokenManager *app.TokenManager

// SetTokenManager 设置TokenManager实例
func SetTokenManager(tm *app.TokenManager) {
	tokenManager = tm
}

// GenerateToken 生成用户Token
func GenerateToken(userInfo string) (string, error) {
	if tokenManager == nil {
		return "", nil
	}
	return tokenManager.GenerateUserToken(userInfo)
}

// ListTokens 获取所有Token列表
func ListTokens() map[string]string {
	if tokenManager == nil {
		return make(map[string]string)
	}
	return tokenManager.ListTokens()
}

// RemoveToken 删除Token
func RemoveToken(tokenString string) {
	if tokenManager != nil {
		tokenManager.RemoveToken(tokenString)
	}
}
