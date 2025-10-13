package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Setup 初始化日志
func Setup() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.TimeKey = "time"
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeDuration = zapcore.SecondsDurationEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder

	encoder := zapcore.NewJSONEncoder(config)

	// 获取配置
	logPath := getLogFilePath()
	errorLogPath := getErrorLogFilePath()

	// 创建日志目录
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		fmt.Printf("创建日志目录失败: %v\n", err)
	}

	// 普通日志
	normalWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100, // MB
		MaxBackups: 10,
		MaxAge:     30, // 天
		Compress:   true,
	})

	// 错误日志
	errorWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   errorLogPath,
		MaxSize:    100, // MB
		MaxBackups: 10,
		MaxAge:     30, // 天
		Compress:   true,
	})

	// 控制台输出
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	consoleWriter := zapcore.AddSync(os.Stdout)

	// 创建核心
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, normalWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		})),
		zapcore.NewCore(encoder, errorWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})),
		zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel),
	)

	// 创建Logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// getLogFilePath 获取日志文件路径
func getLogFilePath() string {
	return fmt.Sprintf("logs/%s.log", time.Now().Format("20060102"))
}

// getErrorLogFilePath 获取错误日志文件路径
func getErrorLogFilePath() string {
	return fmt.Sprintf("logs/error_%s.log", time.Now().Format("20060102"))
}

// Debug 输出Debug级别日志
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Debugf 输出Debug级别日志（支持格式化）
func Debugf(format string, args ...interface{}) {
	Logger.Debug(fmt.Sprintf(format, args...))
}

// Info 输出Info级别日志
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Infof 输出Info级别日志（支持格式化）
func Infof(format string, args ...interface{}) {
	Logger.Info(fmt.Sprintf(format, args...))
}

// Warn 输出Warn级别日志
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Warnf 输出Warn级别日志（支持格式化）
func Warnf(format string, args ...interface{}) {
	Logger.Warn(fmt.Sprintf(format, args...))
}

// Error 输出Error级别日志
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Errorf 输出Error级别日志（支持格式化）
func Errorf(format string, args ...interface{}) {
	Logger.Error(fmt.Sprintf(format, args...))
}

// Fatal 输出Fatal级别日志
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

// Fatalf 输出Fatal级别日志（支持格式化）
func Fatalf(format string, args ...interface{}) {
	Logger.Fatal(fmt.Sprintf(format, args...))
}

// WithFields 添加字段
func WithFields(fields map[string]interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return zapFields
}

// GetCallerInfo 获取调用者信息
func GetCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}