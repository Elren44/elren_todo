package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Elren44/elren_todo/internal/model"
	"github.com/Elren44/elren_todo/internal/service"
	"github.com/Elren44/elren_todo/pkg/client/postgres"
	"github.com/Elren44/elren_todo/pkg/utils"
	"github.com/jackc/pgconn"
	"go.uber.org/zap"
)

type userRepository struct {
	logger *zap.SugaredLogger
	client postgres.Client
}

func NewUserRepository(logger *zap.SugaredLogger, client postgres.Client) service.UserStorage {
	return &userRepository{
		logger: logger,
		client: client,
	}
}

func (ur userRepository) CreateUser(ctx context.Context, user model.User) (string, error) {
	q := `INSERT INTO elren_todo.public.users(email, password_hash) VALUES ($1, $2) RETURNING id`

	ur.logger.Debugf(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))
	row := ur.client.QueryRow(ctx, q, user.Email, user.EncryptedPassword)
	if err := row.Scan(&user.UUID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			ur.logger.Error(newErr)
			return "", newErr
		}
		return user.UUID, err
	}
	return user.UUID, nil
}

func (ur userRepository) GetOneUser(ctx context.Context, email string) (model.User, error) {
	var user model.User
	q := `SELECT id, email, password_hash, created_at FROM elren_todo.public.users WHERE email=$1`
	row := ur.client.QueryRow(ctx, q, email)

	if err := row.Scan(&user.UUID, &user.Email, &user.EncryptedPassword, &user.CreatedAt); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			ur.logger.Error(newErr)
			return user, newErr
		}
		return user, err
	}
	return user, nil
}

func (ur userRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (ur userRepository) UpdateUser(ctx context.Context, uuid string, user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (ur userRepository) DeleteUser(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}
