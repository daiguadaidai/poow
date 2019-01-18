package dao

import (
	"fmt"
	"github.com/daiguadaidai/poow/pili/gdbc"
	"github.com/daiguadaidai/poow/pili/models"
	"github.com/jinzhu/gorm"
)

type HostDao struct {
	DB *gorm.DB
}

func NewHostDao() *HostDao {
	return &HostDao{
		DB: gdbc.GetOrmInstance().DB,
	}
}

// 通过程序ID 和 是否有装用机器获取机器
func (this *HostDao) GetByProgramIDAndDedicate(
	proID int64,
	haveDedicate bool,
	cols []string,
) (*models.Host, error) {

	if !haveDedicate { // 没有装用机器, 直接在共用机器里面查找
		return this.GetByProgramIDAndHaveNotDedicate(proID, cols)
	}

	return this.GetByProgramIDAndHaveDedicate(proID, cols)
}

// 获取 hosto专用
func (this *HostDao) GetByProgramIDAndHaveDedicate(proID int64, cols []string) (*models.Host, error) {
	host := new(models.Host)

	// 有专用机器
	if err := this.DB.Table("program_hosts").Select(cols).
		Joins("INNER JOIN hosts ON program_hosts.host_id = hosts.id").
		Where("hosts.is_valid = ?", true).
		Order("hosts.running_task_count").
		First(host).Error; err != nil {
		return nil, fmt.Errorf("获取装用机器失败. %v", err)
	}

	return host, nil
}

// 获取共用机器
func (this *HostDao) GetByProgramIDAndHaveNotDedicate(proID int64, cols []string) (*models.Host, error) {
	host := new(models.Host)

	if err := this.DB.Table("hosts").Select(cols).
		Where("is_dedicate = ? AND is_valid = ?", false, true).
		Order("running_task_count").First(host).Error; err != nil {
		return nil, fmt.Errorf("获取共用机器失败, %v", err)
	}

	return host, nil
}

// 任务数自增1
func (this *HostDao) IncrTaskByHost(host string) error {
	if err := this.DB.Model(&models.Host{}).Where("host = ?", host).
		Update("running_task_count", gorm.Expr("running_task_count + ?", 1)).
		Error; err != nil {
		return err
	}

	return nil
}

// 任务数自减1
func (this *HostDao) DecrTaskByHost(host string) error {
	if err := this.DB.Model(&models.Host{}).Where("host = ?", host).
		Update("running_task_count", gorm.Expr("running_task_count - ?", 1)).
		Error; err != nil {
		return err
	}

	return nil
}

// 更新host是否可用
func (this *HostDao) UpdateIsValidByHost(host string, isValid bool) error {
	return this.DB.Model(&models.Host{}).Where("host = ?", host).
		Update("is_valid", isValid).Error
}

// 获取选择器需要的东西
func (this *HostDao) QueryForSelector() ([]models.Host, error) {
	hosts := []models.Host{}
	if err := this.DB.Model(&models.Host{}).Select("id, host").
		Where("is_valid = 1").Find(&hosts).Error; err != nil {
		return hosts, err
	}
	return hosts, nil
}

// 通过程序id获取hosts
func (this *HostDao) FindByProgramID(programID int64) ([]*models.Host, error) {
	sql := `
SELECT h.*
FROM program_hosts AS ph
INNER JOIN hosts AS h
    ON ph.host_id = h.id
WHERE ph.program_id = ?
    AND h.is_valid = 1;
`

	var hosts []*models.Host
	if err := this.DB.Raw(sql, programID).Find(&hosts).Error; err != nil {
		return nil, err
	}

	return hosts, nil
}
