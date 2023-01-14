package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {
	// 捕获连接属性.
	cfg := mysql.Config{
		//User:   os.Getenv("DBUSER"), 需要配置环境变量
		User: "root",
		//Passwd: os.Getenv("DBPASS"),
		Passwd: "0178116czjg",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "action",
		//AllowNativePasswords: true,
	}
	// 获取数据库句柄.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// 操作数据库
	//查询多行
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	//查询单行
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	//插入数据
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

}

// 插入数据：addAlbum 将指定的专辑添加到数据库中
// 返回新条目的专辑 ID
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// 查询多行：albumsByArtist对具有指定艺术家姓名的专辑进行查询.
func albumsByArtist(name string) ([]Album, error) {
	// 一个专辑切片，用于保存返回行中的数据.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// 遍历行，使用"扫描"将列数据分配给结构字段.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// 查询单行：albumByID查询指定ID的专辑。
func albumByID(id int64) (Album, error) {
	// 一个专辑将保存返回行的数据。
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}
