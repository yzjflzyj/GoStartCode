package GinExample

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db       *gorm.DB
	linkOnce sync.Once
)

// 用于连接数据库
func linkDb() {
	linkOnce.Do(func() {
		dsn := "root:0178116czjg@tcp(localhost:3306)/action?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
}

func init() {
	linkDb()
}
