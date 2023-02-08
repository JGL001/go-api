package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 转小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 转大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密  @param plainpwd 用户输入的密码 @param salt 定义的自定义变量
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// 解密 password 存入数据库的密码
func ValidPassword(plainpwd, salt string, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}
