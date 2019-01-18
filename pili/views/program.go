package views

import (
	"fmt"
	"github.com/Unknwon/com"
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
	group.GET("", this.List)
	group.GET("/download/:name", this.Download)
	group.POST("/upload_create", this.UploadCreate)
	group.POST("/upload_edit", this.UploadEdit)
	group.POST("", this.Create)
	group.GET("/get/:id", this.GetByID)
}

type ProgramHandler struct{}

// 列表
func (this *ProgramHandler) List(c *gin.Context) {
	pg := form.ParsePaginator(c)
	list, err := controllers.NewProgramController().Query(pg)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}
	returnSuccess(c, list)
}

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

// 上传命令  创建命令
func (this *ProgramHandler) UploadCreate(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}
	data, err := controllers.NewProgramController().UploadCreateProgram(file, c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, data)
}

// 上传命令  编辑命令
func (this *ProgramHandler) UploadEdit(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}
	idStr := c.PostForm("id")
	if len(idStr) == 0 {
		returnError(c, http.StatusInternalServerError, fmt.Errorf("请数据命令ID"))
		return
	}
	id, err := com.StrTo(idStr).Int64()
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	data, err := controllers.NewProgramController().UploadEditProgram(id, file, c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, data)
}

// 创建程序
func (this *ProgramHandler) Create(c *gin.Context) {
	var f form.ProgramCreateForm
	if err := c.ShouldBind(&f); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	if err := controllers.NewProgramController().Create(&f); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, nil)
}

// 通过ID获取Program
func (this *ProgramHandler) GetByID(c *gin.Context) {
	idStr, err := form.GetParam(c, "id")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}
	id, err := com.StrTo(idStr).Int64()
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	data, err := controllers.NewProgramController().GetByID(id)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, data)
}
