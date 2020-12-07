package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app"
)

func (h *Handler) createList(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		newErrorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		newErrorMessage(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		newErrorMessage(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, id, input); err != nil {
		newErrorMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		newErrorMessage(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
