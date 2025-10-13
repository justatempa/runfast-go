package router

import ("github.com/justatempa/runfast-go/pkg/app"
	"github.com/justatempa/runfast-go/service"

	"github.com/gin-gonic/gin"
)

func MsDemo(c *gin.Context) {
	var (
		app   = app.Gin{C: c}
		param = service.DemoParam{}
	)
	if err := c.ShouldBind(&param); err != nil {
		app.ResponseError(err.Error())
		return
	}
	resp, err := service.Demo(c, &param)
	if err != nil {
		app.ResponseError(err.Error())
		return
	}
	app.ResponseSuccess(resp)
}

