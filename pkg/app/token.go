package app

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenManager 负责管理JWT token
type TokenManager struct {
	adminToken string
	tokens     map[string]string // token -> user info
}

// NewTokenManager 创建新的TokenManager实例
func NewTokenManager(adminToken string) *TokenManager {
	return &TokenManager{
		adminToken: adminToken,
		tokens:     make(map[string]string),
	}
}

// GenerateAdminToken 生成Admin Token
func (tm *TokenManager) GenerateAdminToken() (string, error) {
	token := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, token); err != nil {
		return "", err
	}

	return hex.EncodeToString(token), nil
}

// GenerateUserToken 生成用户Token
func (tm *TokenManager) GenerateUserToken(userInfo string) (string, error) {
	claims := jwt.MapClaims{
		"user": userInfo,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token有效期24小时
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tm.adminToken))
	if err != nil {
		return "", err
	}

	// 保存token和用户信息的映射
	tm.tokens[tokenString] = userInfo

	return tokenString, nil
}

// ValidateToken 验证Token
func (tm *TokenManager) ValidateToken(tokenString string) (string, error) {
	// 首先检查是否为有效的用户Token
	if userInfo, exists := tm.tokens[tokenString]; exists {
		// 验证JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(tm.adminToken), nil
		})

		if err != nil {
			return "", err
		}

		if !token.Valid {
			return "", errors.New("无效的token")
		}

		return userInfo, nil
	}

	return "", errors.New("无效的token")
}

// IsAdminToken 检查是否为Admin Token
func (tm *TokenManager) IsAdminToken(tokenString string) bool {
	return tokenString == tm.adminToken
}

// RemoveToken 删除Token
func (tm *TokenManager) RemoveToken(tokenString string) {
	delete(tm.tokens, tokenString)
}

// ListTokens 获取所有Token列表
func (tm *TokenManager) ListTokens() map[string]string {
	return tm.tokens
}
