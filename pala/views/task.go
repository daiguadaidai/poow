package views

import (
	"github.com/daiguadaidai/poow/pala/controllers"
	"github.com/daiguadaidai/poow/pala/views/form"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	handler := new(TaskHandler)
	AddHandlerV1("/pala/tasks", handler) // 添加当前页面的uri路径之前都要添加上这个
}

// 注册route
func (this *TaskHandler) RegisterV1(group *gin.RouterGroup) {
	group.GET("/start", this.Start)
	group.POST("/start", this.Start)
	group.GET("/remove/:program", this.RemoveProgram)
	group.GET("/kill/:uuid", this.KillTask)
	group.GET("/tail", this.Tail)
}

type TaskHandler struct{}

// 执行一个任务
func (this *TaskHandler) Start(c *gin.Context) {
	tsf, err := form.NewTaskStartForm(c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	// 开始任务
	controllers.NewTaskController().Start(tsf)

	returnSuccess(c, nil)
}

// 删除命令文件
func (this *TaskHandler) RemoveProgram(c *gin.Context) {
	// 获取需要删除的程序
	program, err := form.GetParam(c, "program")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	// 删除指定的程序
	if err := controllers.NewTaskController().RemoveCommand(program); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, nil)
}

// 停止一个命令
func (this *TaskHandler) KillTask(c *gin.Context) {
	// 获取任务uuid
	taskUUID, err := form.GetParam(c, "uuid")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	// 停止任务
	if err := controllers.NewTaskController().KillTask(taskUUID); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}
	returnSuccess(c, nil)
}

// 获取日志后几行
func (this *TaskHandler) Tail(c *gin.Context) {
	form, err := form.NewTailFileForm(c)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	tailData, err := controllers.NewTaskController().TailFile(form)
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}
	returnSuccess(c, tailData)
}
