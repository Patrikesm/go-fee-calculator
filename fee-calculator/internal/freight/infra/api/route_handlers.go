package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Patrikesm/fee-calculator/internal/freight/usecase"
)

type RouteHandlers struct {
	CreateRouteUseCase *usecase.CreateRouteUseCase
	UpdateRouteUseCase *usecase.UpdateRouteUseCase
	ListRouteUseCase   *usecase.ListRouteUseCase
}

func NewRouteHandlers(create *usecase.CreateRouteUseCase, update *usecase.UpdateRouteUseCase, listAll *usecase.ListRouteUseCase) *RouteHandlers {
	return &RouteHandlers{
		CreateRouteUseCase: create,
		UpdateRouteUseCase: update,
		ListRouteUseCase:   listAll,
	}
}

func (h *RouteHandlers) CreateRouteHandler(c *gin.Context) {
	var input usecase.CreateRouteInputDTO

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error in field binding",
		})
		fmt.Println(err)
		return
	}

	output, err := h.CreateRouteUseCase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error creating Route",
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *RouteHandlers) ListAllRoutesHandler(c *gin.Context) {
	output, err := h.ListRouteUseCase.Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error getting Routes",
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, output)
}
