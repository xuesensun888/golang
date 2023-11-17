package message

type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，我们必须保证
	//用户信息的json字符串的key 和结构体字段对应的tag名字一直
	UserId     int    `json:"userid"`
	UserPwd    string `json:"userpwd"`
	UserName   string `json:"username"`
	UserStatus int    `json:"userStatus"`
}
