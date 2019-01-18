package views

import (
	"github.com/daiguadaidai/poow/pili/controllers"
	"github.com/daiguadaidai/poow/pili/views/form"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	handler := new(HostHandler)
	AddHandlerV1("/pili/hosts", handler) // 添加当前页面的uri路径之前都要添加上这个
}

// 注册route
func (this *HostHandler) RegisterV1(group *gin.RouterGroup) {
	group.GET("/heartbeat/:host", this.Heartbeat)
	group.GET("/selector", this.ForSelector)
}

type HostHandler struct{}

// 获取host为前端selector
func (this *HostHandler) ForSelector(c *gin.Context) {
	hosts, err := controllers.NewHostController().QueryForSelector()
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
	}
	returnSuccess(c, hosts)
}

// 下载命令
func (this *HostHandler) Heartbeat(c *gin.Context) {
	h, err := form.GetParam(c, "host")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	if err := controllers.NewHostController().UpdateIsValidByHost(h, true); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, nil)
}
