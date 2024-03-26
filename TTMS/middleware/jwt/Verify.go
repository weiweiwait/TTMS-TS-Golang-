package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 核实token，用于中间件的验证
func VerifyToken(c *gin.Context) {
	token := c.Query("token")
	fmt.Printf("%v \t \n", token)
	if len(token) == 0 {
		c.Next()
		return
	}

	Id, err := ParseToken(token)
	if err != nil {
		// 解析错误
		c.Abort()
		// 返回json
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "token is not exist",
		})
	} else {
		// 解析签发时间
		tokenTime, err := ParseTokenTime(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "token is not error",
			})
		}
		// 判断时间
		if GetDays(int64(tokenTime.(float64)), time.Now().Unix()) > 30 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "token is expire",
			})
		}
		// 解析正确
		//str := strconv.FormatFloat(Id, 'E', -1, 64)
		//strconv.ParseInt(str, 10, 64)
		c.Set("Id", Id)
		c.Next()
	}
}

// 通过post请求获取token
func VerifyTokenByPost(c *gin.Context) {
	token := c.PostForm("token")
	fmt.Printf("%v \t \n", token)
	if len(token) == 0 {
		//错误 直接
		c.Abort()
		//返回json
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "token is not exist",
		})
		return
	}
	Id, err := ParseToken(token)
	if err != nil {
		// 解析错误
		c.Abort()
		// 返回json
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "token is erroe",
		})
	} else {
		// 解析正确
		//str := strconv.FormatFloat(Id, 'E', -1, 64)
		//strconv.ParseInt(str, 10, 64)
		c.Set("Id", Id)
		c.Next()
	}
}

func GetDays(start, end int64) int {
	startTime := time.Unix(start, 0)
	endTime := time.Unix(end, 0)
	sub := int(endTime.Sub(startTime).Hours())
	days := sub / 24
	if (sub % 24) > 0 {
		days = days + 1
	}
	return days
}
