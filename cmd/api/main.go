package main

import (
	"bonitification/internal/app/handler"
	"bonitification/internal/app/repository"
	"bonitification/internal/app/service"
	"bonitification/pkg/http/chi"
	"log"
)

func main() {
	// Repositório e serviço
	repo := repository.NewInMemoryNotificationRepository()
	notifService := service.NewService(repo)

	//Handlers
	notifHandler := handler.NewNotificationHandler(notifService)

	//Configuracao do roteador
	router := chi.NewRouter(notifHandler)

	//Iniciando o servidor
	log.Println("Iniciando servidor na porta 8080...")
	err := chi.StartServer(":8080", router)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor ", err)
	}
}

