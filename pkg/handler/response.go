package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type statusResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorMessage(ctx *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
