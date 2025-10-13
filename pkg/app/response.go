package app

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/justatempa/runfast-go/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Time      string      `json:"time"`
	RequestId string      `json:"request_id"`
}

const (
	HttpRequestRawData       = "http_request_raw_data"
	HttpRequestStartUnixNano = "http_request_start_unix_nano"
)

type Gin struct {
	C *gin.Context
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	msg := GetMsg(errCode)
	g.ResponseWithMsgAndCode(httpCode, errCode, msg, data)
}

func (g *Gin) ResponseSuccess(data interface{}) {
	g.Response(http.StatusOK, Success, data)
}

func (g *Gin) ResponseErrorMessage(msg string, data interface{}) {
	g.ResponseWithMsgAndCode(http.StatusOK, Error, msg, data)
}

func (g *Gin) ResponseError(data interface{}) {
	g.Response(http.StatusOK, Error, data)
}

func (g *Gin) ResponseWithMsgAndCode(httpCode, errCode int, msg string, data interface{}) {
	if errCode != Success {
		logger.Infof("Response api_resp_bad ErrCode: %d Msg: %s", errCode, msg)
	}
	var response = Response{
		Status:    errCode,
		Message:   msg,
		Data:      data,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		RequestId: "",
	}
	logMsg := fmt.Sprintf("%s接口access_log", g.C.Request.RequestURI)
	logger.Infof(logMsg)
	g.C.JSON(httpCode, response)
	return
}

func GetRequestId() (requestId string) {
	value := uuid.Must(uuid.New(), nil).String()
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

