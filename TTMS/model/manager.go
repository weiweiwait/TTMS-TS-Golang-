package model

type Manager struct {
	ID       uint   `gorm:"id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Email    string `gorm:"class"`
}
