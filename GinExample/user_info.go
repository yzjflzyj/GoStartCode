package GinExample

import (
	//"GoStartCode/GinExample"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	//ID           uint
	//CreatedAt    time.Time
	//UpdatedAt    time.Time
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func albumOperate() {
	album := Album{Title: "song", Artist: "周杰伦", Price: 246}

	result := db.Create(&album) // 通过数据的指针来创建
	//album.ID             // 返回插入数据的主键
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数
	fmt.Println(result, album.ID, result.Error, result.RowsAffected)
}

func userAdd() {
	users := []*User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
	}
	user := &User{Name: "Jackson", Age: 19, Birthday: time.Now()}
	users = append(users, user)
	result := db.Create(users) // pass a slice to insert multiple ROW
	fmt.Println(result, len(users), result.Error, result.RowsAffected)

	// 下面的批次创建，会开启事物
	//db.CreateInBatches(users, 100)

	// 指定字段创建
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("Jackson", 19, "2023-04-06 11:05:21.775")
	// 忽略字段创建
	db.Omit("Name", "Age", "CreatedAt").Create(&user)
}

func main() {
	//albumOperate()
	userAdd()
}
