package router

import (
	"io/ioutil"
	"log"

	"github.com/justatempa/runfast-go/midderware"
	"github.com/justatempa/runfast-go/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(midderware.Cors, gin.Recovery())

	// 初始化TokenManager
	initTokenManager()

	// 注册Token相关路由
	tokenAPI(r)

	api(r)
	return r
}

// initTokenManager 初始化TokenManager
func initTokenManager() {
	// 读取Admin Token
	adminTokenBytes, err := ioutil.ReadFile("conf/admin_token.txt")
	if err != nil {
		log.Fatalf("读取Admin Token失败: %v", err)
	}

	adminToken := string(adminTokenBytes)

	// 初始化中间件中的TokenManager
	midderware.InitTokenManager(adminToken)

	// 初始化服务中的TokenManager
	service.SetTokenManager(midderware.GetTokenManager())
}

func api(r *gin.Engine) {
	r.GET("ms_demo", MsDemo)

	r.Group("/api", midderware.Verify)
}

// tokenAPI 注册Token相关路由
func tokenAPI(r *gin.Engine) {
	// 需要Admin权限的路由组
	admin := r.Group("/admin")
	admin.Use(midderware.Verify)
	admin.GET("/token/generate", GenerateToken)
	admin.GET("/token/list", ListTokens)
	admin.GET("/token/remove", RemoveToken)
}
