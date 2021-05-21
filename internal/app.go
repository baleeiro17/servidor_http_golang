package internal

import (
	"github.com/gin-gonic/gin"
	"servidor_http_golang/internal/routes"
)

func DefineRoutes(router *gin.Engine, rotas string) {

	// main route
	rotasV1 := router.Group(rotas)

	routes.CreateReservations(rotasV1)

}
