package controllers

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pili/config"
	"github.com/daiguadaidai/poow/pili/dao"
	"github.com/daiguadaidai/poow/pili/models"
	"github.com/daiguadaidai/poow/pili/views/form"
	"github.com/daiguadaidai/poow/utils"
	"github.com/daiguadaidai/poow/utils/types"
	"net/url"
)

type TaskController struct {
	SC  *config.ServerConfig
	DBC *config.DBConfig
}

func NewTaskController() *TaskController {
	return &TaskController{
		SC:  config.GetServerConfig(),
		DBC: config.GetDBConfig(),
	}
}

func (this *TaskController) Start(f *form.TaskStartForm) (*models.Task, error) {
	taskDao := dao.NewTaskDao()
	// 判断指定的父uuid是否存在
	var ptask *models.Task
	if f.NeedTaskPUUID {
		if len(f.PUUID) == 0 {
			return nil, fmt.Errorf("您没有指定puuid参数")
		}
		var err error
		if ptask, err = taskDao.GetByUUID([]string{"id", "host"}, f.PUUID); err != nil {
			if err.Error() != "record not found" {
				return nil, fmt.Errorf("通过指定的puuid: %s 找不到关联的父任务. %v", f.PUUID, err)
			}
			return nil, fmt.Errorf("通过指定的puuid: %s 获取符任务出错. %v", f.PUUID, err)
		}
	}

	// 获取执行的命令
	p, err := dao.NewProgramDao().GetByName(f.Program, []string{"id", "have_dedicate"})
	if err != nil {
		return nil, fmt.Errorf("获取命令出错: %s. %v", f.Program, err)
	}

	// 获取执行命令的机器
	var host string
	hostDao := dao.NewHostDao()
	if f.UseParentHost { // 使用父任务的host
		if ptask == nil {
			return nil, fmt.Errorf("您指定了使用父任务的机器执行当前命令. 但是您没有执行相关父任务")
		}
		host = ptask.Host.String
	} else { // 新选择host
		cols := []string{"hosts.id", "hosts.host"}
		h, err := hostDao.GetByProgramIDAndDedicate(p.ID.Int64, p.HaveDedicate, cols)
		if err != nil {
			return nil, fmt.Errorf("失去执行命令机器失败. %v", err)
		}
		host = h.Host.String
	}

	// POST 启动命令URL/参数
	this.SC.GetPalaTaskStartURL(host)
	uuid := utils.GetUUID()
	postData := f.GetPostData(uuid, this.SC.Address())

	// 创建任务
	task := &models.Task{
		TaskUUID:  types.NewNullString(uuid, false),
		ProgramId: p.ID,
		Host:      types.NewNullString(host, false),
		FileName:  types.NewNullString(postData["program"].(string), false),
		Params:    types.NewNullString(postData["params"].(string), false),
		PUUID:     types.NewNullString(f.PUUID, false),
		Status:    types.NewNullInt64(models.TASK_STATUS_RUNNING, false),
	}
	if err := taskDao.Create(task); err != nil {
		return nil, fmt.Errorf("创建任务失败. %v, %v. %v", postData["program"],
			postData["params"], err)
	}

	// 发post请求给pala进行启动任务
	if _, err := utils.PostURL(this.SC.GetPalaTaskStartURL(host), postData); err != nil {
		switch err.(type) {
		case *url.Error:
			hostDao.UpdateIsValidByHost(task.Host.String, false)
			return task, fmt.Errorf("执行任务机器不可用. %s. %v", task.Host.String, err)
		}
		return task, err
	}

	// host 正在运行该任务数 +1
	if err := hostDao.IncrTaskByHost(host); err != nil {
		seelog.Warnf("任务启动成功. 添加当前host(%v)任务数失败", host)
	}

	return task, nil
}

// 更新任务状态
func (this *TaskController) UpdateStatusByUUID(uuid string, status int) error {
	taskDao := dao.NewTaskDao()
	if err := taskDao.UpdateStatusByUUID(uuid, status); err != nil {
		return fmt.Errorf("设置任务<执行成功>失败. %v", err)
	}

	// 获取任务信息
	t, err := taskDao.GetByUUID([]string{"host"}, uuid)
	if err != nil {
		return fmt.Errorf("获取任务所在host失败(任务状态已经更新成功, host任务数将不会更新). %v", err)
	}

	// 更新host上运行的任务数
	switch status {
	case models.TASK_STATUS_FAIL, models.TASK_STATUS_SUCCESS:
		// 通过任务的host, 将host表中的任务运行数 -1
		if err := dao.NewHostDao().DecrTaskByHost(t.Host.String); err != nil {
			return fmt.Errorf("host任务数减1失败")
		}
	case models.TASK_STATUS_RUNNING:
		// 通过任务的host, 将host表中的任务运行数 -1
		if err := dao.NewHostDao().IncrTaskByHost(t.Host.String); err != nil {
			return fmt.Errorf("host任务数加1失败")
		}
	}

	return nil
}

// 更新task
func (this *TaskController) Update(f *form.UpdateTaskForm) error {
	return dao.NewTaskDao().UpdateByUUID(f.NewTask())
}

func (this *TaskController) TailFile(f *form.TailForm) (interface{}, error) {
	t, err := dao.NewTaskDao().GetByUUID([]string{"log_path"}, f.TaskUUID)
	if err != nil {
		return "", fmt.Errorf("获取日志路径失败. %v", err)
	}

	// 访问 pala tail 接口
	url := this.SC.GetPalaTaskTailURL(t.Host.String)
	queryMap := make(map[string]interface{})
	queryMap["size"] = f.Size
	queryMap["start"] = f.Start
	queryMap["path"] = t.LogPath.String
	query := utils.GetURLQuery(queryMap)

	return utils.GetURL(url, query)
}

// 查询通过program id task
func (this *TaskController) QueryByProgramID(pk int64, pg *utils.Paginator) ([]models.Task, error) {
	return dao.NewTaskDao().QueryByProgramID(pk, pg)
}

func (this *TaskController) GetByTaskUUID(f *form.GetTaskForm) (*models.Task, error) {
	return dao.NewTaskDao().GetByUUID([]string{"*"}, f.TaskUUID)
}
