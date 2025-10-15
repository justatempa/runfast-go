package main

import (
	"fmt"
	"log"

	"github.com/justatempa/runfast-go/conf"
	"github.com/justatempa/runfast-go/model"
	"github.com/justatempa/runfast-go/pkg/database"
)

func main() {
	// 初始化配置
	conf.Setup()

	// 初始化数据库连接
	database.InitDB()

	// 获取数据库连接实例
	db := database.GetDB()

	// 自动迁移创建表
	err := db.AutoMigrate(&model.AdminToken{})
	if err != nil {
		log.Fatalf("创建表失败: %v", err)
	}

	fmt.Println("数据库表创建成功")

	// 插入示例数据
	adminToken := model.AdminToken{
		Token: "example_token_123456",
	}
	result := db.Create(&adminToken)
	if result.Error != nil {
		log.Fatalf("插入示例数据失败: %v", result.Error)
	}

	fmt.Println("示例数据插入成功")

	// 查询数据
	var token model.AdminToken
	result = db.Where("token = ?", "example_token_123456").First(&token)
	if result.Error != nil {
		log.Fatalf("查询数据失败: %v", result.Error)
	} else {
		fmt.Printf("查询到数据: ID=%d, Token=%s\n", token.ID, token.Token)
	}
}