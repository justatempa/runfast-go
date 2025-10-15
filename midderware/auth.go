package midderware

import (
	"github.com/justatempa/runfast-go/service"
	"net/http"
	"strings"

	"github.com/justatempa/runfast-go/conf"
	"github.com/justatempa/runfast-go/pkg/app"

	"github.com/gin-gonic/gin"
)

// TokenManager 实例
var tokenManager *service.TokenManager

// InitTokenManager 初始化TokenManager
func InitTokenManager(adminToken string) {
	tokenManager = service.NewTokenManager(adminToken)
}

// GetTokenManager 获取TokenManager实例
func GetTokenManager() *service.TokenManager {
	return tokenManager
}

func Verify(c *gin.Context) {
	var (
		ginApp = app.Gin{C: c}
	)

	// 从Header中获取Authorization
	authHeader := c.GetHeader(conf.Authorization)
	if authHeader == "" {
		status := 10001 // 未提供Authorization header
		ginApp.ResponseWithMsgAndCode(http.StatusUnauthorized, status, conf.GetAccountErr(status), nil)
		c.Abort()
		return
	}

	// 提取Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		status := 10002 // Authorization header格式错误
		ginApp.ResponseWithMsgAndCode(http.StatusUnauthorized, status, conf.GetAccountErr(status), nil)
		c.Abort()
		return
	}

	// 验证Admin Token
	if tokenManager.IsAdminToken(tokenString) {
		c.Set("user_type", "admin")
		c.Next()
		return
	}

	// 验证普通用户Token
	userInfo, err := tokenManager.ValidateToken(tokenString)
	if err != nil {
		status := 10003 // 无效的token
		ginApp.ResponseWithMsgAndCode(http.StatusUnauthorized, status, conf.GetAccountErr(status), nil)
		c.Abort()
		return
	}

	// 将用户信息保存到上下文
	c.Set("user_type", "user")
	c.Set("user_info", userInfo)
	c.Next()
}
