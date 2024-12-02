package v1

import (
	"github.com/Bakhram74/wallet-app-test-work/internal/service"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	service *service.Service
}

func NewRoutes(service *service.Service) *Routes {
	return &Routes{
		service: service,
	}
}

func (r *Routes) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		v1.POST("/wallet", r.walletOperation)
		v1.GET("/wallets/:walletId", r.getBalance)
	}
}

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}
