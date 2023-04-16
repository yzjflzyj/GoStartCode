package main

import (
	"GoStartCode/data-access/GormTest/models"
	"GoStartCode/data-access/GormTest/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	fmt.Println("先导入fmt包，才能使用")
	r := gin.Default()
	r.Run()
}
