package v1

import (
	"strconv"

	"github.com/Bekhzood/ElectronicWallet/models"
	"github.com/Bekhzood/ElectronicWallet/storage"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckAccount(c *gin.Context) {
	number := c.Param("number")

	res, err := storage.NewAccountRepo(h.storagePostgres).CheckAccount(number)
	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}

func (h *Handler) GetTransactionsHistory(c *gin.Context) {
	number := c.Param("number")
	res, err := storage.NewAccountRepo(h.storagePostgres).GetHistory(number)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}

func (h *Handler) GetBalance(c *gin.Context) {
	number := c.Param("number")

	res, err := storage.NewAccountRepo(h.storagePostgres).GetBalance(number)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}

func (h *Handler) UpdateBalance(c *gin.Context) {
	var updateInfo models.UpdateBalance

	number := c.Param("number")
	numberConverted, err := strconv.Atoi(number)
	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
	}

	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
	}

	res, err := storage.NewAccountRepo(h.storagePostgres).UpdateBalance(numberConverted, updateInfo)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err.Error())
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}
