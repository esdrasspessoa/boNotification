package chi

import (
	"bonitification/internal/app/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(notificationHandler *handler.NotificationHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/notifications", func(r chi.Router) {
		r.Post("/", notificationHandler.CreateNotification)
		r.Get("/{id}", notificationHandler.GetNotification)
		r.Get("/", notificationHandler.GetAllNotifications)
	})

	return router
}

func StartServer(port string, router http.Handler) error {
	return http.ListenAndServe(port, router)
}

