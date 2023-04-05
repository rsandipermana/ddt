package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/rsandipermana/ddt/model"
    "github.com/rsandipermana/ddt/service"
    "net/http"
)

type MaxHandler interface {
    FindMax(ctx *gin.Context)
}

type maxHandler struct {
    service service.MaxService
}

func NewMaxHandler(service service.MaxService) MaxHandler {
    return &maxHandler{service}
}

func (h *maxHandler) FindMax(ctx *gin.Context) {
    var request model.MaxRequest
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    max := h.service.FindMax(request.Numbers)

    ctx.JSON(http.StatusOK, model.MaxResponse{Max: max})
}
