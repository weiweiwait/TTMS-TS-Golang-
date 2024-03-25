package models

import "gorm.io/gorm"

type movie struct {
	gorm.Model
	Info  string
	Name  string
	Money float64
	place place //申请演出厅
}

func (movie movie) TableName() string {
	return "movie_basic"
}
