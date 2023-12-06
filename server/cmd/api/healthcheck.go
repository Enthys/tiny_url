package main

import "net/http"

type healthCheck struct {
	Status     string `json:"status"`
	SystemInfo struct {
		Version string `json:"version"`
	} `json:"system_info"`
}

func (a *application) health(w http.ResponseWriter, r *http.Request) {
	var hc healthCheck

	hc.Status = "available"
	hc.SystemInfo.Version = version

	err := a.writeJSON(w, http.StatusOK, envelope{"health": hc}, nil)
	if err != nil {
		a.logger.Error(err.Error(), nil)
	}
}
