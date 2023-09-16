package main

import (
	"GoStartCode/GinExample/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("0 0/1 * * * ? ", func() {
		log.Println("Run models.CleanAllTag...")
		//models.CleanAllTag()
	})
	c.AddFunc("0 0/2 * * * ? ", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
