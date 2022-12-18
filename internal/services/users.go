package services

import (
	"context"
	"log"

	"github.com/iunary/simply/internal/dto"
	"github.com/iunary/simply/internal/repositories"
	"github.com/spf13/viper"
)

type IUserService interface {
	All(c context.Context) ([]*dto.User, error)
}

type UserService struct {
	logger *log.Logger
	repo   repositories.IUserRepository
	v      *viper.Viper
}

func NewUserService(logger *log.Logger, v *viper.Viper, repo repositories.IUserRepository) IUserService {
	return &UserService{
		logger: logger,
		repo:   repo,
		v:      v,
	}
}

func (us UserService) All(c context.Context) ([]*dto.User, error) {
	us.logger.Println("get users service")
	records, err := us.repo.All(c)
	if err != nil {
		return nil, err
	}

	users := make([]*dto.User, 0)
	for _, record := range records {
		users = append(users, &dto.User{
			Email: record.Email,
		})
	}

	return users, err
}
