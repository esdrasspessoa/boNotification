package repository

import (
	"bonitification/internal/app/domain/model"
	"errors"
	"sync"

	"github.com/google/uuid"
)

// ErrNotFound é retornado quando uma notificação não é encontrada.
var ErrNotFound = errors.New("notificação não encontrada")

// define a interface para operações de repositório de notificações.
type NotificationRepository interface {
	Save(notification *model.Notification) error
	FindById(id uuid.UUID) (*model.Notification, error)
	FindAll() []*model.Notification
}

// implementação em memória do NotificationRepository.
type InMemoryNotificationRepository struct {
	mu            sync.RWMutex
	notifications map[uuid.UUID]*model.Notification
}

// NewInMemoryNotificationRepository cria um novo InMemoryNotificationRepository.
func NewInMemoryNotificationRepository() *InMemoryNotificationRepository{
	return &InMemoryNotificationRepository{
		notifications: make(map[uuid.UUID]*model.Notification),
	}
}

func (r *InMemoryNotificationRepository) Save(notification *model.Notification) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.notifications[notification.ID] = notification
	return nil
}

func (r *InMemoryNotificationRepository) FindById(id uuid.UUID) (*model.Notification, error){
	r.mu.RLock()
	defer r.mu.RUnlock()
	notification, ok := r.notifications[id]
	if !ok {
		return nil, ErrNotFound
	}
	return notification, nil
}

func (r *InMemoryNotificationRepository) FindAll() []*model.Notification {
	r.mu.RLock()
	defer r.mu.RUnlock()
	notifications := make([]*model.Notification, 0, len(r.notifications))
	for _, notification := range r.notifications {
		notifications = append(notifications, notification)
	}
	return notifications
}