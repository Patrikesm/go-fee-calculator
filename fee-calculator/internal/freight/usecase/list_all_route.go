package usecase

import (
	"time"

	"github.com/Patrikesm/fee-calculator/internal/freight/entity"
)

type ListRouteOutputDTO struct {
	ID           string
	Name         string
	Distance     float64
	Status       string
	FreightPrice float64
	StartedAt    time.Time
	FinishedAt   time.Time
}

type ListRouteUseCase struct {
	Repository entity.RouteRepository
}

func NewListRouteUseCase(repo entity.RouteRepository) *ListRouteUseCase {
	return &ListRouteUseCase{Repository: repo}
}

func (u *ListRouteUseCase) Execute() ([]*ListRouteOutputDTO, error) {
	routes, err := u.Repository.FindAll()

	if err != nil {
		return nil, err
	}

	var output []*ListRouteOutputDTO

	for _, route := range routes {
		output = append(output, &ListRouteOutputDTO{
			ID:           route.ID,
			Name:         route.Name,
			Distance:     route.Distance,
			Status:       route.Status,
			FreightPrice: route.FreightPrice,
			StartedAt:    route.StartedAt,
			FinishedAt:   route.FinishedAt,
		})
	}

	return output, nil
}
