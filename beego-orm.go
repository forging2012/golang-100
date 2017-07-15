package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)

	// 注册定义的 model
	orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// 创建 table
	//orm.RunSyncdb("default", false, true)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.Debug = true
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// 插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新表
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}



///usr/local/go/bin/go run /Users/songyawei/gitroot/golang-100/beego-orm.go
//table `user` already exists, skip
//[ORM]2017/07/13 11:26:32  -[Queries/default] - [  OK /     db.Exec /     0.5ms] - [INSERT INTO `user` (`name`) VALUES (?)] - `slene`
//ID: 2, ERR: <nil>
//[ORM]2017/07/13 11:26:32  -[Queries/default] - [  OK /     db.Exec /     0.5ms] - [UPDATE `user` SET `name` = ? WHERE `id` = ?] - `astaxie`, `2`
//NUM: 1, ERR: <nil>
//[ORM]2017/07/13 11:26:32  -[Queries/default] - [  OK / db.QueryRow /     0.2ms] - [SELECT `id`, `name` FROM `user` WHERE `id` = ? ] - `2`
//ERR: <nil>
//[ORM]2017/07/13 11:26:32  -[Queries/default] - [  OK /     db.Exec /     0.4ms] - [DELETE FROM `user` WHERE `id` = ?] - `2`
//NUM: 1, ERR: <nil>
//
//Process finished with exit code 0



///usr/local/go/bin/go run /Users/songyawei/gitroot/golang-100/beego-orm.go
//create table `user`
//-- --------------------------------------------------
//--  Table Structure for `main.User`
//-- --------------------------------------------------
//CREATE TABLE IF NOT EXISTS `user` (
//`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
//`name` varchar(100) NOT NULL DEFAULT ''
//) ENGINE=InnoDB;
//
//ID: 1, ERR: <nil>
//NUM: 1, ERR: <nil>
//ERR: <nil>
//NUM: 1, ERR: <nil>
//
//Process finished with exit code 0
