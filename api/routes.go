package api

import (
	v1 "github.com/Bekhzood/ElectronicWallet/api/handlers/v1"
	"github.com/gin-gonic/gin"
)

func endpointsV1(r *gin.RouterGroup, h *v1.Handler) {
	r.GET("/account", h.GetAccount)
}
