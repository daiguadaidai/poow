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
	ListenHost string // 启动服务绑定的IP
	ListenPort int    // 启动服务绑定的端口

	ProgramPath       string // 命令存放的路径
	UploadProgramPath string // 上传命令临时使用目录

	PalaTaskStartURL string // 通知执行命令的URL
	PalsTaskTailURL  string
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
	return fmt.Sprintf(this.PalsTaskTailURL, host)
}
