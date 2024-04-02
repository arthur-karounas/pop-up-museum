package handlers

import (
	"net/http"
	"strconv"

	"github.com/artur-karunas/pop-up-museum/internal/entities"
	"github.com/artur-karunas/pop-up-museum/pkg/errorhandling"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *Handler) getAllAuthors(ctx *gin.Context) {
	authors, err := h.services.Author.GetAll()
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func (h *Handler) getAuthorById(ctx *gin.Context) {
	authorId, err := strconv.Atoi(ctx.Param("authorId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	author, items, err := h.services.Author.GetById(authorId)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"authorId":          author.AuthorId,
		"authorName":        author.AuthorName,
		"description":       author.Description,
		"pathToAuthorImage": author.PathToAuthorImage,
		"pseudonym":         author.Pseudonym,
		"phoneNumber":       author.PhoneNumber,
		"email":             author.Email,
		"items":             items,
	})
}

func (h *Handler) createAuthor(ctx *gin.Context) {
	var input entities.CreateAuthor

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	authorId, err := h.services.Author.Create(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"authorId": authorId,
	})
}

func (h *Handler) updateAuthor(ctx *gin.Context) {
	authorId, err := strconv.Atoi(ctx.Param("authorId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.UpdateAuthor

	if err := ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	input.AuthorId = authorId

	err = h.services.Author.Update(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteAuthor(ctx *gin.Context) {
	authorId, err := strconv.Atoi(ctx.Param("authorId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	var input entities.DeleteAuthor

	if err = ctx.ShouldBind(&input); err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	input.AuthorId = authorId

	err = h.services.Author.Delete(input)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) updateAuthorImage(ctx *gin.Context) {
	authorId, err := strconv.Atoi(ctx.Param("authorId"))
	if err != nil {
		err = errors.Wrap(errorhandling.ErrBadRequest, err.Error())
		errorhandling.SendError(ctx, err)
		return
	}

	contextImageHandler := h.imageHandler.NewContextImageHandler(ctx)

	err = h.services.Author.UpdatePathToImage(authorId, contextImageHandler)
	if err != nil {
		errorhandling.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
