package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Elren44/elren_todo/internal/model"
	"github.com/Elren44/elren_todo/internal/service"
	"github.com/Elren44/elren_todo/pkg/client/postgres"
	"github.com/Elren44/elren_todo/pkg/utils"
	"go.uber.org/zap"
)

type taskRepository struct {
	client postgres.Client
	logger *zap.SugaredLogger
}

func NewTaskRepository(client postgres.Client, logger *zap.SugaredLogger) service.TaskStorage {
	return &taskRepository{
		client: client,
		logger: logger,
	}
}

func (tr *taskRepository) Create(ctx context.Context, task *model.Task) error {
	//TODO implement me
	panic("implement me")
}

func (tr *taskRepository) GetOne(ctx model.Task, uuid string) (model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (tr *taskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	q := `SELECT * FROM tasks`
	tr.logger.Debugf(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))
	rows, err := tr.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	tasks := make([]model.Task, 0)

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.UUID, &task.Title, &task.Date, &task.Description)
		if err != nil {
			return nil, err
		}
		formattedDate := task.Date.Format(time.RFC1123)
		task.FormattedDate = formattedDate
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *taskRepository) Update(ctx context.Context, uuid string, task *model.Task) error {
	//TODO implement me
	panic("implement me")
}

func (tr *taskRepository) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}
