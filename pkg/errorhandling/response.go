package errorhandling

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func SendError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrBadRequest):
		logrus.Info(formatError(err, 400))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.Cause(err).Error())

	case errors.Is(err, ErrUnauthorized):
		logrus.Info(formatError(err, 401))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.Cause(err).Error())

	case errors.Is(err, ErrForbidden):
		logrus.Info(formatError(err, 403))
		ctx.AbortWithStatusJSON(http.StatusForbidden, errors.Cause(err).Error())

	case errors.Is(err, ErrNotFound):
		logrus.Info(formatError(err, 404))
		ctx.AbortWithStatusJSON(http.StatusNotFound, errors.Cause(err).Error())

	case errors.Is(err, ErrInternalServer):
		logrus.Warn(formatError(err, 500))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.Cause(err).Error())
	default:
		logrus.Info(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid request data or parsing troubles")
	}
}

func formatError(err error, code int) string {
	parts := strings.Split(err.Error(), ": ")
	return fmt.Sprintf("[%d] - [%s: %s]", code, parts[1], parts[0])
}
