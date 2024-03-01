package api

import "github.com/gin-gonic/gin"

func RegisterHTTPEndpoints(router *gin.Engine, rh RouteHandlers) {
	authEndpoints := router.Group("/api")

	{
		authEndpoints.POST("/routes", rh.CreateRouteHandler)
		authEndpoints.GET("/routes", rh.ListAllRoutesHandler)
	}
}
