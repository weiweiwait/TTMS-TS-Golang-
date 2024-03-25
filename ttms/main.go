package main

import (
	router "TTMS_go/ttms/app/api"
	"TTMS_go/ttms/domain/models"
	dto "TTMS_go/ttms/domain/models/dao"
	utils "TTMS_go/ttms/util"
)

func main() {
	utils.InitMysql()
	utils.InitRedis()
	utils.DB.AutoMigrate(models.User{})
	utils.DB.AutoMigrate(dto.UserInfo{})
	r := router.Router()
	r.Run("0.0.0.0:8082")
}
