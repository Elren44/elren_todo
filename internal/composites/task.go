package composites

import (
	"github.com/Elren44/elren_todo/internal/config"
	"github.com/Elren44/elren_todo/internal/handlers"
	"github.com/Elren44/elren_todo/internal/service"
	postgres2 "github.com/Elren44/elren_todo/internal/storage/postgres"
	"github.com/Elren44/elren_todo/pkg/client/postgres"
	"go.uber.org/zap"
)

type TaskComposite struct {
	Storage service.TaskStorage
	Service *service.TaskService
	Handler handlers.Handler
}

func NewTaskComposite(client postgres.Client, logger *zap.SugaredLogger, cfg *config.Config) *TaskComposite {
	logger.Debug("creating task repository")
	storage := postgres2.NewTaskRepository(client, logger)
	logger.Debug("creating task services")
	service := service.NewTaskService(storage, logger)
	logger.Debug("registering task handlers")
	handler := handlers.NewTaskHandler(logger, service, cfg)

	return &TaskComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}
}
