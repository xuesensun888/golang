package process2

import (
	"encoding/json"
	"fmt"
	"gotest/chatroom/client/utils"
	"gotest/chatroom/common/message"
	"gotest/chatroom/server/model"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//先从mes中取出data部分
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json unmarchal err", err)
		return
	}
	//先声明一个resmes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	// if registerMes.User.UserId == 100 && registerMes.User.UserPwd == "123456" {
	// 	registerResMes.Code = 200
	// 	registerResMes.Error = "注册成功"
	// } else {
	// 	registerResMes.Code = 405
	// 	registerResMes.Error = "用户已存在"
	// }
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXITS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXITS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生错误"
		}

	} else {
		registerResMes.Code = 200
	}

	//序列化registerresmes
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json unmarchal fail ", err)
		return
	}
	//将data赋值给res.data
	resMes.Data = string(data)
	//序列化resmes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json unmarchal fail resmes")
		return
	}
	//发送data数据 封装成writepkg
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmarchal fail error=", err)
		return
	}
	//先声明一个resmes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//在声明一个loginresmes 并完成
	var loginResMes message.LoginResMes
	//如果用户id=100 密码=123456 认为合法 否则不合法
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200
	// 	fmt.Println("登录成功")
	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "用户不存在，请重新注册"
	// }
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		//这里因为用户登录成功，我们就把登录成功的用户放入到usermgr中
		//将登陆成功的用户的userid 赋值给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前的在线用户的id 放入到loginresmes.userid
		//遍历usermgr.onlineusers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Println("登录成功", user.UserId)
	}
	//将loginresmes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json marshal fail ", err)
		return
	}
	//2.将data赋值给resmes
	resMes.Data = string(data)
	//将resmes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("res json marchal fail", err)
		return
	}
	//发送data 我们将其封装成writepkg函数
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onlinesers
	for id, up := range userMgr.onlineUsers {
		//过滤到自己
		if id == userId {
			continue
		}
		up.NotifyMeOnlne(userId)
	}

}
func (this *UserProcess) NotifyMeOnlne(userId int) {
	//定义一个mes结构体
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	//定义一个notifyuserstatusmes结构体
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyuserstatusmes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("marchal.err ==", err)
		return
	}
	//将data赋值给mes
	mes.Data = string(data)
	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal =", err)
		return
	}
	//将data发送给各个用户
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyotheronline err", err)
		return
	}
}
