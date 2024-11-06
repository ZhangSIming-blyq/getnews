package model

import (
	"gorm.io/gorm"
	"time"
)

// News represents the news model
type News struct {
	Title string `gorm:"not null" json:"title"`
	Link  string `gorm:"not null" json:"link"`
	Rank  int    `json:"rank"`
	// Source means different news sources, such as Weibo, Zhihu, etc.
	Source    string    `gorm:"not null" json:"source"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewsModel handles news database operations
type NewsModel struct {
	db *gorm.DB
}

// NewNewsModel creates an instance of NewsModel with a database connection
func NewNewsModel(db *gorm.DB) *NewsModel {
	return &NewsModel{db: db}
}

// GetLatestBySource fetches the latest news from a specific source, ordered by Rank in descending order
func (m *NewsModel) GetLatestBySource(source string, limit int) ([]News, error) {
	var news []News
	// Order by Rank in descending order
	err := m.db.Table("news").Where("source = ?", source).Order("`rank` ASC").Limit(limit).Find(&news).Error
	return news, err
}

func (m *NewsModel) DeleteBySource(source string) error {
	return m.db.Where("source = ?", source).Delete(&News{}).Error
}

// BatchCreate inserts a batch of news items into the database
func (m *NewsModel) BatchCreate(news []News) error {
	return m.db.Create(&news).Error
}
