package controllers

import (
	"api/src/person/application"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type StatsController struct {
    statsService *application.CountByGenderService
}

func NewStatsController(statsService *application.CountByGenderService) *StatsController {
    return &StatsController{statsService: statsService}
}

func (c *StatsController) GetGenderStatsLongPoll(ctx *gin.Context) {
    initialMale, initialFemale, err := c.statsService.GetGenderStats()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    timeout := time.After(30 * time.Second) // Timeout after 30 seconds
    tick := time.Tick(1 * time.Second)      // Check every second

    for {
        select {
        case <-timeout:
            ctx.JSON(http.StatusOK, gin.H{
                "male":   initialMale,
                "female": initialFemale,
                "status": "timeout",
            })
            return
        case <-tick:
            currentMale, currentFemale, err := c.statsService.GetGenderStats()
            if err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }

            if currentMale != initialMale || currentFemale != initialFemale {
                ctx.JSON(http.StatusOK, gin.H{
                    "male":   currentMale,
                    "female": currentFemale,
                    "status": "updated",
                })
                return
            }
        }
    }
}