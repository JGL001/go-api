package controller

import (
	"fmt"
	"ginChat/models"
	"ginChat/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 防止跨域站点伪造请求
var UpGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := UpGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws)
}

func MsgHandler(ws *websocket.Conn) {
	msg, err := utils.Subscribe(utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]：%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}

func SendUserMsg(c *gin.Context) {
	fmt.Println(1233444, c.Request)
	models.Chat(c.Writer, c.Request)
}

// SearchFriends
// @Summary 聊天模块
// @Tag 获取朋友列表
// @param userId formData {string} false
// @Success  200 {string} json{"code","data","message"}
// @router /user/searchfriends [post]
func SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("userId"))
	fmt.Println("id====", id)
	users := models.SearchFriend(uint(id))
	utils.RespOkList(c.Writer, users, len(users))
}
