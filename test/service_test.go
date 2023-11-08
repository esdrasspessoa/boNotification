package test

import (
	"bonitification/internal/app/domain/model"
	"bonitification/internal/app/service"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock for the NotificationRepository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(notification *model.Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}

func (m *MockRepository) FindById(id uuid.UUID) (*model.Notification, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Notification), args.Error(1)
}

func (m *MockRepository) FindAll() []*model.Notification {
	args := m.Called()
	return args.Get(0).([]*model.Notification)
}

// Test the CreateNotification method
func TestCreateNotification(t *testing.T) {
	mockRepo := new(MockRepository)
	notificationService := service.NewService(mockRepo)

	title := "Test Title"
	message := "Test Message"

	mockRepo.On("Save", mock.Anything).Return(nil)

	notification, err := notificationService.CreateNotification(title, message)

	assert.NoError(t, err)
	assert.Equal(t, title, notification.Title)
	assert.Equal(t, message, notification.Message)
	mockRepo.AssertExpectations(t)
}

// Test the GetNotification method
func TestGetNotification(t *testing.T) {
	mockRepo := new(MockRepository)
	notificationService := service.NewService(mockRepo)

	testID := uuid.New()
	testNotification := &model.Notification{ID: testID, Title: "Test Title", Message: "Test Message"}

	mockRepo.On("FindById", testID).Return(testNotification, nil)

	notification, err := notificationService.GetNotification(testID)

	assert.NoError(t, err)
	assert.Equal(t, testNotification, notification)
	mockRepo.AssertExpectations(t)
}

// Test the GetAllNotifications method
func TestGetAllNotifications(t *testing.T) {
	mockRepo := new(MockRepository)
	notificationService := service.NewService(mockRepo)

	testNotifications := []*model.Notification{
		{Title: "Test Title 1", Message: "Test Message 1"},
		{Title: "Test Title 2", Message: "Test Message 2"},
	}

	mockRepo.On("FindAll").Return(testNotifications)

	notifications, err := notificationService.GetAllNotifications()

	assert.NoError(t, err)
	assert.Equal(t, testNotifications, notifications)
	mockRepo.AssertExpectations(t)
}
