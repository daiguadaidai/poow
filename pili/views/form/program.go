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

type ProgramCreateForm struct {
	Title         string  `json:"title" form:"titile" binding:"required"`
	FileName      string  `json:"file_name" form:"file_name" binding:"required"`
	TmpFileName   string  `json:"tmp_file_name" form:"tmp_file_name" binding:"required"`
	Params        string  `json:"params" form:"params"`
	HaveDedicate  bool    `json:"have_dedicate" form:"have_dedicate"`
	DedicateHosts []int64 `json:"dedicate_hosts" form:"dedicate_hosts"`
}
