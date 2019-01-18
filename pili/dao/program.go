package dao

import (
	"fmt"
	"github.com/daiguadaidai/poow/pili/gdbc"
	"github.com/daiguadaidai/poow/pili/models"
	"github.com/daiguadaidai/poow/utils"
	"github.com/daiguadaidai/poow/utils/types"
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

func (this *ProgramDao) Query(pg *utils.Paginator) ([]models.Program, error) {
	programs := []models.Program{}
	if err := this.DB.Model(&models.Program{}).
		Order("id DESC").Offset(pg.Offset).Limit(pg.Limit).
		Find(&programs).Error; err != nil {
		return nil, err
	}
	return programs, nil
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

// 获取命令个数通过 程序名
func (this *ProgramDao) CountByFileName(name string) (int, error) {
	var cnt int
	if err := this.DB.Model(&models.Program{}).Where("file_name = ?", name).
		Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

// 创建程序
func (this *ProgramDao) Create(p *models.Program, hosts []int64) error {
	tx := this.DB.Begin()

	if err := tx.Model(&models.Program{}).Create(p).Error; err != nil {
		tx.Rollback()
		return err
	}
	if p.HaveDedicate { // 有专用机器
		if len(hosts) == 0 {
			return fmt.Errorf("该程序指定了专用机器. 但是没有给专用机器")
		}
		for _, hostID := range hosts {
			ph := &models.ProgramHost{
				ProgramId: p.ID,
				HostId:    types.NewNullInt64(hostID, false),
			}
			if err := tx.Model(&models.ProgramHost{}).Create(ph).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()

	return nil
}

// 通过程序ID获取数据
func (this *ProgramDao) GetByID(id int64, cols []string) (*models.Program, error) {
	p := new(models.Program)
	if err := this.DB.Model(&models.Program{}).Select(cols).
		Where("id = ?", id).First(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
