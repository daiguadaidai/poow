package form

import (
	"fmt"
	"github.com/daiguadaidai/poow/pili/models"
	"github.com/daiguadaidai/poow/utils/types"
)

type TaskStartForm struct {
	Program       string `form:"program" json:"program" binding:"required"`
	Params        string `form:"params" json:"params"`
	PUUID         string `form:"puuid", json:"puuid"`
	NeedTaskUUID  bool   `form:"need_task_uuid" json:"need_task_uuid"`
	NeedUpdateAPI bool   `form:"need_update_api" json:"need_update_api"`
	NeedReadAPI   bool   `form:"need_read_api" json:"need_read_api"`
	NeedTaskPUUID bool   `form:"need_task_puuid" json:"need_task_puuid"`
	UseParentHost bool   `form:"use_parent_host" json:"use_parent_host"`
}

// 获取post参数
func (this TaskStartForm) GetPostData(uuid string, addr string) map[string]interface{} {
	data := make(map[string]interface{})
	params := ""
	if this.NeedTaskUUID {
		params += fmt.Sprintf("--task-uuid=%s", uuid)
	}
	if this.NeedUpdateAPI {
		api := fmt.Sprintf("http://%s/api/v1/pili/tasks", addr)
		params += fmt.Sprintf(" --update-api=%#v", api)
	}
	if this.NeedReadAPI {
		api := fmt.Sprintf("http://%s/api/v1/pili/tasks/get", addr)
		params += fmt.Sprintf(" --read-api=%#v", api)
	}
	if this.NeedTaskPUUID {
		params += fmt.Sprintf(" --parent-task-uuid=%s", this.PUUID)
	}

	data["program"] = this.Program
	data["task_uuid"] = uuid
	data["puuid"] = this.PUUID
	data["params"] = fmt.Sprintf("%s %s", this.Params, params)

	return data
}

// update form
type UpdateTaskForm struct {
	TaskUUID   string `form:"task_uuid" json:"task_uuid" binding:"required"`
	Host       string `form:"host" json:"host"`
	FileName   string `form:"file_name" json:"file_name"`
	Params     string `form:"params" json:"params"`
	LogPath    string `form:"log_path" json:"log_path"`
	NotifyInfo string `form:"notify_info" json:"notify_info"`
	RealInfo   string `form:"real_info" json:"real_info"`
	Status     int64  `form:"status" json:"status"`
	PUUID      string `form:"puuid", json:"puuid"`
}

// 通过一个 form 新建任务
func (this *UpdateTaskForm) NewTask() *models.Task {
	return &models.Task{
		TaskUUID:   types.NewNullString(this.TaskUUID, false),
		Host:       types.NewNullString(this.Host, false),
		FileName:   types.NewNullString(this.FileName, false),
		Params:     types.NewNullString(this.Params, false),
		LogPath:    types.NewNullString(this.LogPath, false),
		NotifyInfo: types.NewNullString(this.NotifyInfo, false),
		RealInfo:   types.NewNullString(this.RealInfo, false),
		Status:     types.NewNullInt64(this.Status, false),
		PUUID:      types.NewNullString(this.PUUID, false),
	}
}

type TailForm struct {
	TaskUUID string `form:"task_uuid" json:"task_uuid" binding:"required"`
	Size     int64  `form:"size" json:"size"`
	Start    int64  `form:"start" json:"start"`
}

type GetTaskForm struct {
	TaskUUID string `form:"task_uuid" json:"task_uuid" binding:"required"`
}
