package repository

import (
	"database/sql"
	"fmt"

	"github.com/Patrikesm/fee-calculator/internal/freight/entity"
)

type RouteRepositoryMysql struct {
	DB *sql.DB
}

func NewRouteRepositoryMysql(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		DB: db,
	}
}

func (r *RouteRepositoryMysql) Create(route *entity.Route) error {
	_, err := r.DB.Exec(
		"insert into routes(id, name, distance, status, freight_price) values(?, ?, ?, ?, ?)",
		route.ID, route.Name, route.Distance, route.Status, route.FreightPrice,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RouteRepositoryMysql) FindById(id string) (*entity.Route, error) {
	sqlSmt := "select * from routes where id = ?"
	row := r.DB.QueryRow(sqlSmt, id)

	var startedAt, finishedAt sql.NullTime
	var route entity.Route

	err := row.Scan(
		&route.ID,
		&route.Name,
		&route.Distance,
		&route.Status,
		&route.FreightPrice,
		&startedAt,
		&finishedAt,
	)

	if err != nil {
		return nil, err
	}

	//verificando se o start é valido
	if startedAt.Valid {
		route.StartedAt = startedAt.Time
	}

	if finishedAt.Valid {
		route.FinishedAt = finishedAt.Time
	}

	return &route, nil
}

func (r *RouteRepositoryMysql) Update(route *entity.Route) error {
	startedAt := route.StartedAt.Format("2006-01-02 15:04:05")
	finishedAt := route.FinishedAt.Format("2006-01-02 15:04:05")

	sql := "update routes set status = ?, freight_price = ?, started_at = ?, finished_at = ? where id = ?"

	_, err := r.DB.Exec(sql, route.Status, route.FreightPrice, startedAt, finishedAt, route.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *RouteRepositoryMysql) FindAll() ([]*entity.Route, error) {
	rows, err := r.DB.Query("select * from routes")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var routes []*entity.Route
	var startedAt, finishedAt sql.NullTime

	for rows.Next() {
		var route entity.Route

		err := rows.Scan(
			&route.ID,
			&route.Name,
			&route.Distance,
			&route.Status,
			&route.FreightPrice,
			&startedAt,
			&finishedAt,
		)

		if startedAt.Valid {
			route.StartedAt = startedAt.Time
		}

		if finishedAt.Valid {
			route.FinishedAt = finishedAt.Time
		}

		if err != nil {
			fmt.Println("erro aqui")
			return nil, err
		}

		routes = append(routes, &route)
	}

	return routes, nil
}
