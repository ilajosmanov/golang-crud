package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)

	if header == "" {
		newErrorMessage(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorMessage(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])

	if err != nil {
		newErrorMessage(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userCtx)

	if !ok {
		newErrorMessage(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)

	if !ok {
		newErrorMessage(ctx, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
