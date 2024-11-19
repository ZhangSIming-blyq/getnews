package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;type:bigint unsigned" json:"id"` // Unsigned bigint
	Title     string    `gorm:"type:varchar(255)" json:"title"`                          // Article Title
	Content   string    `gorm:"type:text" json:"content"`                                // Article Content
	CourseID  uint      `gorm:"type:bigint unsigned;not null" json:"course_id"`          // Foreign key to Course
	CreatedAt time.Time `json:"created_at"`                                              // Created Time
	UpdatedAt time.Time `json:"updated_at"`                                              // Updated Time
}

type Course struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;type:bigint unsigned" json:"id"` // Unsigned bigint
	Title       string    `gorm:"type:varchar(255);uniqueIndex" json:"title"`              // Unique Title
	Description string    `gorm:"type:text" json:"description"`                            // Course Description
	Articles    []Article `gorm:"foreignKey:CourseID" json:"articles"`                     // One-to-many relationship
	CreatedAt   time.Time `json:"created_at"`                                              // Created Time
	UpdatedAt   time.Time `json:"updated_at"`                                              // Updated Time
}

// CourseModel handles course database operations
type CourseModel struct {
	db *gorm.DB
}

func NewCourseModel(db *gorm.DB) *CourseModel {
	return &CourseModel{db: db}
}

// GetCourses 返回所有课程
func (m *CourseModel) GetCourses() ([]Course, error) {
	// Return all courses in the database
	var courses []Course
	err := m.db.Find(&courses).Error
	return courses, err
}

// GetCourseByID 返回指定ID的课程
func (m *CourseModel) GetCourseByID(id string) (*Course, error) {
	var course Course
	err := m.db.Preload("Articles").First(&course, "id = ?", id).Error
	return &course, err
}

// CreateCourse 创建新课程
func (m *CourseModel) CreateCourse(course Course) error {
	return m.db.Create(&course).Error
}

// CreateArticle 创建新文章
func (m *CourseModel) CreateArticle(article Article) error {
	return m.db.Create(&article).Error
}

// GetArticleByID 获取指定ID的文章
func (m *CourseModel) GetArticleByID(id string) (*Article, error) {
	var article Article
	err := m.db.Where("id = ?", id).First(&article).Error
	return &article, err
}
