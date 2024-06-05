package repository

import (
	"trainee3/model/entity/mysql/mysql_trainee3"

	"gorm.io/gorm"
)

type IMysql interface {
	Query(db *gorm.DB, model interface{}, condition string, values ...interface{}) error
	QueryAll(db *gorm.DB, model interface{}, rewardList *[]mysql_trainee3.DSReward, condition string, values ...interface{}) error
	Update(db *gorm.DB, model interface{}) error
	UpdateColumns(db *gorm.DB, model interface{}, cols map[string]interface{}) error
	Save(db *gorm.DB, model interface{}) error
}

type Mysql struct{}

func NewMysql() IMysql {
	return &Mysql{}
}

func (r *Mysql) Query(db *gorm.DB, model interface{}, condition string, values ...interface{}) error {
	if err := db.Model(&model).Where(condition, values...).First(&model).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		return err
	}
	return nil
}
func (r *Mysql) QueryAll(db *gorm.DB, model interface{}, rewardList *[]mysql_trainee3.DSReward, condition string, values ...interface{}) error {
	if err := db.Model(&model).Where(condition, values...).Find(rewardList).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		return err
	}
	return nil
}

func (r *Mysql) Update(db *gorm.DB, model interface{}) error {
	if result := db.Model(&model).Updates(model); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Mysql) UpdateColumns(db *gorm.DB, model interface{}, cols map[string]interface{}) error {
	if result := db.Model(model).UpdateColumns(cols); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Mysql) Save(db *gorm.DB, model interface{}) error {
	if result := db.Save(model); result.Error != nil {
		return result.Error
	}

	return nil
}
