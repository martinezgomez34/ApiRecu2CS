package routes

import (
	"api/src/person/application"
	"api/src/person/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, addPersonService *application.AddPersonService, listPersonService *application.ListPersonService, statsService *application.CountByGenderService, getByIDService *application.GetPersonByIDService) {
    addPersonController := controllers.NewPersonController(addPersonService)
	listPersonController := controllers.NewListPersonController(listPersonService)
    statsController := controllers.NewStatsController(statsService)
    getByIDController := controllers.NewGetPersonByIDController(getByIDService)

    api := router.Group("/api")
    {
        api.POST("/addPerson", addPersonController.AddPerson)
        api.GET("/newPersonIsAdded", listPersonController.GetAllPersons)          // Short polling
        api.GET("/countGender", statsController.GetGenderStatsLongPoll)          // Long polling
        api.GET("/martinez/:id", getByIDController.GetByID)
    }
}