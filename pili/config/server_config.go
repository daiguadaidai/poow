package config

import (
	"fmt"
	"github.com/daiguadaidai/poow/utils"
)

const (
	LISTEN_HOST = "0.0.0.0"
	LISTEN_PORT = 19528

	PROGRAM_PATH        = "./pili_programs"
	UPLOAD_PROGRAM_PATH = "./pili_upload_programs"

	PALA_TASK_START_URL = "http://%s:19529/api/v1/pala/tasks/start"
	PALA_TASK_TAIL_URL  = "http://%s:19529/api/v1/pala/tasks/tail"
)

var sc *ServerConfig

type ServerConfig struct {
	ListenHost string `toml:"listen_host"` // 启动服务绑定的IP
	ListenPort int    `toml:"listen_port"` // 启动服务绑定的端口

	ProgramPath       string `toml:"program_path"`        // 命令存放的路径
	UploadProgramPath string `toml:"upload_program_path"` // 上传命令临时使用目录

	PalaTaskStartURL string `toml:"pala_task_start_url"` // 通知执行命令的URL
	PalaTaskTailURL  string `toml:"pals_task_tail_url"`  //  获取pala日志信息URL
}

// 设置 piliStartconfig
func SetServerConfig(scf *ServerConfig) {
	sc = scf
}

func GetServerConfig() *ServerConfig {
	return sc
}

// 检测配置信息, 初始化一些需要的东西
func (this *ServerConfig) CheckAndStore() error {

	// 检测和创建命令存放目录(临时)
	if err := utils.CheckAndCreatePath(this.UploadProgramPath,
		"(临时)命令存放目录"); err != nil {
		return err
	}

	// 检测和创建命令存放目录(最终)
	if err := utils.CheckAndCreatePath(this.ProgramPath,
		"(最终)命令存放目录"); err != nil {
		return err
	}

	return nil
}

// 获取pala监听地址
func (this *ServerConfig) Address() string {
	return fmt.Sprintf("%v:%v", this.ListenHost, this.ListenPort)
}

// 上传文件临时存放路径
func (this *ServerConfig) UploadProgramFilePath(fn string) string {
	return fmt.Sprintf("%v/%v", this.UploadProgramPath, fn)
}

// 命令文件存放位置
func (this *ServerConfig) ProgramFilePath(fileName string) string {
	return fmt.Sprintf("%v/%v", this.ProgramPath, fileName)
}

func (this *ServerConfig) GetPalaTaskStartURL(host string) string {
	return fmt.Sprintf(this.PalaTaskStartURL, host)
}

func (this *ServerConfig) GetPalaTaskTailURL(host string) string {
	return fmt.Sprintf(this.PalaTaskTailURL, host)
}

// 补充默认值
func (this *ServerConfig) SupDefault() {
	// 启动服务绑定的IP
	if len(this.ListenHost) == 0 {
		this.ListenHost = LISTEN_HOST
	}
	// 启动服务绑定的端口
	if this.ListenPort < 1 {
		this.ListenPort = LISTEN_PORT
	}
	// 命令存放的路径
	if len(this.ProgramPath) == 0 {
		this.ProgramPath = PROGRAM_PATH
	}
	// 上传命令临时使用目录
	if len(this.UploadProgramPath) == 0 {
		this.UploadProgramPath = UPLOAD_PROGRAM_PATH
	}
	// 通知执行命令的URL
	if len(this.PalaTaskStartURL) == 0 {
		this.PalaTaskStartURL = PALA_TASK_START_URL
	}
	// 通知pala执行成功 URL
	if len(this.PalaTaskTailURL) == 0 {
		this.PalaTaskTailURL = PALA_TASK_TAIL_URL
	}
}
