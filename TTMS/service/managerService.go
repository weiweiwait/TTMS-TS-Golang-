package service

import (
	"TTMS/conf"
	"TTMS/dao"
	"TTMS/dto"
	"TTMS/middleware/jwt"
	"TTMS/pkg/e"
	"github.com/gin-gonic/gin"
)

type ManagerService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Username string `form:"username" json:"username"`
	Code     string `form:"code" json:"code"`
	Token    string `form:"token" json:"token"`
}

func NewManagerService() *ManagerService {
	return &ManagerService{}
}

// 管理员登录
func (s *ManagerService) ManagerLoginByPassword(ctx *gin.Context) *dto.Result {
	managerDao := dao.NewManagerDao(ctx)
	//检验密码账号
	if isExist := managerDao.IsExistByEmail(s.Email); !isExist {
		//邮箱未注册
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	manager := managerDao.GetUser(s.Email)
	if manager.Password != s.Password {
		//密码不正确
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	//生成token
	token := jwt.SignTokenManager(manager)
	userDto := dto.BuildManager(manager, token)
	//返回用户信息
	return dto.Success(e.Success, userDto)
}
func (s *ManagerService) ManagerLoginByCode(ctx *gin.Context) *dto.Result {
	managerDao := dao.NewManagerDao(ctx)
	redisClient := conf.NewRedisClient()
	//检验密码账号
	if isExist := managerDao.IsExistByEmail(s.Email); !isExist {
		//邮箱未注册
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	if vCode := redisClient.Get(e.VerificationCodeKey + s.Email).Val(); vCode != s.Code || vCode == "" {
		//验证码错误
		return dto.Fail(e.WrongCode, nil)
	}
	manager := managerDao.GetUser(s.Email)
	//生成token
	token := jwt.SignTokenManager(manager)
	userDto := dto.BuildManager(manager, token)
	//返回用户信息
	return dto.Success(e.Success, userDto)
}
