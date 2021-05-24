package routes

import (
	"github.com/gin-gonic/gin"
	"servidor_http_golang/internal/controllers"
)

func CreateRouteReserva(route *gin.RouterGroup) {
	route.GET("/reservations", controllers.GetReservations)
	route.PUT("/reservations/:id", controllers.UpdateReservations)
}
