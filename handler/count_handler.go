package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/rahmat-kurniawan/gin-api-example/model"
    "github.com/rahmat-kurniawan/gin-api-example/service"
    "net/http"
)

type CountHandler interface {
    CountOccurrences(ctx *gin.Context)
}

type countHandler struct {
    service service.CountService
}

func NewCountHandler(service service.CountService) CountHandler {
    return &countHandler{service}
}

func (h *countHandler) CountOccurrences(ctx *gin.Context) {
    var request model.CountRequest
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    count := h.service.CountOccurrences(request.Target, request.Text)

    ctx.JSON(http.StatusOK, model.CountResponse{Count: count})
}
