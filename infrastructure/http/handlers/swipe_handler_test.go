package handlers

import (
	"backend-app/domain/entities"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mocking the SwipeService
type MockSwipeService struct {
	mock.Mock
}

func (m *MockSwipeService) Swipe(userID, profileID uint, action string) error {
	args := m.Called(userID, profileID, action)
	return args.Error(0)
}

func (m *MockSwipeService) GetDailySwipes(userID uint) ([]entities.Swipe, error) {
	args := m.Called(userID)
	return args.Get(0).([]entities.Swipe), args.Error(1)
}

func TestSwipeHandler_Swipe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockSwipeService)
	handler := NewSwipeHandler(mockService)

	t.Run("successful swipe", func(t *testing.T) {
		mockService.On("Swipe", uint(1), uint(1), "like").Return(nil)

		router := gin.Default()
		router.POST("/auth/swipe/:profileID/:action", handler.Swipe)

		req, _ := http.NewRequest(http.MethodPost, "/auth/swipe/1/like", bytes.NewBufferString(""))
		req.Header.Set("userID", "1")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid profileID", func(t *testing.T) {
		router := gin.Default()
		router.POST("/auth/swipe/:profileID/:action", handler.Swipe)

		req, _ := http.NewRequest(http.MethodPost, "/auth/swipe/invalid/like", bytes.NewBufferString(""))
		req.Header.Set("userID", "1")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("invalid action", func(t *testing.T) {
		router := gin.Default()
		router.POST("/auth/swipe/:profileID/:action", handler.Swipe)

		req, _ := http.NewRequest(http.MethodPost, "/auth/swipe/1/invalid", bytes.NewBufferString(""))
		req.Header.Set("userID", "1")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}
