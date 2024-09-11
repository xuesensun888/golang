package models

type User struct {
	Id       int    `form:"id"`
	Username string `form:"username"`
	Age      int    `form:"age"`
	Email    string `form:"email"`
	AddTime  int    `form:"addtime"`
}

// 配置数据库操作的表名
func (User) TableName() string {
	return "user"
}
