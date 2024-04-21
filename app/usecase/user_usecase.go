package usecase

import (
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type userUc struct {
	repo   domain.UserRepository
	client *mongo.Client
}

func NewUserUsecase(repo domain.UserRepository, client *mongo.Client) domain.UserUsecase {
	return &userUc{
		repo:   repo,
		client: client,
	}
}

func (uc *userUc) Create(payload model.CreateUser) error {
	return nil
}
