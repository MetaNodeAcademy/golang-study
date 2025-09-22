package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRef"` //重写外键，使用UserRef作为外键
}

type CreditCard struct {
	gorm.Model
	Number  string
	UserRef uint
}

// type User struct {
//   gorm.Model
//   MemberNumber string
//   CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
// }

// type CreditCard struct {
//   gorm.Model
//   Number     string
//   UserNumber string
// }

func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
	return users, err
}

func main() {

}
