package controllers

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pala/config"
	"github.com/daiguadaidai/poow/pala/server/task"
	"github.com/daiguadaidai/poow/pala/views/form"
	"github.com/daiguadaidai/poow/utils"
)

type TaskController struct {
	SC *config.ServerConfig
}

func NewTaskController() *TaskController {
	return &TaskController{
		SC: config.GetServerConfig(),
	}
}

// 开始执行命令, 将需要执行的任务放入队列中
func (this *TaskController) Start(form *form.TaskStartForm) {
	t := &task.Task{
		Program:  form.Program,
		TaskUUID: form.TaskUUID,
		Params:   form.Params,
		SC:       this.SC,
	}
	task.AddTask(t)
}

// 删除命令
func (this *TaskController) RemoveCommand(program string) error {
	if err := utils.RemoveFile(this.SC.GetProgramFilePath(program)); err != nil {
		return fmt.Errorf("删除指定程序文件失败. %v", err)
	}
	return nil
}

// 停止任务
func (this *TaskController) KillTask(uuid string) error {
	pid, ok := task.GetTask(uuid)
	if !ok {
		seelog.Warnf("没有获取到任务, 任务已经停止. %s", uuid)
		return nil
	}
	if err := utils.KillProcess(pid.(int)); err != nil {
		return fmt.Errorf("停止任务出错 pid: %d. uuid: %s", pid, uuid)
	}
	task.DestroyTask(uuid)

	return nil
}

// 查看文件后几个字节
func (this *TaskController) TailFile(form *form.TailFileForm) (*utils.TailData, error) {
	return utils.TailFile(form.Path, form.Start, form.Size)
}
