package handlers

import (
	"net/http"

	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getInfo(ctx *gin.Context) {
	info, err := h.services.Info.GetInfo()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, info)
}

func (h *Handler) getFAQ(ctx *gin.Context) {
	faq, err := h.services.Info.GetFAQ()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, faq)
}
