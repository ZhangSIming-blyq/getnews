package service

import (
	"getnews/pkg/model"
	"gorm.io/gorm"
)

var newsModel *model.NewsModel

// InitService Initialize the service with a database connection
func InitService(db *gorm.DB) {
	newsModel = model.NewNewsModel(db)
}

// GetLatestNews 获取最新新闻
func GetLatestNews(source string, limit int) ([]model.News, error) {
	return newsModel.GetLatestBySource(source, limit)
}

// RefreshNews 刷新新闻
func RefreshNews(newsModel *model.NewsModel) error {
	// 定义各服务的抓取函数
	ServiceMap := map[string]func() ([]model.News, error){
		// TODO: 添加更多的新闻来源...
		"Weibo": FetchWeiboHot,
	}

	// 遍历服务列表
	for source, serviceFunc := range ServiceMap {
		// 调用抓取函数，获取新闻列表
		news, err := serviceFunc()
		if err != nil {
			return err
		}

		// 删除该来源的旧新闻
		if err := newsModel.DeleteBySource(source); err != nil {
			return err
		}

		// 将新抓取的新闻批量保存到数据库
		if err := newsModel.BatchCreate(news); err != nil {
			return err
		}
	}
	return nil
}
