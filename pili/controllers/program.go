package controllers

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/daiguadaidai/poow/pili/config"
	"github.com/daiguadaidai/poow/pili/dao"
	"github.com/daiguadaidai/poow/pili/models"
	"github.com/daiguadaidai/poow/pili/views/form"
	"github.com/daiguadaidai/poow/utils"
	"github.com/daiguadaidai/poow/utils/types"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type ProgramController struct {
	sc  *config.ServerConfig
	dbc *config.DBConfig
}

func NewProgramController() *ProgramController {
	return &ProgramController{
		sc:  config.GetServerConfig(),
		dbc: config.GetDBConfig(),
	}
}

func (this *ProgramController) Query(pg *utils.Paginator) ([]models.Program, error) {
	return dao.NewProgramDao().Query(pg)
}

// 获取需要下载文件的路径
func (this *ProgramController) GetDownloadFilePath(fileName string) (string, error) {
	// 判断命令是否存在
	fileNamePath := this.sc.ProgramFilePath(fileName)
	if exists, _ := utils.PathExists(fileNamePath); !exists {
		return "", fmt.Errorf("命令文件不存在 %v", fileNamePath)
	}

	return fileNamePath, nil
}

func (this *ProgramController) UploadCreateProgram(file *multipart.FileHeader, c *gin.Context) (map[string]string, error) {
	// 判断文件是否存在
	cnt, err := dao.NewProgramDao().CountByFileName(file.Filename)
	if err != nil {
		return nil, fmt.Errorf("通过文件名获取程序个数失败. %v", err)
	}
	if cnt > 0 {
		return nil, fmt.Errorf("程序:%v, 已经存在.", file.Filename)
	}

	// 获取唯一的临时文件名, 并且进行保存
	tmpFileName := utils.GetUUID()
	filePath := this.sc.UploadProgramFilePath(tmpFileName)
	seelog.Warnf("上传临时文件名称 Name: %v, 路径: %v", file.Filename, filePath)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return nil, fmt.Errorf("保存临时文件出错. %v", err)
	}

	data := make(map[string]string)
	data["file_name"] = file.Filename
	data["tmp_file_name"] = tmpFileName

	return data, nil
}

func (this *ProgramController) UploadEditProgram(
	id int64,
	file *multipart.FileHeader,
	c *gin.Context,
) (map[string]string, error) {
	// 通过名称获取文件
	if program, err := dao.NewProgramDao().GetByName(file.Filename, []string{"id"}); err != nil {
		if err.Error() != "record not found" {
			return nil, fmt.Errorf("通过文件名获取程序失败. %v", err)
		}
	} else {
		// 判断是否和其他命令冲突
		if program.ID.Int64 != id {
			return nil, fmt.Errorf("上传的文件名已经被其他程序使用. %v", err)
		}
	}
	// 获取唯一的临时文件名, 并且进行保存
	tmpFileName := utils.GetUUID()
	filePath := this.sc.UploadProgramFilePath(tmpFileName)
	seelog.Warnf("上传临时文件名称 Name: %v, 路径: %v", file.Filename, filePath)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return nil, fmt.Errorf("保存临时文件出错. %v", err)
	}

	data := make(map[string]string)
	data["file_name"] = file.Filename
	data["tmp_file_name"] = tmpFileName

	return data, nil
}

// 创建
func (this *ProgramController) Create(f *form.ProgramCreateForm) error {
	// 移动临时命令到最终目录
	srcProgramPath := this.sc.UploadProgramFilePath(f.TmpFileName)
	stdProgramPath := this.sc.ProgramFilePath(f.FileName)
	if err := utils.FileCopy(srcProgramPath, stdProgramPath); err != nil {
		return err
	}

	p := &models.Program{
		Title:        types.NewNullString(f.Title, false),
		FileName:     types.NewNullString(f.FileName, false),
		HaveDedicate: f.HaveDedicate,
		Params:       types.NewNullString(f.Params, false),
	}
	if err := dao.NewProgramDao().Create(p, f.DedicateHosts); err != nil {
		return err
	}

	return nil
}

// 通过ID获取数据
func (this *ProgramController) GetByID(id int64) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	program, err := dao.NewProgramDao().GetByID(id, []string{"*"})
	if err != nil {
		return nil, err
	}
	data["program"] = program

	if program.HaveDedicate {
		hosts, err := dao.NewHostDao().FindByProgramID(program.ID.Int64)
		if err != nil {
			return nil, err
		}
		data["hosts"] = hosts
	} else {
		data["hosts"] = make([]models.Host, 0)
	}

	return data, err
}
