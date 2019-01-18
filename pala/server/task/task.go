package task

import (
	"bufio"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pala/config"
	"github.com/daiguadaidai/poow/utils"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

var taskChan chan *Task
var runningTaskMap sync.Map

func init() {
	taskChan = make(chan *Task, 10000)
}

type Task struct {
	Program  string
	TaskUUID string
	Params   string
	LogPath  string
	LogFile  *os.File
	SC       *config.ServerConfig
}

// 添加任务
func AddTask(t *Task) {
	taskChan <- t
}

// 保存但前正在运行的 进程
func CacheTask(uuid string, pid int) {
	runningTaskMap.Store(uuid, pid)
}

// 销毁正在运行的任务
func DestroyTask(uuid string) {
	runningTaskMap.Delete(uuid)
}

// 获取任务
func GetTask(uuid string) (interface{}, bool) {
	return runningTaskMap.Load(uuid)
}

// 开始一个任务task
func (this *Task) Start() {
	this.InitLogFile()
	defer this.LogFile.Close()

	if err := this.UpdateLogPath(); err != nil {
		this.LogErrorf(err.Error())
	}

	// 检测命令是否存在
	if err := this.CheckAndDownloadProgram(); err != nil {
		this.TaskRunFail(err)
		return
	}

	// 检测命令是否有执行权限
	if err := this.CheckProgramIsExecutable(); err != nil {
		this.TaskRunFail(err)
		return
	}

	// 运行
	if err := this.Run(); err != nil {
		this.TaskRunFail(err)
		return
	}

	this.TaskRunSuccess()
}

// 运行命令
func (this *Task) Run() error {
	wg := new(sync.WaitGroup) // 再次创建一个并发控制器. 只提供运行命令中使用

	// 获取命令路经
	cmdPath := this.SC.GetProgramFilePath(this.Program)
	args, err := utils.GetArgs(this.Params)
	if err != nil {
		return fmt.Errorf("解析命令参数出错. params: %s. %v", this.Params, err)
	}
	// 创建命令执行器
	cmd := exec.Command(cmdPath, args...)
	// 主进程退出子进程也退出
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("获取输出管道错误. task uuid: %s. %s %s. %v",
			this.TaskUUID, cmdPath, this.Params, err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("获取<错误>输出管道错误. task uuid: %s. %s %s. %v",
			this.TaskUUID, cmdPath, this.Params, err)
	}

	// 开始执行命令
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("启动命令出错. task uuid: %s. %s %s. %v",
			this.TaskUUID, cmdPath, this.Params, err)
	}
	this.LogInfof("命令执行中. pid: %v. task uuid: %s. %s %s",
		cmd.Process.Pid, this.TaskUUID, cmdPath, this.Params)

	// 保存执行命令 pid
	CacheTask(this.TaskUUID, cmd.Process.Pid)
	defer DestroyTask(this.TaskUUID)

	// 记录命令的输出
	wg.Add(1)
	go this.LogOutput(wg, stdout)
	wg.Add(1)
	go this.LogOutput(wg, stderr)

	wg.Wait()
	// 等待结束
	if err := cmd.Wait(); err != nil {
		this.LogErrorf("Wait err: %v", err)
	}

	// 执行失败
	if !cmd.ProcessState.Success() {
		return fmt.Errorf("命令执行失败. pid: %v. task uuid: %s. %s %s",
			cmd.Process.Pid, this.TaskUUID, cmdPath, this.Params)
	}

	return nil
}

// 出入info日志
func (this *Task) LogOutput(_wg *sync.WaitGroup, _stdout io.ReadCloser) {
	defer _wg.Done()

	outputBuf := bufio.NewReader(_stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				this.LogErrorf("获取输出错误. %v", err)
				return
			}
		}
		this.Log(string(output))
	}
}

// 检测和下载命令
func (this *Task) CheckAndDownloadProgram() error {
	commandPath := fmt.Sprintf("%s/%s", this.SC.ProgramPath, this.Program)
	exists, err := utils.PathExists(commandPath)
	if err != nil {
		return err
	}
	if !exists {
		this.LogWarnf("命令不存在: %s", this.Program)
		if err1 := utils.DownloadFile(this.SC.GetPiliDownloadProgramURL(this.Program), commandPath); err1 != nil {
			return fmt.Errorf("%v. %s", err1, this.SC.GetPiliDownloadProgramURL(this.Program))
		}
		this.LogWarnf("命令下载成功: %s", this.Program)
	}

	return nil
}

func (this *Task) CheckProgramIsExecutable() error {
	commandPath := fmt.Sprintf("%s/%s", this.SC.ProgramPath, this.Program)
	executable, err := utils.FileIsExecutable(commandPath)
	if err != nil {
		return err
	}
	if !executable {
		this.LogWarnf("命令不可执行: %s", this.Program)
		if err1 := utils.ChmodFile(commandPath); err1 != nil {
			return err1
		}
		this.LogWarnf("命令可执行权限设置成功: %s", this.Program)
	}

	return nil
}

// 通知任务执行成功
func (this *Task) TaskRunSuccess() {
	if _, err := utils.GetURL(this.SC.GetPiliTaskSuccessURL(this.TaskUUID), ""); err != nil {
		this.LogErrorf("通知失败<任务完成>. UUID: %s, command: %s, params: %s. %v",
			this.TaskUUID, this.Program, this.Params, err)
		return
	}

	this.LogInfof("通知成功<任务完成>. UUID: %s, command: %s, params: %s",
		this.TaskUUID, this.Program, this.Params)
}

// 通知任务执行失败
func (this *Task) TaskRunFail(err error) {
	this.LogErrorf(err.Error())

	if _, err := utils.GetURL(this.SC.GetPiliTaskFailURL(this.TaskUUID), ""); err != nil {
		this.LogErrorf("通知失败<任务执行失败>. UUID: %s, command: %s, params: %s. %v",
			this.TaskUUID, this.Program, this.Params, err)
		return
	}

	this.LogInfof("通知成功<任务执行失败>. UUID: %s, command: %s, params: %s",
		this.TaskUUID, this.Program, this.Params)
}

// 初始化日志文件
func (this *Task) InitLogFile() {
	this.LogPath = this.SC.GetLogPath(this.TaskUUID)
	seelog.Infof("任务: %s. 命令: %s. 输出文件: %s", this.TaskUUID, this.Program, this.LogPath)

	var err error
	this.LogFile, err = os.OpenFile(this.LogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		seelog.Errorf("创建错误日志文件失败. task uuid: %s. logfile: %s",
			this.TaskUUID, this.LogPath)
		return
	}
}

func (this *Task) UpdateLogPath() error {
	data := make(map[string]string)
	data["task_uuid"] = this.TaskUUID
	data["log_path"] = this.LogPath

	if _, err := utils.PutURL(this.SC.GetPiliTaskUpdateURL(), data); err != nil {
		return fmt.Errorf("更新任务日志地址出错: %v", err)
	}

	return nil
}

func (this *Task) Log(info string) {
	if _, err := fmt.Fprintln(this.LogFile, info); err != nil {
		seelog.Errorf("写入自建日志出错. %v", err)
	}
}

// 记录输出信息
func (this *Task) LogInfof(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	seelog.Info(msg)
	this.Log(msg)
}

// 记录错误输出信息
func (this *Task) LogErrorf(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	seelog.Error(msg)
	this.Log(msg)
}

// 记录警告输出信息
func (this *Task) LogWarnf(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	seelog.Warn(msg)
	this.Log(msg)
}
