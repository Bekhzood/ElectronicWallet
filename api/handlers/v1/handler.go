package v1

import (
	"github.com/Bekhzood/ElectronicWallet/config"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	cfg             config.Config
	storagePostgres *sqlx.DB
}

func New(cfg config.Config, db *sqlx.DB) *Handler {
	return &Handler{
		cfg:             cfg,
		storagePostgres: db,
	}
}
