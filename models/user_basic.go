package models

import (
	"fmt"
	"ginChat/utils"
	"time"

	"gorm.io/gorm"
)

// 定义用户基本信息结构体
type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	PassWord      string    `json:"password"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Identity      string    `json:"identity"`
	ClientIp      string    `json:"clientIp"`
	ClientPort    string    `json:"clientPort"`
	LoginTime     time.Time `json:"loginTime"`
	HeartbeatTime time.Time `json:"heartbeatTime"`
	LogoutTime    time.Time `json:"logoutTime"`
	IsLogout      bool      `json:"isLogout"`
	DeviceInfo    string    `json:"deviceInfo"`
	Salt          string    `json:"salt"`
}

// 通过用户名获取用户信息
func FindUserByName(name string) UserBasic {
	// 获取实例
	user := UserBasic{}
	fmt.Println(123)
	utils.DB.Where("name=?", name).First(&user)
	return user
}

// 通过邮箱获取用户信息
func FindUserByEmail(email string) UserBasic {
	// 获取实例
	user := UserBasic{}
	utils.DB.Where("email=?", email).First(&user)
	return user
}

// 通过手机号获取用户信息
func FindUserByPhone(phone string) UserBasic {
	// 获取实例
	user := UserBasic{}
	utils.DB.Where("phone=?", phone).First(&user)
	return user
}

// 创建表名称
func (table *UserBasic) TableName() string {
	return "user_basic"
}

// 创建用户信息
func CreateUser(user UserBasic) UserBasic {
	utils.DB.Create(&user)
	return user
}

// 删除用户信息
func DelectUser(user UserBasic) {
	utils.DB.Delete(&user)
}

// 修改用户信息
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord})
}

// 获取用户列表
func GetUserList() []*UserBasic {
	// 创建一个容量为10的结构体指针切片
	data := make([]*UserBasic, 10)
	// 调用DB方法获取数据
	utils.DB.Find(&data)
	return data
}
