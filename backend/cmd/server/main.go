package main

import (
	"context"
	"log"

	"github.com/dirgefigure-source/enterprise-flow/internal/config"
	"github.com/dirgefigure-source/enterprise-flow/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

func main(){
	cfg := config.Load()

	ctx := context.Background()

	router := gin.Default()

	db, err := infrastructure.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	
	// defer标记的代码，不管函数怎么结束，一定会在最后执行一次
	defer db.Close()

	// 注册http://localhost:8080/health路由，当用户访问这个路由，执行Hanlder，返回HTTP为200的JSON
	router.GET("/health", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	log.Printf("Server starting on :%s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil{
		panic(err)
	}
}