package dto

import (
	utils "TTMS_go/ttms/util"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Wallet int
	Ticket string
}

func (user UserInfo) TableName() string {
	return "user_info"
}
func FindUserInfo(id string) UserInfo {
	u := &UserInfo{}
	utils.DB.Where("id = ?", id).First(&u)
	return *u
}

func RefreshUserInfo(userInfo UserInfo) {

}
