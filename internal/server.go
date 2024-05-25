package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	contractsHandler "ddd/internal/application/contracts/handlers"
	clientsInfra "ddd/internal/infrastructure/clients"
	contractsInfra "ddd/internal/infrastructure/contracts"
	managersInfra "ddd/internal/infrastructure/managers"
)

type Server struct {
	// settings
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {

	// создаем репозитории для хэндлеров
	contractRepo := contractsInfra.NewInMemoryRepo()
	clientRepo := clientsInfra.NewInMemoryRepo()
	managerRepo := managersInfra.NewInMemoryRepo()

	// создаем хэндлер для контрактов
	contractHandler := contractsHandler.NewHandler(
		contractRepo,
		managerRepo,
		clientRepo,
	)

	// создаем хэндлер для аварий
	// crashHandler := crashHandler.NewHandler()

	// создаем роутер
	routerChi := chi.NewRouter()

	// настраиваем роутер для маршрутизации контрактов
	s.initContractRoutes(routerChi, contractHandler)

	// настраиваем роутер для маршрутизации аварий
	// s.initCrashRoutes(routerChi, crashHandler)

	return http.ListenAndServe("localhost:8080", routerChi)
}

func (s *Server) initContractRoutes(router *chi.Mux, handler *contractsHandler.Handler) {

	// добавление нового контракта
	router.Post(`/contract/add`, handler.CreateContract)

}
