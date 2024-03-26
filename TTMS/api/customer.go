package api

import (
	"TTMS/dto"
	"TTMS/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发送邮箱验证码
func SendCode(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.SendCode())
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

// 用户注册
func Register(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.Register(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

//用户登录

func LoginByPassword(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.LoginByPassword(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func LoginByCode(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.LoginByCode(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}
