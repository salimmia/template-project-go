package dbrepo

import (
	"database/sql"

	"github.com/salimmia/bookings/internal/config"
	"github.com/salimmia/bookings/internal/repository"
)

type mysqlDBRepo struct{
	App *config.AppConfig
	DB *sql.DB
}

func NewMySQLRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo{
	return &mysqlDBRepo{
		App: a,
		DB: conn,
	}
}