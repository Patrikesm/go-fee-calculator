package server

import (
	"database/sql"
	"log"

	"github.com/Patrikesm/fee-calculator/internal/freight/entity"
	"github.com/Patrikesm/fee-calculator/internal/freight/infra/api"
	"github.com/Patrikesm/fee-calculator/internal/freight/infra/repository"
	"github.com/Patrikesm/fee-calculator/internal/freight/usecase"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func NewApp() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/routes?parseTime=true")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewRouteRepositoryMysql(db)
	freight := entity.NewFreight(10)

	createRouteUseCase := usecase.NewCreateRouteUseCase(repository, freight)
	updateRouteUseCase := usecase.NewUpdateRouteUseCase(repository)
	listRouteUseCase := usecase.NewListRouteUseCase(repository)

	RunApi(db, createRouteUseCase, updateRouteUseCase, listRouteUseCase)
}

func RunApi(db *sql.DB, create *usecase.CreateRouteUseCase, update *usecase.UpdateRouteUseCase, listAll *usecase.ListRouteUseCase) {
	router := gin.Default()

	routeHandlers := api.NewRouteHandlers(
		create, update, listAll,
	)

	api.RegisterHTTPEndpoints(router, *routeHandlers)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

// func RunKafka() {
// 	msgChan := make(chan *ckafka.Message)
// 	topics := []string{"routes"}
// 	servers := "host.docker.internal:9094"

// 	go kafka.Consume(topics, servers, msgChan)

// 	for msg := range msgChan {
// 		input := usecase.CreateRouteInputDTO{}
// 		json.Unmarshal(msg.Value, &input)

// 		switch input.Event {
// 		case "RouteCreated":
// 			output, err := createRouteUseCase.Execute(input)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println(output)
// 		case "RouteUpdated", "RouteDeleted":
// 			input := usecase.UpdateRouteInputDTO{}
// 			json.Unmarshal(msg.Value, &input)

// 			output, err := updateRouteUseCase.Excute(input)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println(output)
// 		}
// 	}
// }
