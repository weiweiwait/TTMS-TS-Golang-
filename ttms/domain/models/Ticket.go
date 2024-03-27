package models

import (
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	gorm.Model
	//影片名
	Name string
	//价格
	Money float64
	//演出厅
	Place int
	//座位
	Seat seat
	//状态
	Issold bool
	//影片开始结束时间
	Begintime time.Time
	Aftertime time.Time
}

type seat struct {
	x int
	y int
}

func (ticket Ticket) TableName() string {
	return "ticket_basic"
}
