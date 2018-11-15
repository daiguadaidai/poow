package service

import (
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pili/config"
	"github.com/daiguadaidai/poow/pili/server"
	"sync"
)

func Start(sc *config.ServerConfig, dbc *config.DBConfig) {
	defer seelog.Flush()
	logger, _ := seelog.LoggerFromConfigAsBytes([]byte(config.LogDefautConfig()))
	seelog.ReplaceLogger(logger)

	// 检测和创建指定和需要的目录
	err := sc.CheckAndStore()
	if err != nil {
		seelog.Errorf("检测启动配置文件错误: %v", err)
		return
	}
	err = dbc.Check()
	if err != nil {
		seelog.Errorf("检测链接数据库配置错误: %v", err)
		return
	}

	config.SetServerConfig(sc) // 设置全局的http配置文件
	config.SetDBConfig(dbc)    // 设置全局的数据库配置文件

	wg := new(sync.WaitGroup)

	// 启动palaserver
	wg.Add(1)
	go server.StartHttpServer(wg)

	wg.Wait()
}
