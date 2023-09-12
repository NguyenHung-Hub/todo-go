package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"social_todo/common"
	ginitem "social_todo/modules/item/transport/gin"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PUT("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}
	r.GET("/ping", func(ctx *gin.Context) {

		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])

		}()
		ctx.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})

	r.Run(":3001")

}
