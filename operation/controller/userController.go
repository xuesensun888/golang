package controller

import (
	"fmt"
	"operation/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (con UserController) Add(c *gin.Context) {
	var user models.User

	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println("结构体绑定报错", err)
	}
	err = models.DB.Create(user).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	c.String(200, "增加数据成功")
}
func (con UserController) Index(c *gin.Context) {
	// userList := []models.User{}
	// //查询所有用户,把结果保存到userlist中
	// models.DB.Find(&userList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": userList,
	// })
	// userList := models.User{Id: 5}
	// models.DB.Find(&userList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": userList,
	// })
	// user1 := models.User{}
	// models.DB.First(&user1, 39)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": user1,
	// })
	// user2 := []models.User{}
	// models.DB.Where("id > ? and id < ?", 20, 50).Find(&user2)
	// c.JSON(http.StatusOK, user2)
	// user3 := []models.User{}
	// models.DB.Where("id in ?", []int{5, 49}).Find(&user3)
	// c.JSON(http.StatusOK, user3)
	// user4 := []models.User{}
	// models.DB.Where("email like ?", "%test%").Find(&user4)
	// c.JSON(http.StatusOK, user4)
	// user5 := []models.User{}
	// models.DB.Where("id between ? and ?", 20, 50).Find(&user5)
	// c.JSON(http.StatusOK, user5)
	// user6 := []models.User{}
	// //models.DB.Where("id=? or id = ?", 39, 20).Find(&user6)
	// models.DB.Where("id = ?", 39).Or("id = ?", 5).Find(&user6)
	// c.JSON(http.StatusOK, user6)
	// user7 := []models.User{}
	// models.DB.Select("id", "email").Where("id=?", 5).Or("id=?", 49).Find(&user7)
	// c.JSON(http.StatusOK, user7)
	// user8 := []models.User{}
	// models.DB.Order("id asc").Offset(2).Limit(3).Find(&user8)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": user8,
	// })
	// user9 := []models.User{}
	// var num int64
	// models.DB.Where("id > ? and id < ?", 4, 50).Find(&user9).Count(&num)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": num,
	// })
	type Result struct {
		Username string
		Age      int
	}
	// var result Result
	// models.DB.Table("user").Select("username", "age").Where("username=?", "sxs").Scan(&result)
	// c.JSON(http.StatusOK, result)
	// models.DB.Raw("select username,age from user where username = ?", "sxs2").Scan(&result)
	// c.JSON(200, result)
	var result []models.User
	models.DB.Raw("select * from user").Scan(&result)
	fmt.Println(result)
}
func (con UserController) Edit(c *gin.Context) {
	// userList := models.User{Id: 100}
	// models.DB.Find(&userList)

	// c.JSON(http.StatusOK, gin.H{
	// 	"result": userList,
	// })
	// userList.Username = "你好"
	// userList.Age = 111

	// models.DB.Save(&userList)
	// user1 := models.User{}
	// models.DB.Model(&user1).Where("id=100").Update("username", "大")
	user2 := models.User{}
	models.DB.Where("id = ?", 29).Find(&user2)
	user2.Username = "好"
	user2.Age = 312
	models.DB.Save(&user2)
	c.JSON(200, gin.H{
		"success": user2,
	})
}

func (con UserController) Delete(c *gin.Context) {
	// user := models.User{Id: 29}
	// models.DB.Delete(&user)
	user1 := models.User{}
	models.DB.Where("Age > ?", 25).Delete(&user1)
}
