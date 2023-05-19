package main

// 通过test创建数据库表，并新建数据

import (
	"fmt"
	"supplyChainPrivacyComputeSystem/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	 db, err := gorm.Open(mysql.Open("root:mysql123@tcp(127.0.0.1:3306)/CarbonBasic?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// create
	user := &model.User{}
	// user.UserName = "wangbucuo"
	// user.Passwd = "wangbucuo01"
	// user.Identity = 0
	// user.UserDesc = "root"
	// model.CreateUser(*user)

	// read
	fmt.Println(db.First(user, 1))  // 根据整型主键查找
	newuser := model.FindUserByName("wangbucuo")
	fmt.Println(newuser.UserName)
	// db.First(user, "code=?", "D42") // 查找code字段值为D42的记录

	// update
	// db.Model(&user).Update("PassWord", "1234")

}
