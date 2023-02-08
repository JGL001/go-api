package controller

import (
	"fmt"
	"ginChat/utils"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	req := c.Request
	strFile, header, err := req.FormFile("file")

	if err != nil {
		utils.RespFial(c.Writer, err.Error())
	}
	var suffix string
	fileName := header.Filename
	tem := strings.Split(fileName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	} else {
		utils.RespFial(c.Writer, "暂不支持该文件格式")
	}
	resFileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstFile, err := os.Create("./assets/image/" + fileName)
	if err != nil {
		utils.RespFial(c.Writer, err.Error())
	}
	_, err = io.Copy(dstFile, strFile)
	if err != nil {
		utils.RespFial(c.Writer, err.Error())
	}
	url := "./assets/image/" + resFileName
	utils.RespOk(c.Writer, url, "发送图片成功")
}
