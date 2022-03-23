package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/Elren44/elren_todo/internal/model"
	"go.uber.org/zap"
)

const salt = "as54d5532145fhfjk"

type UserStorage interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
	GetOneUser(ctx context.Context, email string) (model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, uuid string, user *model.User) error
	DeleteUser(ctx context.Context, uuid string) error
}

type UserService struct {
	logger     *zap.SugaredLogger
	repository UserStorage
}

func NewUserService(logger *zap.SugaredLogger, repository UserStorage) *UserService {
	return &UserService{
		logger:     logger,
		repository: repository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user model.User) (string, error) {
	user.EncryptedPassword = generatePasswordHash(user.Password)
	_, err := us.repository.GetOneUser(ctx, user.Email)
	if err == nil {
		//us.logger.Errorf("user with this email already exists: %v", err)
		return "", errors.New("user exists")
	}

	return us.repository.CreateUser(ctx, user)
}

func (us *UserService) GetOneUser(ctx context.Context, user model.User) (model.User, error) {

	user, err := us.repository.GetOneUser(ctx, user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
