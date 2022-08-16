package models

import "time"

type User struct {
	Id        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `form:"username"`
	Password  string `form:"password"`
	Phone     string `form:"phone"`
	Email     string `form:"email"`
}

func (User) TableName() string {
	return "user"
}
