package config

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/utils"
)

const (
	LISTEN_HOST = "0.0.0.0"
	LISTEN_PORT = 19529

	PROGRAM_PATH           = "./pala_programs"
	RUN_PROGRAM_LOG_PATH   = "./log"
	RUN_PROGRAM_PARALLER   = 8
	IS_LOG_DIR_PREFIX_DATE = true
	HEARTBEAT_INTERVAL     = 60

	PILI_SERVER = "localhost:19528"

	PILI_API_VERSTION         = "api/v1"
	PILI_DOWNLOAD_PROGRAM_URL = "http://%s/%s/pili/programs/download/%s"
	PILI_TASK_SUCCESS_URL     = "http://%s/%s/pili/tasks/success/%s"
	PILI_TASK_FAIL_URL        = "http://%s/%s/pili/tasks/fail/%s"
	PILI_HEARTBEAT_URL        = "http://%s/%s/pili/hosts/heartbeat/%s"
	PILI_TASK_UPDATE_URL      = "http://%s/%s/pili/tasks"
)

var sc *ServerConfig

type ServerConfig struct {
	ListenHost string // 启动服务绑定的IP
	ListenPort int    // 启动服务绑定的端口

	ProgramPath        string // 命令存放的路径
	RunProgramLogPath  string // 运行命令接收日志的输出位置
	RunProgramParaller int    // 运行命令并发数
	IsLogDirPrefixDate bool   // 日志的目录是否需要使用日期切割
	HeartbeatInterval  int    // 心跳检测间隔时间

	PiliServer string // 需要访问pili的host

	PiliAPIVersion         string
	PiliDownloadProgramURL string
	PiliTaskSuccessURL     string
	PiliTaskFailURL        string
	PiliHeartbeatURL       string
	PiliTaskUpdateURL      string
}

// 设置 palaStartconfig
func SetServerConfig(scf *ServerConfig) {
	sc = scf
}

func GetServerConfig() *ServerConfig {
	return sc
}

// 检测配置信息, 初始化一些需要的东西
func (this *ServerConfig) CheckAndStore() error {

	// 检测和创建命令存放目录
	if err := utils.CheckAndCreatePath(this.ProgramPath,
		"命令存放目录"); err != nil {
		return err
	}

	if err := utils.CheckAndCreatePath(this.RunProgramLogPath,
		"被执行命令的(错误)日志目录"); err != nil {
		return err
	}

	return nil
}

// 获取pili监听地址
func (this *ServerConfig) PiliAddress() string {
	return this.PiliServer
}

// 获取pala监听地址
func (this *ServerConfig) PalaAddress() string {
	return fmt.Sprintf("%v:%v", this.ListenHost, this.ListenPort)
}

// 获取命令相对路径
func (this *ServerConfig) GetProgramFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", this.ProgramPath, fileName)
}

func (this *ServerConfig) GetLogDir() string {
	if !this.IsLogDirPrefixDate {
		return this.RunProgramLogPath
	}

	dir := fmt.Sprintf("%v/%v", this.RunProgramLogPath, utils.GetDateStr())
	if err := utils.CheckAndCreatePath(dir, "被执行命令的(正常)日志目录"); err != nil {
		seelog.Warnf("创建不了输出日志目录. dir: %s. 使用默认的目录: %s. %v",
			dir, this.RunProgramLogPath, err)
		return this.RunProgramLogPath
	}
	return dir
}

// 获取日志路径
func (this *ServerConfig) GetLogPath(taskUUID string) string {
	return fmt.Sprintf("%s/%s.log", this.GetLogDir(), taskUUID)
}

// 获取pili下载命令url
func (this *ServerConfig) GetPiliDownloadProgramURL(command string) string {
	return fmt.Sprintf(this.PiliDownloadProgramURL, this.PiliServer, this.PiliAPIVersion, command)
}

// 获取pili任务成功url
func (this *ServerConfig) GetPiliTaskSuccessURL(taskUUID string) string {
	return fmt.Sprintf(this.PiliTaskSuccessURL, this.PiliServer, this.PiliAPIVersion, taskUUID)
}

// 获取pili任务失败url
func (this *ServerConfig) GetPiliTaskFailURL(taskUUID string) string {
	return fmt.Sprintf(this.PiliTaskFailURL, this.PiliServer, this.PiliAPIVersion, taskUUID)
}

// 心跳检测api
func (this *ServerConfig) GetPiliHeartbeatURL(host string) string {
	return fmt.Sprintf(this.PiliHeartbeatURL, this.PiliServer, this.PiliAPIVersion, host)
}

// 更新任务接口
func (this *ServerConfig) GetPiliTaskUpdateURL() string {
	return fmt.Sprintf(this.PiliTaskUpdateURL, this.PiliServer, this.PiliAPIVersion)
}
