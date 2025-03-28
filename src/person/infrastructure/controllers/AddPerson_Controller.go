package controllers

import (
	"api/src/person/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddPersonController struct {
    personService *application.AddPersonService
}

func NewPersonController(personService *application.AddPersonService) *AddPersonController {
    return &AddPersonController{personService: personService}
}

func (c *AddPersonController) AddPerson(ctx *gin.Context) {
    var request struct {
        Name   string `json:"name" binding:"required"`
        Age    int    `json:"age" binding:"required"`
        Gender bool   `json:"gender" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.personService.AddPerson(request.Name, request.Age, request.Gender); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "Person added successfully"})
}
