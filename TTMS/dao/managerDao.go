package dao

import (
	"TTMS/conf"
	"TTMS/model"
	"context"
	"gorm.io/gorm"
)

type ManagerDao struct {
	db *gorm.DB
}

func NewManagerDao(ctx context.Context) *ManagerDao {
	return &ManagerDao{
		db: conf.NewDBClient(ctx),
	}
}

func (dao *ManagerDao) IsExistByEmail(email string) bool {
	var count int64
	if dao.db.Table("manager").Model(&model.Manager{}).Where("email = ?", email).Count(&count); count == 1 {
		return true
	}
	return false
}

func (dao *ManagerDao) CreateCustomer(user *model.Manager) error {
	//插入user
	if err := dao.db.Table("manager").Create(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func (dao *ManagerDao) UpdatePassword(username string, password string) error {
	return dao.db.Table("manager").Model(&model.Manager{}).Where("username = ?", username).Update("password", password).Error
}

func (dao *ManagerDao) IsExistByNickName(name string) bool {
	var count int64
	if dao.db.Table("manager").Model(&model.Manager{}).Where("username = ?", name).Count(&count); count == 1 {
		return true
	}
	return false
}

func (dao *ManagerDao) UpdateUserName(name string, newName string) error {
	return dao.db.Table("manager").Model(&model.Manager{}).Where("username = ?", name).Update("username", newName).Error
}
func (dao *ManagerDao) GetUser(email string) *model.Manager {
	u := &model.Manager{}
	dao.db.Table("manager").Where("email = ?", email).Find(u)
	return u
}
func (dao *ManagerDao) GetUserByName(name string) *model.Manager {
	u := &model.Manager{}
	dao.db.Table("manager").Where("email = ?", name).Find(u)
	return u
}
