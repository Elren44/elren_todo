package service

import (
	"context"
	"fmt"

	"github.com/Elren44/elren_todo/internal/model"
	"go.uber.org/zap"
)

type TaskStorage interface {
	Create(ctx context.Context, task *model.Task) error
	GetOne(ctx model.Task, uuid string) (model.Task, error)
	GetAll(ctx context.Context) ([]model.Task, error)
	Update(ctx context.Context, uuid string, task *model.Task) error
	Delete(ctx context.Context, uuid string) error
}

type TaskService struct {
	repository TaskStorage
	logger     *zap.SugaredLogger
}

func NewTaskService(repository TaskStorage, logger *zap.SugaredLogger) *TaskService {
	return &TaskService{
		repository: repository,
		logger:     logger,
	}
}

func (ts *TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	tasks, err := ts.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get array tasks: %v", err)
	}
	return tasks, nil
}
