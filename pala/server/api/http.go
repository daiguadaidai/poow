package api

import (
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pala/config"
	"github.com/daiguadaidai/poow/pala/middleware"
	"github.com/daiguadaidai/poow/pala/views"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// 启动http服务
func Start(_wg *sync.WaitGroup) {
	defer _wg.Done()

	// 注册路由
	router := gin.Default()
	router.Use(middleware.Cors())
	views.Register(router)

	// 获取pala启动配置信息
	sc := config.GetServerConfig()
	s := &http.Server{
		Addr:           sc.PalaAddress(),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	seelog.Infof("Pala监听地址为: %v", sc.PalaAddress())
	err := s.ListenAndServe()
	if err != nil {
		seelog.Errorf("pala启动服务出错: %v", err)
	}
}
