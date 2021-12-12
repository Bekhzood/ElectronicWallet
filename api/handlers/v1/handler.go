package v1

import (
	"fmt"

	"github.com/Bekhzood/ElectronicWallet/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	cfg             config.Config
	storagePostgres *sqlx.DB
}

// SuccessModel ...
type SuccessModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorModel ...
type ErrorModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func New(cfg config.Config, db *sqlx.DB) *Handler {
	return &Handler{
		cfg:             cfg,
		storagePostgres: db,
	}
}

func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, SuccessModel{
		Code:    fmt.Sprint(code),
		Message: message,
		Data:    data,
	})
}

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, ErrorModel{
		Code:    fmt.Sprint(code),
		Message: message,
		Error:   err,
	})
}
