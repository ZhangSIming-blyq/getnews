package controller

import (
	"errors"
	"getnews/pkg/model"
	"getnews/pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the service layer to isolate the tests
type MockService struct {
	mock.Mock
}

func (m *MockService) GetLatestNews(source string, limit int) ([]model.News, error) {
	args := m.Called(source, limit)
	return args.Get(0).([]model.News), args.Error(1)
}

func (m *MockService) RefreshNews() error {
	args := m.Called()
	return args.Error(0)
}

// TestGetNewsList tests the GetNewsList handler
func TestGetNewsList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockService)
	service.GetLatestNews = mockService.GetLatestNews

	t.Run("Success", func(t *testing.T) {
		// Arrange
		expectedNews := []model.News{
			{Title: "Sample News 1", Link: "http://example.com/news1", Source: "Source 1"},
			{Title: "Sample News 2", Link: "http://example.com/news2", Source: "Source 2"},
		}
		mockService.On("GetLatestNews", "", 10).Return(expectedNews, nil)

		// Act
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		GetNewsList(ctx)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `[{"Title":"Sample News 1","Link":"http://example.com/news1","Source":"Source 1"},
			             {"Title":"Sample News 2","Link":"http://example.com/news2","Source":"Source 2"}]`, w.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		// Arrange
		mockService.On("GetLatestNews", "", 10).Return(nil, errors.New("service error"))

		// Act
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		GetNewsList(ctx)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error":"service error"}`, w.Body.String())
	})
}

// TestRefreshNewsList tests the RefreshNewsList handler
func TestRefreshNewsList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockService)
	service.RefreshNews = mockService.RefreshNews

	t.Run("Success", func(t *testing.T) {
		// Arrange
		mockService.On("RefreshNews").Return(nil)

		// Act
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		RefreshNewsList(ctx)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message":"success"}`, w.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		// Arrange
		mockService.On("RefreshNews").Return(errors.New("refresh error"))

		// Act
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		RefreshNewsList(ctx)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error":"refresh error"}`, w.Body.String())
	})
}
