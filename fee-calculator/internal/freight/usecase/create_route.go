package usecase

import "github.com/Patrikesm/fee-calculator/internal/freight/entity"

type CreateRouteInputDTO struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
	Event    string  `json:"event"`
}

type CreateRouteOutputDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Distance     float64 `json:"distance"`
	Status       string  `json:"status"`
	FreightPrice float64 `json:"freight_price"`
}

// coloco a interface aqui para inverter a depêndencia
type CreateRouteUseCase struct {
	Repository entity.RouteRepository
	Freight    entity.FreightInterface
}

func NewCreateRouteUseCase(repository entity.RouteRepository, freight entity.FreightInterface) *CreateRouteUseCase {
	return &CreateRouteUseCase{
		Repository: repository,
		Freight:    freight,
	}
}

func (c *CreateRouteUseCase) Execute(input CreateRouteInputDTO) (*CreateRouteOutputDTO, error) {
	route := entity.NewRoute(input.ID, input.Name, input.Distance)
	c.Freight.Calculate(route)
	err := c.Repository.Create(route)

	if err != nil {
		return nil, err
	}

	return &CreateRouteOutputDTO{
		ID:           route.ID,
		Name:         route.Name,
		Distance:     route.Distance,
		Status:       route.Status,
		FreightPrice: route.FreightPrice,
	}, nil
}
