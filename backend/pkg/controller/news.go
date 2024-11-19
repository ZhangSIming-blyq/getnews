package controller

import (
	"getnews/pkg/model"
	"getnews/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetNewsList godoc
//
//	@Summary		获取新闻列表
//	@Description	获取新闻列表
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			source	query		string	false	"新闻来源"
//	@Param			limit	query		int		false	"数量"
//	@Success		200		{object}	[]model.News
//	@Router			/news [get]
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
//
//	@Summary		刷新新闻列表
//	@Description	强制刷新新闻数据
//	@Tags			news
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/news/refresh [post]
func RefreshNewsList(ctx *gin.Context) {
	err := service.RefreshNews()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetCourses 获取课程列表
//
//	@Summary		获取课程列表
//	@Description	获取所有课程
//	@Tags			courses
//	@Produce		json
//	@Success		200	{array}		model.Course
//	@Failure		500	{object}	map[string]string
//	@Router			/courses [get]
func GetCourses(ctx *gin.Context) {
	courses, err := service.GetCourses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

// GetCourseByID 获取指定ID的课程
//
//	@Summary		获取课程详情
//	@Description	根据课程ID获取课程详细信息
//	@Tags			courses
//	@Produce		json
//	@Param			id	path		string	true	"课程ID"
//	@Success		200	{object}	model.Course
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/courses/{id} [get]
func GetCourseByID(ctx *gin.Context) {
	id := ctx.Param("id")
	course, err := service.GetCourseByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}

// CreateCourse 创建课程
//
//	@Summary		创建新课程
//	@Description	创建一个新的课程
//	@Tags			courses
//	@Accept			json
//	@Produce		json
//	@Param			course	body		model.Course	true	"课程信息"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/courses [post]
func CreateCourse(ctx *gin.Context) {
	var course model.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateCourse(course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// CreateArticle 创建文章
//
//	@Summary		创建新文章
//	@Description	创建一个新的文章
//	@Tags			articles
//	@Accept			json
//	@Produce		json
//	@Param			article	body		model.Article	true	"文章信息"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/articles [post]
func CreateArticle(ctx *gin.Context) {
	var article model.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateArticle(article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetArticleByID 获取指定ID的文章
//
//	@Summary		获取文章详情
//	@Description	根据文章ID获取文章详细信息
//	@Tags			articles
//	@Produce		json
//	@Param			id	path		string	true	"文章ID"
//	@Success		200	{object}	model.Article
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/articles/{id} [get]
func GetArticleByID(ctx *gin.Context) {
	id := ctx.Param("id")
	article, err := service.GetArticleByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, article)
}
