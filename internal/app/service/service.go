package service

import (
	"bonitification/internal/app/domain/model"
	"bonitification/internal/app/repository"

	"github.com/google/uuid"
)

type NotificationService interface {
	CreateNotification(title, message string) (*model.Notification, error)
	GetNotification(id uuid.UUID) (*model.Notification, error)
	GetAllNotifications() ([]*model.Notification, error)
}

type service struct {
	repo repository.NotificationRepository
}

func NewService(repo repository.NotificationRepository) NotificationService {
	return &service{
		repo: repo,
	}
}

// CreateNotification cria uma nova notificação e a armazena no repositório.
func (s *service) CreateNotification(title string, message string) (*model.Notification, error) {
	notification := model.NewNotification(title, message)
	err := s.repo.Save(notification)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

// GetNotification busca uma notificação pelo ID.
func (s *service) GetNotification(id uuid.UUID) (*model.Notification, error) {
	return s.repo.FindById(id)
}	

// GetAllNotifications retorna todas as notificações armazenadas.
func (s *service) GetAllNotifications() ([]*model.Notification, error) {
	return s.repo.FindAll(), nil 
}




