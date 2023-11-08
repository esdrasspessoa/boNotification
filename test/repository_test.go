package test

import (
	"bonitification/internal/app/domain/model"
	"bonitification/internal/app/repository"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestSave checks if the notification is saved correctly
func TestSave(t *testing.T) {
	repo := repository.NewInMemoryNotificationRepository()
	notification := model.NewNotification("Test Title", "Test Message")

	err := repo.Save(notification)
	assert.NoError(t, err)

	savedNotification, err := repo.FindById(notification.ID)
	assert.NoError(t, err)
	assert.Equal(t, notification, savedNotification)
}

// TestFindById checks if the notification can be retrieved by its ID
func TestFindById(t *testing.T) {
	repo := repository.NewInMemoryNotificationRepository()
	testID := uuid.New()
	_, err := repo.FindById(testID)
	assert.Error(t, err, repository.ErrNotFound)

	notification := model.NewNotification("Test Title", "Test Message")
	repo.Save(notification)

	foundNotification, err := repo.FindById(notification.ID)
	assert.NoError(t, err)
	assert.Equal(t, notification, foundNotification)
}

// TestFindAll checks if all notifications are returned
func TestFindAll(t *testing.T) {
	repo := repository.NewInMemoryNotificationRepository()
	notification1 := model.NewNotification("Test Title 1", "Test Message 1")
	notification2 := model.NewNotification("Test Title 2", "Test Message 2")

	repo.Save(notification1)
	repo.Save(notification2)

	notifications := repo.FindAll()
	assert.Len(t, notifications, 2)
	assert.Contains(t, notifications, notification1)
	assert.Contains(t, notifications, notification2)
}
