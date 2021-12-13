package api

import (
	v1 "github.com/Bekhzood/ElectronicWallet/api/handlers/v1"
	"github.com/gin-gonic/gin"
)

func endpointsV1(r *gin.RouterGroup, h *v1.Handler) {
	authRoutes := r.Group("/").Use(h.DigestAuth())
	authRoutes.POST("/accounts/:number", h.CheckAccount)
	authRoutes.POST("/accounts/:number/history", h.GetTransactionsHistory)
	authRoutes.POST("/accounts/:number/balance", h.GetBalance)
	authRoutes.POST("/accounts/:number/balance/send", h.UpdateBalance)
}
