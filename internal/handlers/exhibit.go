package handlers

import (
	"net/http"
	"strconv"

	"github.com/artur-karunas/pop-up-museum/internal/entities"
	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllExhibits(ctx *gin.Context) {
	exhibits, err := h.services.Exhibit.GetAll()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, exhibits)
}

func (h *Handler) getExhibitById(ctx *gin.Context) {
	exhibitId, err := strconv.Atoi(ctx.Param("exhibitId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	exhibit, authors, items, err := h.services.Exhibit.GetById(exhibitId)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"exhibitId":          exhibit.ExhibitId,
		"exhibitName":        exhibit.ExhibitName,
		"pathToExhibitImage": exhibit.PathToExhibitImage,
		"startDate":          exhibit.StartDate,
		"endDate":            exhibit.EndDate,
		"type":               exhibit.Type,
		"status":             exhibit.Status,
		"description":        exhibit.Description,
		"location":           exhibit.Location,
		"website":            exhibit.Website,
		"authors":            authors,
		"items":              items,
	})
}

func (h *Handler) createExhibit(ctx *gin.Context) {
	var input entities.CreateExhibit

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	exhibitId, err := h.services.Exhibit.Create(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"exhibitId": exhibitId,
	})
}

func (h *Handler) updateExhibit(ctx *gin.Context) {
	exhibitId, err := strconv.Atoi(ctx.Param("exhibitId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.UpdateExhibit

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	input.ExhibitId = exhibitId

	err = h.services.Exhibit.Update(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteExhibit(ctx *gin.Context) {
	exhibitId, err := strconv.Atoi(ctx.Param("exhibitId"))
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.DeleteExhibit

	if err = ctx.ShouldBind(&input); err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	input.ExhibitId = exhibitId

	err = h.services.Exhibit.Delete(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) updateExhibitImage(ctx *gin.Context) {
	exhibitId, err := strconv.Atoi(ctx.Param("exhibitId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	contextImageHandler := h.imageHandler.NewContextImageHandler(ctx)

	err = h.services.Exhibit.UpdatePathToImage(exhibitId, contextImageHandler)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getExhibitStatuses(ctx *gin.Context) {
	statuses, err := h.services.Exhibit.GetStatuses()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, statuses)
}
