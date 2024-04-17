package handler

import "github.com/umardev500/banksampah/domain"

type userH struct{}

func NewUserHandler() domain.UserHandler {
	return &userH{}
}
