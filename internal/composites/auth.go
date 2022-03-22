package composites

import (
	"github.com/Elren44/elren_todo/internal/config"
	"github.com/Elren44/elren_todo/internal/handlers"
	"github.com/Elren44/elren_todo/internal/service"
	postgres2 "github.com/Elren44/elren_todo/internal/storage/postgres"
	"github.com/Elren44/elren_todo/pkg/client/postgres"
	"go.uber.org/zap"
)

type UserComposite struct {
	Storage service.UserStorage
	Service *service.UserService
	Handler handlers.Handler
}

func NewUserComposite(client postgres.Client, logger *zap.SugaredLogger, cfg *config.Config) *UserComposite {
	logger.Debug("creating user repository")
	storage := postgres2.NewUserRepository(logger, client)
	logger.Debug("creating user services")
	service := service.NewUserService(logger, storage)
	logger.Debug("registering user handlers")
	handler := handlers.NewUserHandler(logger, service, cfg)

	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}
}
