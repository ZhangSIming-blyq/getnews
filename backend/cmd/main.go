package main

import (
	"fmt"
	"getnews/pkg/config"
	"getnews/pkg/controller"
	"getnews/pkg/model"
	"getnews/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	// 加载配置
	if err := config.Load("config.json"); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 初始化数据库
	db := initDB()

	// 初始化模型和服务
	model.NewNewsModel(db)
	service.InitService(db)
	// 立刻刷新新闻一次
	service.RefreshNews()

	// 设置定时任务
	c := cron.New()
	cfg := config.Get()
	cronSpec := fmt.Sprintf("*/%d * * * *", cfg.Crawler.Interval)
	_, err := c.AddFunc(cronSpec, func() {
		if err := service.RefreshNews(); err != nil {
			log.Printf("Failed to refresh news: %v", err)
		}
	})
	if err != nil {
		log.Fatal("Failed to setup cron job:", err)
	}
	c.Start()

	// 设置路由
	r := gin.Default()

	// CORS 中间件
	r.Use(CORSMiddleware())

	// API 路由
	r.GET("/api/news", controller.GetNewsList)
	r.POST("/api/news/refresh", controller.RefreshNewsList)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// 自动迁移数据库
	if err := db.AutoMigrate(&model.News{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}
