package service

import "getnews/pkg/model"

// GetCourses 获取全部Course
func GetCourses() ([]model.Course, error) {
	return courseModel.GetCourses()
}

// GetCourseByID 获取指定ID的Course
func GetCourseByID(id string) (*model.Course, error) {
	return courseModel.GetCourseByID(id)
}

func CreateCourse(course model.Course) error {
	return courseModel.CreateCourse(course)
}

func CreateArticle(article model.Article) error {
	return courseModel.CreateArticle(article)
}

func GetArticleByID(id string) (*model.Article, error) {
	return courseModel.GetArticleByID(id)
}
