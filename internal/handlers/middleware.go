package handlers

import (
	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) optionalIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		ctx.Set("userId", 0)
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(header)
	if err != nil {
		err := errors.Wrap(errorhandling.ErrUnauthorized, "expired token")
		errorhandling.SendError(ctx, err)
		return
	}

	if role != 0 {
		err := errors.Wrap(errorhandling.ErrForbidden, "no access")
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.Set("userId", userId)
}

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		err := errors.Wrap(errorhandling.ErrUnauthorized, "empty authorization field")
		errorhandling.SendError(ctx, err)
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(header)
	if err != nil {
		err := errors.Wrap(errorhandling.ErrUnauthorized, "expired token")
		errorhandling.SendError(ctx, err)
		return
	}

	if role != 0 {
		err := errors.Wrap(errorhandling.ErrForbidden, "no access")
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.Set("userId", userId)
}

func (h *Handler) adminIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		err := errors.Wrap(errorhandling.ErrUnauthorized, "empty authorization field")
		errorhandling.SendError(ctx, err)
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(header)
	if err != nil {
		err := errors.Wrap(errorhandling.ErrUnauthorized, "expired token")
		errorhandling.SendError(ctx, err)
		return
	}

	if role != 1 {
		err := errors.Wrap(errorhandling.ErrForbidden, "no access")
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.Set("userId", userId)
}

func (h *Handler) moderatorIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		err := errors.Wrap(errorhandling.ErrUnauthorized, "empty authorization field")
		errorhandling.SendError(ctx, err)
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(header)
	if err != nil {
		err := errors.Wrap(errorhandling.ErrUnauthorized, "expired token")
		errorhandling.SendError(ctx, err)
		return
	}

	if role != 2 {
		err := errors.Wrap(errorhandling.ErrForbidden, "no access")
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.Set("userId", userId)
}

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get("userId")
	if !ok {
		return 0, errors.Wrap(errorhandling.ErrBadRequest, "empty userId parameter")
	}

	userIdInt, ok := userId.(int)
	if !ok {
		return 0, errors.Wrap(errorhandling.ErrForbidden, "invalid userId parameter")
	}

	return userIdInt, nil
}
