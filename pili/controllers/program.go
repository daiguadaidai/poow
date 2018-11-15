package controllers

import (
	"fmt"
	"github.com/daiguadaidai/poow/pili/config"
	"github.com/daiguadaidai/poow/utils"
)

type ProgramController struct {
	sc  *config.ServerConfig
	dbc *config.DBConfig
}

func NewProgramController() *ProgramController {
	return &ProgramController{
		sc:  config.GetServerConfig(),
		dbc: config.GetDBConfig(),
	}
}

// 获取需要下载文件的路径
func (this *ProgramController) GetDownloadFilePath(fileName string) (string, error) {
	// 判断命令是否存在
	fileNamePath := this.sc.ProgramFilePath(fileName)
	if exists, _ := utils.PathExists(fileNamePath); !exists {
		return "", fmt.Errorf("命令文件不存在 %v", fileNamePath)
	}

	return fileNamePath, nil
}
