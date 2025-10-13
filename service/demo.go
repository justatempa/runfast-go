package service

import "github.com/gin-gonic/gin"

type DemoParam struct {
}
type DemoResult struct {
	Msg string `json:"msg"`
}

func Demo(r *gin.Context, param *DemoParam) (*DemoResult, error) {
	return &DemoResult{Msg: "hello world"}, nil
}
