package v1

import (
	"github.com/Bekhzood/ElectronicWallet/models"
	"github.com/Bekhzood/ElectronicWallet/storage"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckAccount(c *gin.Context) {
	id := c.Param("id")

	res, err := storage.NewAccountRepo(h.storagePostgres).CheckAccount(id)
	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}

func (h *Handler) GetTransactionsHistory(c *gin.Context) {
	storage.NewAccountRepo(h.storagePostgres).GetHistory()
}

func (h *Handler) GetBalance(c *gin.Context) {
	id := c.Param("id")

	res, err := storage.NewAccountRepo(h.storagePostgres).GetBalance(id)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}

func (h *Handler) UpdateBalance(c *gin.Context) {
	number := c.Param("id")
	wallet := models.Wallet{}

	if err := c.ShouldBindJSON(&wallet); err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
	}

	res, err := storage.NewAccountRepo(h.storagePostgres).UpdateBalance(number, wallet.Balance)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}
