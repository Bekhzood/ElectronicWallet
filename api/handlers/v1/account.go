package v1

import (
	"github.com/Bekhzood/ElectronicWallet/storage"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAccount(c *gin.Context) {
	storage.NewAccountRepo(h.storagePostgres).GetList()
}
