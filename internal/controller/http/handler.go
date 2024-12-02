package http

import (
	"github.com/Bakhram74/wallet-app-test-work/config"

	v1 "github.com/Bakhram74/wallet-app-test-work/internal/controller/http/v1"
	"github.com/Bakhram74/wallet-app-test-work/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	config  *config.Config
	service *service.Service
}

func NewHandler(config *config.Config, service *service.Service) *Handler {

	return &Handler{
		service: service,
		config:  config,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	api := router.Group("/api")
	handlerV1 := v1.NewRoutes(h.service)
	{
		handlerV1.Init(api)
	}
}
