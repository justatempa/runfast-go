package router

import ("github.com/justatempa/runfast-go/midderware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(midderware.Cors, gin.Recovery())

	api(r)
	return r
}

func api(r *gin.Engine) {
	r.GET("ms_demo", MsDemo)
}

