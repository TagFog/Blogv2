package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
const (
	Success = 0
	Error = -1
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int ,data any ,msg string,c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg: msg,
	})
}

// 响应封装
func Ok(data any,msg string,c *gin.Context){
	Result(Success, data, msg, c)
}
func OkWithData(data any,msg string,c *gin.Context){
	Result(Success, data, "成功", c)
}
func OkWithMessage(data any,msg string,c *gin.Context){
	Result(Success, map[string]any{}, msg, c)
}

func Fail(data any,msg string,c *gin.Context){
	Result(Error, data, msg, c)
}
func FailWithMessage(msg string,c *gin.Context){
	Result(Error, map[string]any{}, msg, c)
}
func FailWithCode(code ErrorCode,c *gin.Context){
	msg,ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
	}
	Result(Error, map[string]any{}, "未知错误", c)
}