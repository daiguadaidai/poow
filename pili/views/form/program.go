package form

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// 解析下载程序名称
func ParseDownloadProgramName(c *gin.Context) (string, error) {
	fileName := c.Param("name")
	if strings.TrimSpace(fileName) == "" {
		return "", fmt.Errorf("没有指定下载的命令.")
	}
	return fileName, nil
}
