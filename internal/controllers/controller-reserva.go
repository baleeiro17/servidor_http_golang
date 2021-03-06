package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servidor_http_golang/internal/models"
	"servidor_http_golang/internal/repositories"
)

func GetReservations(ctx *gin.Context) {

	// get all reservations in database.
	reserva, err := repositories.GetReserva2()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return reservations.
	ctx.JSON(http.StatusOK, reserva)
}

func UpdateReservations(ctx *gin.Context) {
	var reserva models.Result

	// id da reserva.
	id := ctx.Params.ByName("id")

	// status da reserva.
	if err := ctx.ShouldBindJSON(&reserva); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// update reservations in database
	reserva, err := repositories.UpdateReserva(id, reserva)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return reservations.
	ctx.JSON(http.StatusOK, reserva)
}
