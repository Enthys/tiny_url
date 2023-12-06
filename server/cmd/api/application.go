package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Enthys/book-tracker/internal/data"
	"github.com/Enthys/book-tracker/internal/logger"
	"github.com/julienschmidt/httprouter"
)

type application struct {
	config *config
	logger logger.Logger
	models data.Models
}

func New(config *config, logger logger.Logger, models data.Models) *application {
	return &application{
		config: config,
		logger: logger,
		models: models,
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.router(),

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     log.New(app.logger, "", 0),
	}

	app.logger.Info("Starting server", map[string]interface{}{"port": app.config.port})
	return srv.ListenAndServe()
}

func (app *application) router() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundError)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedError)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.health)

	return app.recoverPanic(router)
}
