package midderware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Authorization,Token,project_id,Project-Id,appcode,Content-Type,Upgrade,Connection,X-Real-IP")
	c.Header("Access-Control-Max-Age", "86400")
	if c.Request.Method == "OPTIONS" {
		c.Writer.WriteHeader(http.StatusOK)
		c.Abort()
	} else {
		c.Next()
	}
}
