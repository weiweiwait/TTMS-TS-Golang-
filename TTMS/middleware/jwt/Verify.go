package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
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
		c.JSON(http.StatusBadRequest, model.BaseResponseInstance.FailMsg(config.TokenIsNotExist))
	} else {
		// 解析签发时间
		tokenTime, err := ParseTokenTime(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.BaseResponseInstance.FailMsg(config.TokenParseErr))
		}
		// 判断时间
		if GetDays(int64(tokenTime.(float64)), time.Now().Unix()) > 30 {
			c.JSON(http.StatusBadRequest, model.BaseResponseInstance.FailMsg(config.TokenIsExpire))
		}
		// 解析正确
		//str := strconv.FormatFloat(Id, 'E', -1, 64)
		//strconv.ParseInt(str, 10, 64)
		c.Set("Id", Id)
		c.Next()
	}
}

//  通过post请求获取token

func VerifyTokenByPost(c *gin.Context) {
	token := c.PostForm("token")
	fmt.Printf("%v \t \n", token)
	if len(token) == 0 {
		//错误 直接
		c.Abort()
		//返回json
		c.JSON(http.StatusBadRequest, model.BaseResponseInstance.FailMsg(config.TokenIsNotExist))
		return
	}
	Id, err := ParseToken(token)
	if err != nil {
		// 解析错误
		c.Abort()
		// 返回json
		c.JSON(http.StatusBadRequest, model.BaseResponseInstance.FailMsg(config.TokenParseErr))
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

type Limiter struct {
	mu         sync.Mutex
	counters   map[string]int
	threshold  int
	windowSize time.Duration
}

// 创建 NewLimiter 函数
func NewLimiter(threshold int, windowSize time.Duration) *Limiter {
	return &Limiter{
		counters:   make(map[string]int),
		threshold:  threshold,
		windowSize: windowSize,
	}
}
func (l *Limiter) Allow(id string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	if counter, ok := l.counters[id]; ok {
		if now.Sub(time.Unix(int64(counter), 0)) <= l.windowSize {
			if counter >= l.threshold {
				return false
			}
			l.counters[id] = counter + 1
		} else {
			l.counters[id] = 1
		}
	} else {
		l.counters[id] = 1
	}

	return true
}
