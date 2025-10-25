package controllers

import (
    "api/src/person/application"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type GetPersonByIDController struct {
    service *application.GetPersonByIDService
}

func NewGetPersonByIDController(service *application.GetPersonByIDService) *GetPersonByIDController {
    return &GetPersonByIDController{service: service}
}

func (c *GetPersonByIDController) GetByID(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    person, err := c.service.GetByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Persona no encontrada"})
        return
    }

    ctx.JSON(http.StatusOK, person)
}
