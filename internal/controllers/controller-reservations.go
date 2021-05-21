package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servidor_http_golang/internal/models"
	"servidor_http_golang/internal/repositories"
)

func GetReservations(ctx *gin.Context) {
	var reserva models.Reserva

	// get all reservations in database.
	reserva, err := repositories.GetReserva()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return reservations.
	ctx.JSON(http.StatusOK, reserva)
}
