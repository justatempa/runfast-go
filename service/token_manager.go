package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/justatempa/runfast-go/model"
	"github.com/justatempa/runfast-go/pkg/database"
	"gorm.io/gorm"
)

// TokenManager 负责管理JWT token
type TokenManager struct {
	adminToken string
	db         *gorm.DB
}

// NewTokenManager 创建新的TokenManager实例
func NewTokenManager(adminToken string) *TokenManager {
	return &TokenManager{
		adminToken: adminToken,
		db:         database.GetDB(),
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
func (tm *TokenManager) GenerateUserToken(tokenName string) (string, error) {
	expire := time.Now().Add(time.Hour * 24)
	claims := jwt.MapClaims{
		"user": tokenName,
		"exp":  expire.Unix(), // Token有效期24小时
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tm.adminToken))
	if err != nil {
		return "", err
	}

	// 保存token到数据库
	userToken := model.AdminToken{
		Name:      tokenName,
		Token:     tokenString,
		Role:      "user",
		ExpiredAt: &expire,
	}
	result := tm.db.Create(&userToken)
	if result.Error != nil {
		return "", result.Error
	}

	return tokenString, nil
}

// ValidateToken 验证Token
func (tm *TokenManager) ValidateToken(tokenString string) (string, error) {
	// 从数据库查找token
	var adminToken model.AdminToken
	result := tm.db.Where("token = ?", tokenString).First(&adminToken)
	if result.Error != nil {
		// 检查是否为Admin Token
		if tokenString == tm.adminToken {
			return "admin", nil
		}
		return "", errors.New("无效的token")
	}

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

	return adminToken.Name, nil
}

// IsAdminToken 检查是否为Admin Token
func (tm *TokenManager) IsAdminToken(tokenString string) bool {
	return tokenString == tm.adminToken
}

// RemoveToken 删除Token
func (tm *TokenManager) RemoveToken(tokenString string) error {
	result := tm.db.Where("token = ?", tokenString).Delete(&model.AdminToken{})
	return result.Error
}

// ListTokens 获取所有Token列表
func (tm *TokenManager) ListTokens() ([]model.AdminToken, error) {
	var tokens []model.AdminToken
	result := tm.db.Find(&tokens)
	return tokens, result.Error
}
