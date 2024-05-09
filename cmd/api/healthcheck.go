package main

import (
	"net/http"
)

func (app *application) healthckeckHandler(w http.ResponseWriter, r *http.Request) {

	envelope := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, envelope, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and cound not process your request", http.StatusInternalServerError)
	}
}
