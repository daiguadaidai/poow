package views

import (
	"github.com/daiguadaidai/poow/pili/controllers"
	"github.com/daiguadaidai/poow/pili/views/form"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	handler := new(ProgramHandler)
	AddHandlerV1("/pili/programs", handler) // 添加当前页面的uri路径之前都要添加上这个
}

// 注册route
func (this *ProgramHandler) RegisterV1(group *gin.RouterGroup) {
	group.GET("/download/:name", this.Download)
}

type ProgramHandler struct{}

// 下载命令
func (this *ProgramHandler) Download(c *gin.Context) {
	fileName, err := form.ParseDownloadProgramName(c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	path, err := controllers.NewProgramController().GetDownloadFilePath(fileName)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnFile(c, fileName, path)
}
