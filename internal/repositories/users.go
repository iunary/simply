package repositories

import (
	"context"
	"log"

	"github.com/iunary/simply/internal/entities/ent"
	"github.com/spf13/viper"
)

type IUserRepository interface {
	All(c context.Context) ([]*ent.User, error)
}

type UserRepository struct {
	logger *log.Logger
	db     *ent.Client
	v      *viper.Viper
}

func NewUserRepository(db *ent.Client, v *viper.Viper, logger *log.Logger) IUserRepository {
	return &UserRepository{
		logger: logger,
		db:     db,
		v:      v,
	}
}

func (repo *UserRepository) All(c context.Context) ([]*ent.User, error) {
	repo.logger.Println("get user by email from database")
	return repo.db.User.Query().All(c)
}
