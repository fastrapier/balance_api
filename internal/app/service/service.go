package service

import (
	"balance_api/internal/app/model"
	"log"
	"time"
)

type BalanceRepository interface {
	FindByUserId(userId string) (model.Balance, error)
}

type BalanceService struct {
	r BalanceRepository
}

func New(r BalanceRepository) *BalanceService {
	return &BalanceService{
		r: r,
	}
}

func (s *BalanceService) Ping() time.Time {
	return time.Now()
}

func (s *BalanceService) UpdateBalance(userId string) {
	log.Println("User id from service: " + userId)
	balance, err := s.r.FindByUserId(userId)

	if err != nil {
		log.Println(err)
	}

	log.Println(balance)
}
