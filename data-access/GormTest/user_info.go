package main

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func dbOperate() {
	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	album := Album{Title: "Jinzhu", Artist: "周杰伦", Price: 123}

	result := db.Create(&album) // 通过数据的指针来创建
	fmt.Println(result)

	//nuser.ID             // 返回插入数据的主键
	//nresult.Error        // 返回 error
	//nresult.RowsAffected // 返回插入记录的条数
}

func main() {
	dbOperate()
}
