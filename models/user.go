package models

import (
	"deercoder-gin/util"
	"deercoder-gin/util/db"
)

/*用户*/
type User struct {
	ID           uint          `json:"id"`
	Username     string        `json:"username"`
	Userpassword string        `json:"userpassword"`
	Createtime   util.JsonTime `json:"createtime"`
}

//获得用户,根据id
func GetUserById(id string) interface{} {
	db.DB.AutoMigrate(&User{})
	var user = User{}
	return db.GetDataById(&user, id)
}

//获得用户,分页/查询
func GetUserBySearch(args map[string][]string) interface{} {

	db.DB.AutoMigrate(&User{})
	var user = []*User{}
	return db.GetDataBySearch(&user, "user", args)
}

//删除用户,根据id
func DeleteUserByid(id string) interface{} {

	return db.DeleteDataByName("user", "id", id)
}

//修改用户
func UpdateUser(args map[string][]string) interface{} {

	return db.UpdateData("user", args)
}

//创建用户
func CreateUser(args map[string][]string) interface{} {

	return db.CreateData("user", args)
}

/*//修改账号密码
func UpdateAccount(uid, nickname, userpassword string) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	if err == nil {
		num, _ = res.RowsAffected()
	}
	if num == 0 {
		info = db.MapError
	} else {
		info = db.MapUpdate
	}
	return info
}
*/
