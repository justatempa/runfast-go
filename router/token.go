package router

import (
	"net/http"

	"github.com/justatempa/runfast-go/pkg/app"
	"github.com/justatempa/runfast-go/service"

	"github.com/gin-gonic/gin"
)

// GenerateToken 生成用户Token接口
func GenerateToken(c *gin.Context) {
	var (
		ginApp = app.Gin{C: c}
	)

	// 从请求中获取用户信息
	userInfo := c.Query("user_info")
	if userInfo == "" {
		status := 10004 // 用户信息不能为空
		ginApp.ResponseWithMsgAndCode(http.StatusBadRequest, status, "用户信息不能为空", nil)
		return
	}

	// 生成Token
	token, err := service.GenerateToken(userInfo)
	if err != nil {
		status := 10005 // 生成Token失败
		ginApp.ResponseWithMsgAndCode(http.StatusInternalServerError, status, "生成Token失败", nil)
		return
	}

	ginApp.ResponseWithMsgAndCode(http.StatusOK, 0, "success", gin.H{
		"token": token,
	})
}

// ListTokens 获取Token列表接口
func ListTokens(c *gin.Context) {
	var (
		ginApp = app.Gin{C: c}
	)

	tokens := service.ListTokens()
	ginApp.ResponseWithMsgAndCode(http.StatusOK, 0, "success", tokens)
}

// RemoveToken 删除Token接口
func RemoveToken(c *gin.Context) {
	var (
		ginApp = app.Gin{C: c}
	)

	// 从请求中获取要删除的Token
	token := c.Query("token")
	if token == "" {
		status := 10006 // Token不能为空
		ginApp.ResponseWithMsgAndCode(http.StatusBadRequest, status, "Token不能为空", nil)
		return
	}

	// 删除Token
	service.RemoveToken(token)
	ginApp.ResponseWithMsgAndCode(http.StatusOK, 0, "success", nil)
}
