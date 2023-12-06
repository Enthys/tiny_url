package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type config struct {
	port int
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	cors struct {
		trustedOrigins []string
	}
}

func (c config) Check() error {
	if c.port <= 0 {
		return fmt.Errorf("invalid port number %d", c.port)
	}

	if c.db.dsn == "" {
		return fmt.Errorf("missing database DSN")
	}

	return nil
}

func loadConfig() *config {
	var cfg config

	flag.IntVar(&cfg.port, "port", 80, "API server port")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Func("cors-allowed-origins", "CORS allowed origins", func(s string) error {
		cfg.cors.trustedOrigins = strings.Fields(s)
		return nil
	})

	flag.Parse()

	if os.Getenv("DB_DSN") != "" && cfg.db.dsn == "" {
		cfg.db.dsn = os.Getenv("DB_DSN")
	}

	return &cfg
}
