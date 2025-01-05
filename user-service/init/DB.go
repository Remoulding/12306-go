package init

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

// 初始化 db 实例的代码
var (
	once sync.Once
	DB   *gorm.DB
)

func InitDBInstance() {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(mysql.Open("user:password@tcp(localhost:3306)/dbname"), &gorm.Config{})
		if err != nil {
			log.Default().Panic(err)
		}
	})
}
