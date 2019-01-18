package controllers

import (
	"github.com/daiguadaidai/poow/pili/config"
	"github.com/daiguadaidai/poow/pili/dao"
	"github.com/daiguadaidai/poow/pili/models"
)

type HostController struct {
	SC  *config.ServerConfig
	DBC *config.DBConfig
}

func NewHostController() *HostController {
	return &HostController{
		SC:  config.GetServerConfig(),
		DBC: config.GetDBConfig(),
	}
}

// 更新host是否可用
func (this *HostController) UpdateIsValidByHost(host string, isValid bool) error {
	return dao.NewHostDao().UpdateIsValidByHost(host, isValid)
}

// 获取选择器的hosts
func (this *HostController) QueryForSelector() ([]models.Host, error) {
	return dao.NewHostDao().QueryForSelector()
}
