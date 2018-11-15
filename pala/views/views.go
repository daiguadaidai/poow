package views

import (
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// 返回错误
func returnError(c *gin.Context, code int, err error) {
	resp := &response{Data: make([]interface{}, 0)}
	resp.Status = false
	resp.Message = err.Error()
	seelog.Error(resp.Message)
	c.JSON(code, resp)
}

// 成功并放回数据
func returnSuccess(c *gin.Context, obj interface{}) {
	resp := &response{Data: obj}
	resp.Status = true
	resp.Message = "success"
	c.JSON(http.StatusOK, resp)
}
