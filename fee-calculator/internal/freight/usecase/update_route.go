package usecase

import (
	"time"

	"github.com/Patrikesm/fee-calculator/internal/freight/entity"
)

type UpdateRouteInputDTO struct {
	ID         string            `json:"id"`
	StartedAt  entity.CustomTime `json:"startedAt"`
	FinishedAt entity.CustomTime `json:"finishedAt"`
	Event      string            `json:"event"`
}

type UpdateRouteOutputDTO struct {
	ID         string            `json:"id"`
	Status     string            `json:"status"`
	StartedAt  entity.CustomTime `json:"startedAt"`
	FinishedAt entity.CustomTime `json:"finishedAt"`
}

type UpdateRouteUseCase struct {
	Repository entity.RouteRepository
}

func NewUpdateRouteUseCase(repo entity.RouteRepository) *UpdateRouteUseCase {
	return &UpdateRouteUseCase{
		Repository: repo,
	}
}

func (u *UpdateRouteUseCase) Excute(input UpdateRouteInputDTO) (*UpdateRouteOutputDTO, error) {
	route, err := u.Repository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	if input.Event == "RouteStarted" {
		route.Start(time.Time(input.StartedAt))
	}

	if input.Event == "RouteFinished" {
		route.Finished(time.Time(input.FinishedAt))
	}

	err = u.Repository.Update(route)

	if err != nil {
		return nil, err
	}

	return &UpdateRouteOutputDTO{
		ID:         route.ID,
		Status:     route.Status,
		StartedAt:  entity.CustomTime(route.StartedAt),
		FinishedAt: entity.CustomTime(route.FinishedAt),
	}, nil
}
