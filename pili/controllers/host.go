package controllers

import (
	"github.com/daiguadaidai/poow/pili/config"
	"github.com/daiguadaidai/poow/pili/dao"
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
