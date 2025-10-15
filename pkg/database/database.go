package database

import (
	"fmt"
	"log"
	"time"

	"github.com/justatempa/runfast-go/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// 从配置文件获取数据库配置
	dbConfig := conf.GetDatabaseConfig()

	// 构建数据源名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	// 打开数据库连接
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("打开数据库连接失败: %v", err)
	}

	// 获取通用数据库对象 sql.DB 以设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取sql.DB失败: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(25)                 // 最大连接数
	sqlDB.SetMaxIdleConns(25)                 // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // 最大连接生命周期

	// 测试数据库连接
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	log.Println("数据库连接成功")
}

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return db
}
