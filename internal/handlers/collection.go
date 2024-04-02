package handlers

import (
	"net/http"

	"github.com/artur-karunas/pop-up-museum/internal/entities"
	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllCollection(ctx *gin.Context) {
	authors, items, err := h.services.Collection.GetAll()
	if err != nil {
		errorhandling.SendError(ctx, err)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"authors": authors,
		"items":   items,
	})
}

func (h *Handler) getCollectionById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	items, err := h.services.Collection.GetById(userId)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (h *Handler) addToUserCollection(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.Collection

	if err := ctx.ShouldBind(&input); err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	input.UserId = userId

	err = h.services.Collection.Create(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"itemId": input.ItemId,
	})
}

func (h *Handler) deleteFromUserCollection(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.Collection

	if err := ctx.ShouldBind(&input); err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	input.UserId = userId

	err = h.services.Collection.Delete(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
