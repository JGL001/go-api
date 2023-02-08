package controller

import (
	"fmt"
	"ginChat/models"
	"ginChat/utils"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Login
// @Summary 用户登录
// @Tags 用户模块
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string}  json{"code","message","data"}
// @Router /user/login [post]
func Login(c *gin.Context) {
	// 获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	dbuser := models.FindUserByName(name)
	if dbuser.Name == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "用户未注册，请先注册",
		})
		return
	}
	flag := utils.ValidPassword(password, dbuser.Salt, dbuser.PassWord)
	if !flag {
		c.JSON(401, gin.H{
			"message": "账号密码不匹配，请重新登录",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "登录成功",
		"userId":   dbuser.ID,
		"userName": dbuser.Name,
	})
}

// CreateUser
// @Summary 注册用户
// @Tags 用户模块
// @param name formData string fasle "用户名"
// @param password formData string fasle "密码"
// @param repassword formData string fasle "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/registry [post]
func CreateUser(c *gin.Context) {
	// 实例化UserBasic结构体
	user := models.UserBasic{}
	// 获取参数name和password
	name := c.PostForm("name")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	if password != repassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	// 定义获取随机变量
	salt := fmt.Sprintf("%666d", rand.Int31())
	// 通过 name 查找数据库中的数据判断是否该用户已注册过
	data := models.FindUserByName(name)
	fmt.Println(data)
	if data.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "该用户已注册",
		})
	} else {
		user.Name = name
		user.PassWord = utils.MakePassword(password, salt)
		user.Salt = salt
		// fmt.Printf("%#v", user)
		// 将值传入创建数据的表中
		dbuser := models.CreateUser(user)
		// 返回数据
		c.JSON(200, gin.H{
			"message":  "注册成功",
			"userId":   dbuser.ID,
			"userName": dbuser.Name,
		})
	}
}

// DelectUser
// @Summary 删除用户
// @Tags 用户模块
// @param id formData string false "id"
// @Router /user/delectuser [post]
func DelectUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	models.DelectUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 更新用户信息
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @Router /user/updateuser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	password := c.PostForm("password")
	// 定义获取随机变量
	salt := fmt.Sprintf("%666d", rand.Int31())
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "更新成功",
	})
}

// GetUserList
// @Summary 获取用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code","data","message"}
// @router /user/getuserlist [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"data":    data,
		"message": "获取数据成功",
	})
}
