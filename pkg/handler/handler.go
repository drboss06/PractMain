package handler

import (
	"authPract/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	team := router.Group("/team", h.userIdentity)
	{
		team.POST("/create", h.createTeam)
		team.POST("/send", h.sendMail)
	}
	return router
}
