package form

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type TaskStartForm struct {
	Program  string `form:"program" json:"program" binding:"required"`
	TaskUUID string `form:"task_uuid" json:"task_uuid" binding:"required"`
	Params   string `form:"params" json:"params"`
}

func NewTaskStartForm(c *gin.Context) (*TaskStartForm, error) {
	form := new(TaskStartForm)
	if err := c.ShouldBind(form); err != nil {
		return nil, err
	}

	return form, nil
}

// 获取参数
func GetParam(c *gin.Context, key string) (string, error) {
	v := c.Param(key)
	if strings.TrimSpace(v) == "" {
		return "", fmt.Errorf("必须输入参数 %s 值")
	}
	return v, nil
}

// 查看文件的form
type TailFileForm struct {
	Path string `form:"path" json:"path" binding:"required"`
	Size int64  `form:"size" json:"size"`
}

const (
	DEFAULT_TAIL_SIZE = 20480
)

func NewTailFileForm(c *gin.Context) (*TailFileForm, error) {
	form := new(TailFileForm)
	if err := c.ShouldBind(form); err != nil {
		return nil, err
	}
	if form.Size <= 0 {
		seelog.Warnf("需要查看的文件大小0. 将使用默认大小: %d", DEFAULT_TAIL_SIZE)
		form.Size = DEFAULT_TAIL_SIZE
	}
	// 如果查看的字节超出的文件大小则设置成文件大小
	fileSize, err := utils.FileSize(form.Path)
	if err != nil {
		return nil, fmt.Errorf("获取查看文件信息大小错误. %v", err)
	}
	if form.Size > fileSize {
		seelog.Warnf("指定查看的信息大小超出了文件大小. 将使用文件大小: %d", fileSize)
		form.Size = fileSize
	}

	return form, nil
}
