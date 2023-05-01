package endpoint

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type BalanceService interface {
	Ping() time.Time
	UpdateBalance(userId string)
}

type Endpoint struct {
	s BalanceService
}

func New(s BalanceService) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Ping(ctx *gin.Context) {
	now := e.s.Ping()

	ctx.JSON(http.StatusOK, gin.H{
		"pong": now,
	})
}

func (e *Endpoint) UpdateBalance(ctx *gin.Context) {
	userId := ctx.Param("userId")

	log.Println("User id from endpoint: " + userId)

	e.s.UpdateBalance(userId)

	ctx.JSON(http.StatusOK, gin.H{
		"test": true,
	})
}
