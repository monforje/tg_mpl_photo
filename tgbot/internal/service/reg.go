package service

import (
	"tgbot/internal/core/repo"
)

type RegService struct {
	userRepo repo.UserRepo
}

func NewRegService(userRepo repo.UserRepo) *RegService {
	return &RegService{
		userRepo: userRepo,
	}
}

func (r *RegService) Reg(
	tgID int64,
	username string,
) (int64, error) {
	// TODO: implement
	return 0, nil
}
