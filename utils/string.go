package utils

import (
	"fmt"
	"github.com/flynn-archive/go-shlex"
	"math/rand"
	"net"
	"time"
)

const (
	UUID_LEN      = 30
	UUID_TIME_LEN = 24
)

// 获取唯一自增ID
func GetUUID() string {
	t := time.Now()
	uuid := t.Format("20060102150405123456")
	currUUIDLen := len(uuid)
	for i := 0; i < UUID_TIME_LEN-currUUIDLen; i++ {
		uuid += "0"
	}
	randLen := 6
	if currUUIDLen > UUID_TIME_LEN {
		randLen = UUID_LEN - currUUIDLen
	}
	return fmt.Sprintf("%s%s", uuid, RandString(randLen))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 用掩码实现随机字符串
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, r.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// 获取日期字符串
func GetDateStr() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func GetArgs(s string) ([]string, error) {
	return shlex.Split(s)
}

// 获取第一个非非回环网卡IP
func GetIntranetIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("没有获取到可以用IP")
}
