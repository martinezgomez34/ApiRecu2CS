package controllers

import (
	"api/src/person/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListPersonController struct {
    personService *application.ListPersonService
}

func NewListPersonController(personService *application.ListPersonService) *ListPersonController {
    return &ListPersonController{personService: personService}
}

func (c *ListPersonController) GetAllPersons(ctx *gin.Context) {
    persons, err := c.personService.GetAllPersons()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, persons)
}