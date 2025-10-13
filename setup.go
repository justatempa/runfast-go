package main

import (
	"os"
	"path"
	"runtime"
	"time"

	"github.com/justatempa/runfast-go/conf"
	"github.com/justatempa/runfast-go/pkg/logger"
)

func Setup(chdir bool) {
	/*
		chdir=true 运行单元测试时，默认的working dir是单元测试文件路径，需要切换到项目根目录，否则无法找到配置文件
		chdir=false 从main入口运行时，会根据main文件或build出二进制的相对路径搜索配置文�?	*/
	if chdir {
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "..")
		err := os.Chdir(dir)
		if err != nil {
			panic("setup chdir error")
		}
	}
	// 手动指定时区
	os.Setenv("TZ", "Asia/Shanghai")
	var CSTFromTZI, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = CSTFromTZI

	conf.Setup()

	logger.Setup()

	logger.Info("项目启动完毕")
}
