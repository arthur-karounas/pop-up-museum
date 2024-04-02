package handlers

import (
	"net/http"

	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	user, appeals, reservations, err := h.services.User.GetById(userId)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"userId":          user.UserId,
		"pathToUserImage": user.PathToUserImage,
		"firstName":       user.FirstName,
		"lastName":        user.LastName,
		"middleName":      user.MiddleName,
		"appeals":         appeals,
		"reservations":    reservations,
	})
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	contextImageHandler := h.imageHandler.NewContextImageHandler(ctx)

	err = h.services.User.Delete(userId, contextImageHandler)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) updateUserImage(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	contextImageHandler := h.imageHandler.NewContextImageHandler(ctx)

	err = h.services.User.UpdatePathToImage(userId, contextImageHandler)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
