package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) router() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundError)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedError)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.health)

	//  ShortUrl routes
	router.HandlerFunc(http.MethodGet, "/url/:shortUrl", app.RedirectToUrl)
	router.HandlerFunc(http.MethodPost, "/v1/short-urls", app.CreateShortUrl)
	router.HandlerFunc(http.MethodGet, "/v1/short-urls", app.GetShortUrls)
	router.HandlerFunc(http.MethodDelete, "/v1/short-urls/:id", app.DeleteShortUrl)

	return app.enableCORS(app.recordVisit(app.recoverPanic(router)))
}
