package test

import (
	"bonitification/internal/app/domain/model"
	"bonitification/internal/app/handler"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the NotificationService
type MockNotificationService struct {
	mock.Mock
}

func (m *MockNotificationService) CreateNotification(title, message string) (*model.Notification, error) {
	args := m.Called(title, message)
	return args.Get(0).(*model.Notification), args.Error(1)
}

func (m *MockNotificationService) GetNotification(id uuid.UUID) (*model.Notification, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Notification), args.Error(1)
}

func (m *MockNotificationService) GetAllNotifications() ([]*model.Notification, error) {
	args := m.Called()
	return args.Get(0).([]*model.Notification), args.Error(1)
}

// TestCreateNotification tests the POST /notifications handler
func TestCreateNotificationHandler(t *testing.T) {
	serviceMock := new(MockNotificationService)
	handler := handler.NewNotificationHandler(serviceMock)
	router := chi.NewRouter()
	router.Post("/notifications", handler.CreateNotification)

	notification := &model.Notification{
		Title:   "Test Title",
		Message: "Test Message",
	}
	serviceMock.On("CreateNotification", notification.Title, notification.Message).Return(notification, nil)

	body, _ := json.Marshal(notification)
	req := httptest.NewRequest(http.MethodPost, "/notifications", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	serviceMock.AssertExpectations(t)
}

// TestGetNotification tests the GET /notifications/{id} handler
func TestGetNotificationHandler(t *testing.T) {
	serviceMock := new(MockNotificationService)
	handler := handler.NewNotificationHandler(serviceMock)
	router := chi.NewRouter()
	router.Get("/notifications/{id}", handler.GetNotification)

	id := uuid.New()
	notification := &model.Notification{
		ID:      id,
		Title:   "Test Title",
		Message: "Test Message",
	}
	serviceMock.On("GetNotification", id).Return(notification, nil)

	req := httptest.NewRequest(http.MethodGet, "/notifications/"+id.String(), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	serviceMock.AssertExpectations(t)
}

// TestGetAllNotifications tests the GET /notifications handler
func TestGetAllNotificationsHandler(t *testing.T) {
	serviceMock := new(MockNotificationService)
	handler := handler.NewNotificationHandler(serviceMock)
	router := chi.NewRouter()
	router.Get("/notifications", handler.GetAllNotifications)

	notifications := []*model.Notification{
		{
			Title:   "Test Title 1",
			Message: "Test Message 1",
		},
		{
			Title:   "Test Title 2",
			Message: "Test Message 2",
		},
	}
	serviceMock.On("GetAllNotifications").Return(notifications, nil)

	req := httptest.NewRequest(http.MethodGet, "/notifications", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	serviceMock.AssertExpectations(t)
}
