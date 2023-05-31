package main

import (
	"fmt"
	"math/big"
)

// 通过test创建数据库表，并新建数据

func main() {
	// db, err := gorm.Open(mysql.Open("root:qhdwsx130324@tcp(127.0.0.1:3306)/CarbonBasic?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// create
	// user := model.User{}
	// user.UserName = "王树新"
	// user.Passwd = "123456"
	// user.Identity = 1
	// user.UserDesc = "审计员"
	// model.CreateUser(user)

	// read
	// fmt.Println(db.First(user, 1))  // 根据整型主键查找
	// newuser := model.FindUserByName("wangbucuo")
	// fmt.Println(newuser.UserName)
	// db.First(user, "code=?", "D42") // 查找code字段值为D42的记录

	// update
	// db.Model(&user).Update("PassWord", "1234")

	// c := ""
	// a, _ := strconv.Atoi(c)
	// fmt.Println(a)

	var num uint64 = 0xc000088000
	bi := big.NewInt(0)
	bi.SetUint64(num)
	fmt.Println(bi)
}
