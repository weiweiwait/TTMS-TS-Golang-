package dao

import (
	"TTMS/conf"
	"TTMS/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		db: conf.NewDBClient(ctx),
	}
}

func (dao *UserDao) IsExistByEmail(email string) bool {
	var count int64
	if dao.db.Table("customer").Model(&model.Customer{}).Where("email = ?", email).Count(&count); count == 1 {
		return true
	}
	return false
}

func (dao *UserDao) CreateCustomer(user *model.Customer) error {
	//插入user
	if err := dao.db.Table("customer").Create(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func (dao *UserDao) UpdatePassword(username string, password string) error {
	return dao.db.Table("customer").Model(&model.Customer{}).Where("username = ?", username).Update("password", password).Error
}

func (dao *UserDao) IsExistByNickName(name string) bool {
	var count int64
	if dao.db.Table("customer").Model(&model.Customer{}).Where("username = ?", name).Count(&count); count == 1 {
		return true
	}
	return false
}

func (dao *UserDao) UpdateUserName(name string, newName string) error {
	return dao.db.Table("customer").Model(&model.Customer{}).Where("username = ?", name).Update("username", newName).Error
}
func (dao *UserDao) GetUser(email string) *model.Customer {
	u := &model.Customer{}
	dao.db.Table("customer").Where("email = ?", email).Find(u)
	return u
}
func (dao *UserDao) GetUserByName(name string) *model.Customer {
	u := &model.Customer{}
	dao.db.Table("customer").Where("email = ?", name).Find(u)
	return u
}
