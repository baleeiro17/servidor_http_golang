package routes

import (
	"github.com/gin-gonic/gin"
	"servidor_http_golang/internal/controllers"
)

func CreateReservations(route *gin.RouterGroup) {
	route.GET("/reservations", controllers.GetReservations)
}
