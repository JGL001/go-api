package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

var ctx = context.Background()

// 初始化配置文件
func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(viper.GetString("mysql.dns"))
}

// 初始化数据库
func InitMySQL() {
	// 添加日志打印
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	fmt.Println("database init done")
}

// 初始化Redis
func InitRedis() {
	// 使用viper插件
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	//  ping 下检验下是否连接通畅
	pone, err := Red.Ping(ctx).Result()
	if err != nil {
		fmt.Println("failed init", err)
	} else {
		fmt.Println("init done", pone)
	}
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(channel, msg string) error {
	var err error
	Red.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	return msg.Payload, err
}
