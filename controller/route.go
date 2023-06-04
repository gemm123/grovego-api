package controller

import (
	"bytes"
	"encoding/json"
	"gemm123/grovego-api/models"
	"gemm123/grovego-api/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type controllerRoute struct {
	serviceRoute service.ServiceRoute
}

func NewControllerRoute(serviceRoute service.ServiceRoute) *controllerRoute {
	return &controllerRoute{serviceRoute: serviceRoute}
}

func (ctr *controllerRoute) RecommendationRoute(c *gin.Context) {
	var input models.InputRecommendationRoute
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed: " + err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	sendRecommendationRouteToML, _ := json.Marshal(input)

	req, err := http.NewRequest("POST", "http://103.67.186.184:5000/recommend", bytes.NewBuffer(sendRecommendationRouteToML))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed: " + err.Error(),
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var recommendationRouteFromML []models.RecommendationRouteFromML
	if err := json.NewDecoder(resp.Body).Decode(&recommendationRouteFromML); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"data":    recommendationRouteFromML,
	})
}

func (ctr *controllerRoute) Finish(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var input models.InputFinishRoute
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed: " + err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	parsedTime, err := time.Parse("15:04:05", input.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed: " + err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	route := models.Route{
		UserID:          userID,
		RouteCoordinate: input.RouteCoordinate,
		RouteName:       input.RouteName,
		Distance:        input.Distance,
		Duration:        parsedTime,
	}

	err = ctr.serviceRoute.CreateRouteUser(route)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed: " + err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
	})
}
