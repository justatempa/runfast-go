package cmd

import (
	"fmt"

	"github.com/justatempa/runfast-go/conf"
	"github.com/justatempa/runfast-go/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "run web application",
	Long:  `runfast-go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web called")

		//---------------------------------------
		//启动web程序
		gin.SetMode(conf.GlobalConfig.Server.RunMode)
		routersInit := router.Router()
		fmt.Println("step http port : ", conf.GlobalConfig.Server.HttpPort)

		//http
		routersInit.Run(fmt.Sprintf(":%d", conf.GlobalConfig.Server.HttpPort))
	},
}
