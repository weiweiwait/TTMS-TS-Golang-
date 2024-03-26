package service

import (
	"TTMS/conf"
	"TTMS/dao"
	"TTMS/dto"
	"TTMS/middleware/jwt"
	"TTMS/model"
	"TTMS/pkg/e"
	"TTMS/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Username string `form:"username" json:"username"`
	Code     string `form:"code" json:"code"`
	Token    string `form:"token" json:"token"`
}

func NewUserService() *UserService {
	return &UserService{}
}

// 发送验证码
func (s *UserService) SendCode() *dto.Result {
	redisClient := conf.NewRedisClient()
	code := e.Success
	//校验email的正确性
	if isTrue := util.VerifyEmailFormat(s.Email); !isTrue {
		code = e.InvalidEmail
		return dto.Fail(code, nil)
	}
	//检查验证码是否重复发送
	if cnt := redisClient.Exists(e.VerificationCodeKey + s.Email).Val(); cnt == 1 {
		//60s内已发送过验证码
		code := e.RepeatSending
		return dto.Fail(code, nil)
	}
	//获取随机验证码并发送
	vCode := util.SendCode(s.Email)
	redisClient.Set(e.VerificationCodeKey+s.Email, vCode, e.VerificationCodeKeyTTL)
	return dto.Success(code, "验证码发送成功")
}

// 用户注册
func (s *UserService) Register(ctx context.Context) *dto.Result {

	redisClient := conf.NewRedisClient()
	userDao := dao.NewUserDao(ctx)

	//检验密码格式
	if isTrue := util.VerifyPasswordFormat(s.Password); !isTrue {
		return dto.Fail(e.WrongPasswordFormat, nil)
	}

	//验证码校验
	vCode := redisClient.Get(e.VerificationCodeKey + s.Email).Val()
	if vCode != s.Code || vCode == "" {
		return dto.Fail(e.WrongCode, nil)
	}

	//判断用户是否已经注册
	if isExist := userDao.IsExistByEmail(s.Email); isExist {
		return dto.Fail(e.RepeatRegister, nil)
	}

	//创建用户并持久化
	user := &model.Customer{
		Email:    s.Email,
		Password: util.Encryption(s.Password),
		Username: s.Username,
	}
	if err := userDao.CreateCustomer(user); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, "注册成功")
}

// 用户登录
func (s *UserService) LoginByPassword(ctx *gin.Context) *dto.Result {
	userDao := dao.NewUserDao(ctx)
	//检验密码账号
	if isExist := userDao.IsExistByEmail(s.Email); !isExist {
		//邮箱未注册
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	user := userDao.GetUser(s.Email)
	if user.Password != util.Encryption(s.Password) {
		//密码不正确
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	//生成token
	token := jwt.SignToken(user)
	userDto := dto.BuildUser(user, token)
	//返回用户信息
	return dto.Success(e.Success, userDto)
}
func (s *UserService) LoginByCode(ctx *gin.Context) *dto.Result {
	userDao := dao.NewUserDao(ctx)
	redisClient := conf.NewRedisClient()
	//检验密码账号
	if isExist := userDao.IsExistByEmail(s.Email); !isExist {
		//邮箱未注册
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	if vCode := redisClient.Get(e.VerificationCodeKey + s.Email).Val(); vCode != s.Code || vCode == "" {
		//验证码错误
		return dto.Fail(e.WrongCode, nil)
	}
	user := userDao.GetUser(s.Email)
	//生成token
	token := jwt.SignToken(user)
	userDto := dto.BuildUser(user, token)
	//返回用户信息
	return dto.Success(e.Success, userDto)
}
