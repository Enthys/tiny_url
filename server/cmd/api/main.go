package main

import (
	"log"
	"os"

	"github.com/Enthys/book-tracker/internal/data"
	"github.com/Enthys/book-tracker/internal/logger"
	_ "github.com/joho/godotenv/autoload"
)

const version = "1.0.0"

func main() {
	logger := logger.NewZeroLogLogger(os.Stdout)

	config := loadConfig()

	if err := config.Check(); err != nil {
		logger.Fatal(err.Error(), nil)
	}

	db, err := openDB(*config)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}
	logger.Info("Database connection established", nil)

	models := data.Models{
		ShortUrl:     data.ShortUrlModel{DB: db},
		ClientVisits: data.ClientVisitModel{DB: db},
	}

	app := New(config, logger, models)

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
