package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/fatih/set"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// 消息体
type Message struct {
	gorm.Model
	FormId      int64  `json:"formId"`      // 发送者
	TargetId    int64  `json:"targetId"`    // 接收者
	ChatType    int8   `json:"chatType"`    // 聊天类型 群聊 私聊 广播
	MessageType int8   `json:"messageType"` // 消息类型 文字 图片 音频/视频
	Content     string `json:"content"`     // 消息内容
	Pic         string `json:"pic"`         // 图片
	Url         string `json:"url"`         // 上传相关的连接
	Desc        string `json:"desc"`        // 描述
	Amount      int    // 其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// 发送和接收消息
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 获取参数并做token校验
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	fmt.Println(userId, query)
	// msgType := query.Get("messageType")
	// targetId := query.Get("targetId")
	// context := query.Get("context")
	isValida := true
	conn, err := (&websocket.Upgrader{
		// token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isValida
		},
	}).Upgrade(writer, request, nil)
	// defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// var uuid = uuid.NewV4().String()
	// 获取conn 实例化Node结构体
	node := &Node{
		// Id:        uuid,
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 用户关系(先空着)
	// userId和node实例绑定 并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	// 完成发送逻辑
	go sendProc(node)
	// 完成接收逻辑
	go recvProc(node)
	// sendMsg(userId, []byte("欢迎进入聊天室"))
}

func sendProc(n *Node) {
	for {
		select {
		case data := <-n.DataQueue:
			fmt.Println("sendProc +======", string(data))
			err := n.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("recvProc +======", string(data))
		// dispatch(data)
		broadMsg(data)
		fmt.Println(data)
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成udp数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		select {
		case data := <-udpsendChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	fmt.Println("udpRecvProc +======")
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("msgtype +======", msg.ChatType)
	switch msg.ChatType {
	case 1: // 私信
		fmt.Println("dispatch +======", string(data))
		sendMsg(msg.TargetId, data)
	}
}

func sendMsg(userId int64, msg []byte) {
	fmt.Println("sendMsg +======", string(msg), userId)
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
