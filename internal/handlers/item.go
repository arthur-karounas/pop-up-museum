package handlers

import (
	"net/http"
	"strconv"

	"github.com/artur-karunas/pop-up-museum/internal/entities"
	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllItems(ctx *gin.Context) {
	items, err := h.services.Item.GetAll()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(ctx *gin.Context) {
	itemId, err := strconv.Atoi(ctx.Param("itemId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	item, err := h.services.Item.GetById(itemId)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (h *Handler) createItem(ctx *gin.Context) {
	var input entities.CreateItem

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	itemId, err := h.services.Item.Create(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"itemId": itemId,
	})
}

func (h *Handler) updateItem(ctx *gin.Context) {
	itemId, err := strconv.Atoi(ctx.Param("itemId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.UpdateItem

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	input.ItemId = itemId

	err = h.services.Item.Update(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteItem(ctx *gin.Context) {
	itemId, err := strconv.Atoi(ctx.Param("itemId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.DeleteItem

	if err = ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	input.ItemId = itemId

	err = h.services.Item.Delete(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) updateItemImage(ctx *gin.Context) {
	itemId, err := strconv.Atoi(ctx.Param("itemId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	contextImageHandler := h.imageHandler.NewContextImageHandler(ctx)

	err = h.services.Item.UpdatePathToImage(itemId, contextImageHandler)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getItemStatuses(ctx *gin.Context) {
	statuses, err := h.services.Item.GetStatuses()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, statuses)
}
