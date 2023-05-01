package app

import (
	"balance_api/internal/app/endpoint"
	"balance_api/internal/app/repository"
	"balance_api/internal/app/service"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type App struct {
	r   *repository.BalanceRepository
	s   *service.BalanceService
	e   *endpoint.Endpoint
	gin *gin.Engine
	db  *sql.DB
}

func New() (*App, error) {
	a := &App{}

	a.r = repository.New(a.db)

	a.s = service.New(a.r)

	a.e = endpoint.New(a.s)

	a.gin = gin.Default()

	a.gin.GET("/ping", a.e.Ping)

	a.gin.POST("/balance/:userId", a.e.UpdateBalance)

	return a, nil
}

func (a *App) setUpDatabase() error {
	pingErr := a.db.Ping()

	if pingErr != nil {
		panic("Ping database failed!")
	}
	_, err := a.db.Exec("USE balance")

	if err != nil {
		_, err := a.db.Exec("CREATE DATABASE balance")

		if err != nil {
			panic(err)
		}

		_, err = a.db.Exec("USE balance")

		if err != nil {
			panic(err)
		}
	}

	_, err = a.db.Query("SELECT * from balance;")

	if err != nil {
		log.Println("table not there")
		_, err := a.db.Exec("CREATE TABLE balance (id integer primary key, balance integer, user_id integer)")
		if err != nil {
			panic(err)
		}
		log.Println("table created")
	} else {
		log.Println("table is here")
	}

	return nil
}

func (a *App) Run() error {

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
	}

	var err error

	a.db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(a.db)

	err = a.setUpDatabase()

	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("PORT", "10305")

	if err != nil {
		log.Fatal(err)
	}

	err = a.gin.Run()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("server running")
	return nil
}
