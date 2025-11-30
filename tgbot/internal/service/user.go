package service

import (
	"tgbot/internal/core/repo"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (r *UserService) Reg(
	tgID int64,
	username string,
) error {
	id := uuid.New()
	timeNow := time.Now()

	if err := r.userRepo.CreateUser(id, tgID, username, timeNow); err != nil {
		return err
	}

	return nil
}
