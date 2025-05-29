package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// DB 初始化 db 实例的代码
var (
	DB *gorm.DB
)

func InitDBInstance() {
	log.Println("InitDBInstance")
	// 自定义 GORM 日志
	newLogger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold:             time.Second, // 1s 以上的查询认为是慢查询
			LogLevel:                  logger.Info, // 设置日志级别，Info 级别会打印 SQL 语句
			IgnoreRecordNotFoundError: true,        // 忽略 "record not found" 错误
			Colorful:                  true,        // 是否开启彩色日志（终端可见）
		},
	)
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("root:hokitz@tcp(%s:3306)/train_db?parseTime=true", host)), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Default().Panic(err)
	}
	if DB == nil {
		log.Default().Panic("DB is nil")
	}
	// 测试连接
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("ping db success")
}
