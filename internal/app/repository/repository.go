package repository

import (
	"balance_api/internal/app/model"
	"database/sql"
	"fmt"
	"log"
)

type BalanceRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BalanceRepository {
	return &BalanceRepository{
		db: db,
	}
}

func (r *BalanceRepository) FindByUserId(userId string) (model.Balance, error) {

	log.Println("User id from repository: " + userId)

	var balance model.Balance

	row := r.db.QueryRow("SELECT * from balance where user_id = ?", userId)

	if err := row.Scan(&balance.Id, &balance.Balance, &balance.UserId); err != nil {
		if err == sql.ErrNoRows {
			return balance, fmt.Errorf("balance with user id %d: no such album", userId)
		}
		return balance, fmt.Errorf("balance with user id %d: %v", userId, err)
	}

	return balance, nil
}
