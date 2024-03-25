package models

import (
	"gorm.io/gorm"
	"time"
)

type place struct {
	gorm.Model
	Name      string
	Seat      [][]bool
	N         int
	M         int
	Movie     int
	Inuse     bool
	Begintime time.Time
	Aftertime time.Time
}

func (p *place) Init() {
	p.Seat = make([][]bool, p.N)
	for i := 0; i < p.N; i++ {
		p.Seat[i] = make([]bool, p.M)
	}
}
func (place place) TableName() string {
	return "place_basic"
}
