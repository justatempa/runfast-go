package service

import (
	"github.com/justatempa/runfast-go/model"
)

// TokenManager 实例指针
var tokenManager *TokenManager

// SetTokenManager 设置TokenManager实例
func SetTokenManager(tm *TokenManager) {
	tokenManager = tm
}

// GenerateToken 生成用户Token
func GenerateToken(tokenName string, expireMinus int) (string, error) {
	if tokenManager == nil {
		return "", nil
	}
	return tokenManager.GenerateUserToken(tokenName, expireMinus)
}

// ListTokens 获取所有Token列表
func ListTokens() ([]model.AdminToken, error) {
	if tokenManager == nil {
		return make([]model.AdminToken, 0), nil
	}
	return tokenManager.ListTokens()
}

// RemoveToken 删除Token
func RemoveToken(tokenString string) error {
	if tokenManager != nil {
		return tokenManager.RemoveToken(tokenString)
	}
	return nil
}
