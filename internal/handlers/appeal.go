package handlers

import (
	"net/http"
	"strconv"

	"github.com/artur-karunas/pop-up-museum/internal/entities"
	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllAppeals(ctx *gin.Context) {
	appeals, err := h.services.Appeal.GetAll()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, appeals)
}

func (h *Handler) createAppeal(ctx *gin.Context) {
	var input entities.CreateAppeal

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

	appealId, err := h.services.Appeal.Create(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"appealId": appealId,
	})
}

func (h *Handler) confirmAppeal(ctx *gin.Context) {
	appealId, err := strconv.Atoi(ctx.Param("appealId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.UpdateAppeal

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	input.AppealId = appealId

	err = h.services.Appeal.Update(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getAppealStatuses(ctx *gin.Context) {
	statuses, err := h.services.Appeal.GetStatuses()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, statuses)
}
