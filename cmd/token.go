package cmd

import (
	"fmt"
	"github.com/justatempa/runfast-go/service"
	"github.com/spf13/cobra"
	"os"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Token management",
	Long:  `Generate and manage admin token`,
}

var generateTokenCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate admin token",
	Long:  `Generate a new admin token and save it to conf/admin_token.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建TokenManager实例
		tokenManager := service.NewTokenManager("")

		// 生成新的Admin Token
		adminToken, err := tokenManager.GenerateAdminToken()
		if err != nil {
			fmt.Printf("生成Admin Token失败: %v\n", err)
			os.Exit(1)
		}

		// 将Admin Token保存到文件
		file, err := os.Create("conf/admin_token.txt")
		if err != nil {
			fmt.Printf("创建token文件失败: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.WriteString(adminToken)
		if err != nil {
			fmt.Printf("写入token文件失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Admin Token生成成功: %s\n", adminToken)
		fmt.Println("Token已保存到 conf/admin_token.txt")
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	tokenCmd.AddCommand(generateTokenCmd)
}
