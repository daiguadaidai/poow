package task

import (
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pala/config"
	"sync"
)

// 启动任务执行服务
func Start(wg *sync.WaitGroup) {
	defer wg.Done()

	sc := config.GetServerConfig()

	seelog.Info("启动任务服务")
	for i := 0; i < sc.RunProgramParaller; i++ {
		seelog.Infof("启动任务服务 %d 号 启动成功", i)
		wg.Add(1)
		go startTask(wg)
	}
}

// 启动task
func startTask(wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		go task.Start()
	}
}
