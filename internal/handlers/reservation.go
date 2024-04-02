package handlers

import (
	"net/http"
	"strconv"

	"github.com/artur-karunas/pop-up-museum/internal/entities"
	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllReservations(ctx *gin.Context) {
	reservations, err := h.services.Reservation.GetAll()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

func (h *Handler) createReservation(ctx *gin.Context) {
	var input entities.CreateReservation

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	userId, err := getUserId(ctx)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	input.UserId = userId

	reservationId, err := h.services.Reservation.Create(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"reservationId": reservationId,
	})
}

func (h *Handler) confirmReservation(ctx *gin.Context) {
	reservationId, err := strconv.Atoi(ctx.Param("reservationId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.UpdateReservation

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	input.ReservationId = reservationId

	err = h.services.Reservation.Update(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getReservationStatuses(ctx *gin.Context) {
	statuses, err := h.services.Reservation.GetStatuses()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, statuses)
}
