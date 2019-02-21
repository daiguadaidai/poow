package service

import (
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pala/config"
	"github.com/daiguadaidai/poow/pala/server/api"
	"github.com/daiguadaidai/poow/pala/server/heartbeat"
	"github.com/daiguadaidai/poow/pala/server/task"
	"sync"
)

func Start(cfg *config.Config) {
	defer seelog.Flush()
	logger, _ := seelog.LoggerFromConfigAsBytes([]byte(cfg.LC.Raw()))
	seelog.ReplaceLogger(logger)

	// 检测和创建指定和需要的目录
	err := cfg.SC.CheckAndStore()
	if err != nil {
		seelog.Errorf("检测启动配置文件错误: %v", err)
		return
	}

	config.SetServerConfig(cfg.SC)

	wg := new(sync.WaitGroup)

	// 启动palaserver
	wg.Add(1)
	go api.Start(wg)
	wg.Add(1)
	go task.Start(wg)
	wg.Add(1)
	go heartbeat.Start(wg)

	wg.Wait()
}
