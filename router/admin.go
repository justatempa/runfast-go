package router

import (
	"errors"
	"net/http"

	"github.com/justatempa/runfast-go/pkg/app"
	"github.com/justatempa/runfast-go/service"

	"github.com/gin-gonic/gin"
)

// AdminLoginRequest 管理员登录请求参数
type AdminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var adminService *service.AdminService

func init() {
	adminService = service.NewAdminService()
}

// AdminLogin 管理员登录接口
func AdminLogin(c *gin.Context) {
	var (
		ginApp = app.Gin{C: c}
		req    AdminLoginRequest
	)

	// 绑定请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		ginApp.ResponseWithMsgAndCode(http.StatusBadRequest, app.Error, "请求参数错误", nil)
		return
	}

	// 执行登录
	token, err := adminService.Login(req.Username, req.Password)
	if err != nil {
		var appErr *app.AppError
		if errors.As(err, &appErr) {
			ginApp.ResponseWithMsgAndCode(http.StatusUnauthorized, appErr.Code, appErr.Message, nil)
		} else {
			ginApp.ResponseWithMsgAndCode(http.StatusUnauthorized, app.Error, "登录失败", nil)
		}
		return
	}

	// 登录成功，返回token
	ginApp.ResponseWithMsgAndCode(http.StatusOK, app.Success, "登录成功", gin.H{
		"token": token,
	})
}
