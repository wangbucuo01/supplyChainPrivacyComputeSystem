package model

import (
	"supplyChainPrivacyComputeSystem/utils"

	"gorm.io/gorm"
)

type User struct {
	UID      int    `json:"uid"`
	UserName string `json:"username"`
	Passwd   string `json:"passwd"`
	Identity int    `json:"identity"`
	State    int    `json:"state"`
	UserDesc string `json:"user_desc"`
}

func (u User) TableName() string {
	return "user"
}

func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

func FindUserByName(name string) User {
	user := User{}
	utils.DB.Where("username = ?", name).First(&user)
	return user
}
