package dao

import (
	"github.com/daiguadaidai/poow/pili/gdbc"
	"github.com/daiguadaidai/poow/pili/models"
	"github.com/jinzhu/gorm"
)

type ProgramDao struct {
	DB *gorm.DB
}

func NewProgramDao() *ProgramDao {
	return &ProgramDao{
		DB: gdbc.GetOrmInstance().DB,
	}
}

// 通过程序名称获取数据
func (this *ProgramDao) GetByName(name string, cols []string) (*models.Program, error) {
	p := new(models.Program)
	if err := this.DB.Model(&models.Program{}).Select(cols).
		Where("file_name = ?", name).First(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
