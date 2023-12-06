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
	stmt := `INSERT INTO client_visits (ip, request_url, user_agent) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.DB.ExecContext(ctx, stmt, visit.IP, visit.RequestUrl, visit.UserAgent)
	return err
}

type ClientVisit struct {
	IP         string    `json:"ip"`
	RequestUrl string    `json:"request_url"`
	UserAgent  string    `json:"user_agent"`
	CreatedAt  time.Time `json:"created_at"`
}
