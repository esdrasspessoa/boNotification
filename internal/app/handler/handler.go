package handler

import (
	"bonitification/internal/app/domain/model"
	"bonitification/internal/app/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type NotificationHandler struct {
	Service service.NotificationService
}

// NewNotificationHandler cria um novo NotificationHandler com as dependências necessárias.
func NewNotificationHandler(s service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		Service: s,
	}
}

// CreateNotification manipula as requisições POST para criar uma nova notificação.
func (h *NotificationHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	var notif model.Notification
	if err := json.NewDecoder(r.Body).Decode(&notif); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	notification, err := h.Service.CreateNotification(notif.Title, notif.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(notification)
}

// GetNotification manipula as requisições GET para recuperar uma notificação pelo ID.
func (h *NotificationHandler) GetNotification(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	notification, err := h.Service.GetNotification(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(notification)
}

// GetAllNotifications manipula as requisições GET para listar todas as notificações.
func (h *NotificationHandler) GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := h.Service.GetAllNotifications()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(notifications)
}
