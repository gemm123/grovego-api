package controller

import (
	"bytes"
	"encoding/json"
	"gemm123/grovego-api/models"
	"gemm123/grovego-api/service"
	"log"
	"net/http"

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

// func (ctr *controllerRoute) Finish(c *gin.Context) {
// 	userID := c.MustGet("userID").(string)
// 	var input models.InputFinishRoute
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "failed: " + err.Error(),
// 			"status":  http.StatusBadRequest,
// 		})
// 	}

// 	route := models.Route{
// 		UserID: userID,
// 		RouteCoordinate: input.RouteCoordinate,
// 		RouteName: input.RouteName,
// 		Distance: input.Distance,
// 		Duration: time.Now(),
// 	}

// 	err := ctr.serviceRoute.CreateRouteUser(route)
// }
