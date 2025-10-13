package midderware

import (
	"net/http"
	"github.com/justatempa/runfast-go/conf"
	"github.com/justatempa/runfast-go/pkg/app"

	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	var (
		app = app.Gin{C: c}
	)

	auth := true

	if auth == false {
		status := 10000
		app.ResponseWithMsgAndCode(http.StatusOK, status, conf.GetAccountErr(status), nil)
		c.Abort()
		return
	}

	c.Next()
}

