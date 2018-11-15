package server

import (
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pili/config"
	"github.com/daiguadaidai/poow/pili/middleware"
	"github.com/daiguadaidai/poow/pili/views"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// 启动http服务
func StartHttpServer(_wg *sync.WaitGroup) {
	defer _wg.Done()

	// 注册路由
	router := gin.Default()
	router.Use(middleware.Cors())
	views.Register(router)

	// 获取pala启动配置信息
	sc := config.GetServerConfig()
	s := &http.Server{
		Addr:           sc.Address(),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	seelog.Infof("Pili监听地址为: %v", sc.Address())
	err := s.ListenAndServe()
	if err != nil {
		seelog.Errorf("pili启动服务出错: %v", err)
	}
}
