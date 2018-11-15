package views

import (
	"github.com/daiguadaidai/poow/pili/controllers"
	"github.com/daiguadaidai/poow/pili/models"
	"github.com/daiguadaidai/poow/pili/views/form"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	handler := new(TaskHandler)
	AddHandlerV1("/pili/tasks", handler) // 添加当前页面的uri路径之前都要添加上这个
}

// 注册route
func (this *TaskHandler) RegisterV1(group *gin.RouterGroup) {
	group.PUT("", this.Update)
	group.GET("/start", this.Start)
	group.POST("/start", this.Start)
	group.GET("/success/:uuid", this.TaskSuccess)
	group.GET("/fail/:uuid", this.TaskFail)
	group.GET("/running/:uuid", this.TaskRunning)
	group.GET("/tail", this.TaskTail)
}

type TaskHandler struct{}

// 执行一个任务
func (this *TaskHandler) Start(c *gin.Context) {
	f, err := form.NewForm(c, &form.TaskStartForm{})
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	// 开始任务
	if err := controllers.NewTaskController().Start(f.(*form.TaskStartForm)); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, nil)
}

// 标记任务成功
func (this *TaskHandler) TaskSuccess(c *gin.Context) {
	uuid, err := form.GetParam(c, "uuid")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	if err := controllers.NewTaskController().UpdateStatusByUUID(uuid, models.TASK_STATUS_SUCCESS); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, nil)
}

// 标记任务失败
func (this *TaskHandler) TaskFail(c *gin.Context) {
	uuid, err := form.GetParam(c, "uuid")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	if err := controllers.NewTaskController().UpdateStatusByUUID(uuid, models.TASK_STATUS_FAIL); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, nil)
}

// 标记任务正在运行
func (this *TaskHandler) TaskRunning(c *gin.Context) {
	uuid, err := form.GetParam(c, "uuid")
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	if err := controllers.NewTaskController().UpdateStatusByUUID(uuid, models.TASK_STATUS_RUNNING); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, nil)
}

// 更新任务
func (this *TaskHandler) Update(c *gin.Context) {
	f, err := form.NewForm(c, &form.UpdateTaskForm{})
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	if err := controllers.NewTaskController().Update(f.(*form.UpdateTaskForm)); err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}
	returnSuccess(c, nil)

}

// 查看日志文件
func (this *TaskHandler) TaskTail(c *gin.Context) {
	f, err := form.NewForm(c, &form.TailForm{})
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	data, err := controllers.NewTaskController().TailFile(f.(*form.TailForm))
	if err != nil {
		returnError(c, http.StatusInternalServerError, err)
		return
	}

	returnSuccess(c, data)
}
