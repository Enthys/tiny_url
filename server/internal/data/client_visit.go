package data

import (
	"context"
	"database/sql"
	"time"
)

type ClientVisitModel struct {
	DB *sql.DB
}

func (c ClientVisitModel) Insert(visit ClientVisit) error {
	stmt := `INSERT INTO client_visits (ip, request_url, user_agent, method, origin) VALUES ($1, $2, $3, $4, $5)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.DB.ExecContext(ctx, stmt, visit.IP, visit.RequestUrl, visit.UserAgent, visit.Method, visit.Origin)
	return err
}

type ClientVisit struct {
	IP         string
	Origin     string
	RequestUrl string
	Method     string
	UserAgent  string
	CreatedAt  time.Time
}
