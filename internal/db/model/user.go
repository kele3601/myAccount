package model

import "fmt"

type User struct {
	BaseModel
	Name  string `gorm:"column:name" json:"name"`
	Pass  string `gorm:"column:pass" json:"pass"`
	Email string `gorm:"column:email" json:"email"`
	Phone string `gorm:"column:phone" json:"phone"`
	Level int    `gorm:"column:level" json:"level"`
}

func (u *User) BeforeCreate() error {
	return u.BaseModel.beforeCreate()
}

func (u *User) BeforeUpdate() error {
	u.BaseModel.beforeUpdate()
	return nil
}

func (u *User) BeforeDelete() error {
	u.BaseModel.beforeDelete()
	return nil
}

func (u *User) ToString() string {
	return fmt.Sprintf("User:{id: %s, name: %s, pass: %s, email: %s, phone: %s, level: %d}",
		u.ID, u.Name, u.Pass, u.Email, u.Phone, u.Level)
}

func (u *User) TableName() string {
	return "user"
}
