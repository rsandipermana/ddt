package main

import (
    "github.com/gin-gonic/gin"
    "github.com/rahmat-kurniawan/gin-api-example/handler"
    "github.com/rahmat-kurniawan/gin-api-example/service"
)

func main() {
    r := gin.Default()

    maxHandler := handler.NewMaxHandler(service.NewMaxService())
    r.GET("/max", maxHandler.FindMax)

    countHandler := handler.NewCountHandler(service.NewCountService())
    r.POST("/count", countHandler.CountOccurrences)

    r.Run(":8080")
}
