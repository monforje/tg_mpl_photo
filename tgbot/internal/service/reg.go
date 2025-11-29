package service

import (
	"tgbot/internal/core/repo"
	"time"

	"github.com/google/uuid"
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
) error {
	id := uuid.New()
	timeNow := time.Now()

	if err := r.userRepo.CreateUser(id, tgID, username, timeNow); err != nil {
		return err
	}

	return nil
}
