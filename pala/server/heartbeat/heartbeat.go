package heartbeat

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pala/config"
	"github.com/daiguadaidai/poow/utils"
	"sync"
	"time"
)

func Start(wg *sync.WaitGroup) {
	defer wg.Done()

	sc := config.GetServerConfig()
	ticker := time.NewTicker(time.Second * time.Duration(sc.HeartbeatInterval))

	for {
		select {
		case <-ticker.C:
			if err := heartBeat(sc); err != nil {
				seelog.Errorf("上报心跳失败. %v", err)
			}
		}
	}
}

// 执行heartbeat
func heartBeat(sc *config.ServerConfig) error {
	ip, err := utils.GetIntranetIp()
	if err != nil {
		return err
	}

	if _, err = utils.GetURL(sc.GetPiliHeartbeatURL(ip), ""); err != nil {
		return fmt.Errorf("%v %s", err, sc.GetPiliHeartbeatURL(ip))
	}

	return nil
}
