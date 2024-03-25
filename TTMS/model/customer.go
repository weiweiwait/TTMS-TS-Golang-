package model

type Customer struct {
	ID       uint   `gorm:"id" json:"id"`
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
	Email    string `gorm:"class" json:"email"`
}
