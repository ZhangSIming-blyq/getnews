package controller

import (
	"getnews/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetNewsList 获取新闻列表, 传递参数 source 和 limit
func GetNewsList(ctx *gin.Context) {
	source := ctx.Query("source")
	limit := ctx.DefaultQuery("limit", "10")
	limitInt, _ := strconv.Atoi(limit)
	news, err := service.GetLatestNews(source, limitInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, news)
}

// RefreshNewsList 刷新新闻列表
func RefreshNewsList(ctx *gin.Context) {
	err := service.RefreshNews()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
