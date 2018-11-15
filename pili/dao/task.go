package dao

import (
	"github.com/daiguadaidai/poow/pili/gdbc"
	"github.com/daiguadaidai/poow/pili/models"
	"github.com/jinzhu/gorm"
)

type TaskDao struct {
	DB *gorm.DB
}

func NewTaskDao() *TaskDao {
	return &TaskDao{
		DB: gdbc.GetOrmInstance().DB,
	}
}

// 创建一个任务
func (this *TaskDao) Create(t *models.Task) error {
	return this.DB.Create(t).Error
}

// 根据uuid 更新任务状态
func (this *TaskDao) UpdateStatusByUUID(uuid string, status int) error {
	return this.DB.Model(&models.Task{}).Where("task_uuid = ?", uuid).
		Update("status", status).Error
}

// 通过uuid 获取任务信息
func (this *TaskDao) GetByUUID(cols []string, uuid string) (*models.Task, error) {
	t := new(models.Task)
	if err := this.DB.Model(&models.Task{}).Select(cols).Where("task_uuid = ?", uuid).
		First(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

// 通过uuid更新任务
func (this *TaskDao) UpdateByUUID(task *models.Task) error {
	ormInstance := gdbc.GetOrmInstance()

	if err := ormInstance.DB.Model(&models.Task{}).
		Where("task_uuid = ?", task.TaskUUID.String).
		Update(task).Error; err != nil {
		return err
	}

	return nil
}
