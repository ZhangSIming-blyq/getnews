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
func RefreshNews() error {
	// 微博
	news, err := FetchWeiboHot()
	if err != nil {
		return err
	}

	// 删除该来源的旧新闻
	if err := newsModel.DeleteBySource("Weibo"); err != nil {
		return err
	}

	// 将新抓取的新闻批量保存到数据库
	if err := newsModel.BatchCreate(news); err != nil {
		return err
	}

	// 知乎
	news, err = FetchZhihuHot()
	if err != nil {
		return err
	}

	// 删除该来源的旧新闻
	if err := newsModel.DeleteBySource("Zhihu"); err != nil {
		return err
	}

	// 将新抓取的新闻批量保存到数据库
	if err := newsModel.BatchCreate(news); err != nil {
		return err
	}

	// b站
	news, err = FetchBilibiliPopular()
	if err != nil {
		return err
	}

	// 删除该来源的旧新闻
	if err := newsModel.DeleteBySource("BiliBili"); err != nil {
		return err
	}

	// 将新抓取的新闻批量保存到数据库
	if err := newsModel.BatchCreate(news); err != nil {
		return err
	}

	return nil
}
