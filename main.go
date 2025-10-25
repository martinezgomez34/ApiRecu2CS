package main

import (
    "api/src/person/infrastructure"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    server := infrastructure.NewServer(router)
    server.Run(":8080")
}