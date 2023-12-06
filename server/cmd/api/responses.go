package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type envelope map[string]interface{}

func (app *application) writeJson(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJson(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) badRequestError(w http.ResponseWriter, err error) {
	app.writeJson(w, http.StatusBadRequest, envelope{"error": err.Error()}, nil)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and count not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request) {
	app.writeJson(w, http.StatusNotFound, envelope{"error": "resource not found"}, nil)
}

func (app *application) methodNotAllowedError(w http.ResponseWriter, r *http.Request) {
	methods := strings.Fields(w.Header().Get("Allow"))
	app.writeJson(
		w,
		http.StatusMethodNotAllowed,
		envelope{
			"error":   "method not allowed",
			"allowed": methods,
		},
		nil,
	)
}
