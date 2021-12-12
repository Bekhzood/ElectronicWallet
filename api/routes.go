package api

import (
	v1 "github.com/Bekhzood/ElectronicWallet/api/handlers/v1"
	"github.com/gin-gonic/gin"
)

func endpointsV1(r *gin.RouterGroup, h *v1.Handler) {
	authRoutes := r.Group("/").Use(h.DigestAuth())
	authRoutes.POST("/account/:id", h.CheckAccount)
	authRoutes.POST("/account/:id/history", h.GetTransactionsHistory)
	authRoutes.POST("/account/:id/balance", h.GetBalance)
	authRoutes.POST("/account/:id/balance/update", h.UpdateBalance)
}
