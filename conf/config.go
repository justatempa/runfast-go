package conf

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

const (
	Authorization = "Authorization"
)

// AppConfig 应用配置
type AppConfig struct {
	AppName     string
	LogSavePath string
	LogSaveName string
	LogLevel    string
	TimeFormat  string
}

type ServerConfig struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Node         int
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
}

var GlobalConfig = &Config{}

// LoadConfig 加载配置文件
func LoadConfig() error {
	viper.SetConfigName("runfast-go")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		return fmt.Errorf("解析配置到结构体失败: %w", err)
	}

	// 转换超时时间为秒
	GlobalConfig.Server.ReadTimeout = time.Duration(viper.GetInt("server.ReadTimeout")) * time.Second
	GlobalConfig.Server.WriteTimeout = time.Duration(viper.GetInt("server.WriteTimeout")) * time.Second

	return nil
}

func Setup() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("配置初始化失败 %v", err)
	}
}

// GetConfig 获取配置
func GetConfig() *Config {
	return GlobalConfig
}

// GetDatabaseConfig 获取数据库配置
func GetDatabaseConfig() *DatabaseConfig {
	return &GlobalConfig.Database
}
