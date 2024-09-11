package models

type Fault struct { // 结构体首字母大写, 和数据库表名对应, 默认访问数据表users, 可以设置访问数据表的方法
	Id            int    `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	Type          string `form:"type" json:"type"`
	Branch        string `form:"branch" json:"branch"`
	Person        string `form:"person" json:"person"` // 大驼峰命名
	Processperson string `form:"processperson" json:"processperson"`
	Result        string `form:"result" json:"result"`
}
// 配置数据库操作的表名称
func (Fault) TableName() string {
	return "fault"
}
