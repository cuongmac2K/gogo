package user

import "time"

type User struct {
	Id       int       `json:"id" gorm:"primaryKey,autoIncrement:true"`
	Name     string    `json:"name"`
	UseName  string    `json:"usename"`
	Date     time.Time `json:"date"`
	CreateAt time.Time `json:"createat"`
	UpdateAt time.Time `json:"updateat"`
	PassWord string    `json:"password"`
	Token    string    `json:"token"`
	Note     string    `json:"note"`
}
