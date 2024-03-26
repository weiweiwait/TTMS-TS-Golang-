package api

import (
	"TTMS/dto"
	"TTMS/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//管理员登录

func ManagerLoginByPassword(ctx *gin.Context) {
	managerService := service.NewManagerService()
	if err := ctx.ShouldBind(managerService); err == nil {
		ctx.JSON(http.StatusOK, managerService.ManagerLoginByPassword(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func ManagerLoginByCode(ctx *gin.Context) {
	managerService := service.NewManagerService()
	if err := ctx.ShouldBind(managerService); err == nil {
		ctx.JSON(http.StatusOK, managerService.ManagerLoginByCode(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}
