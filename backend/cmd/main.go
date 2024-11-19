// @title			News and Courses API
// @version		1.0
// @description	This API provides endpoints to manage news and courses.
// @contact.name	API Support
// @contact.url	http://example.com/support
// @contact.email	support@example.com
// @host			localhost:8080
// @BasePath		/api
package main

import (
	"fmt"
	_ "getnews/docs"
	"getnews/pkg/config"
	"getnews/pkg/controller"
	"getnews/pkg/model"
	"getnews/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有域
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许所有方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		// 允许所有头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// 允许cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 预检请求有效期
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")
		// 继续处理请求
		if c.Request.Method == http.MethodOptions {
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
	model.NewCourseModel(db)

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

	// 设置路由, 监听到0.0.0.0:8080
	r := gin.Default()

	// CORS 中间件
	r.Use(CORSMiddleware())

	// API 路由
	r.GET("/api/news", controller.GetNewsList)
	r.POST("/api/news/refresh", controller.RefreshNewsList)
	// 获取所有课程
	r.GET("/api/courses", controller.GetCourses)
	// 获取指定课程
	r.GET("/api/courses/:id", controller.GetCourseByID)
	// 创建新课程
	r.POST("/api/courses", controller.CreateCourse)
	// 创建新文章
	r.POST("/api/articles", controller.CreateArticle)
	// 获取指定文章
	r.GET("/api/articles/:id", controller.GetArticleByID)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动服务器
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Server.Port)
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
	if err := db.AutoMigrate(&model.News{}, &model.Course{}, &model.Article{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}
