package routes

import (
	"github.com/gin-gonic/gin"
	"servidor_http_golang/internal/controllers"
)

func CreateRouteReserva(route *gin.RouterGroup) {
	route.GET("/reservations/all/waiting_deploy", controllers.GetReservations)
	route.PUT("reservations/:id/update_logs", controllers.UpdateReservations)
}
